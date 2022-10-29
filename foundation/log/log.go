package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/medivh-jay/lfshook"
	"github.com/olivere/elastic/v7"
	esconfig "github.com/olivere/elastic/v7/config"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"shequn1/foundation/app"
	"strings"
	"time"
)

var (
	// Terminal 默认终端log输出
	Terminal = logrus.New()
	// Logger 对外 Logger 对象
	Logger = Terminal
	config logConfig
)

type logConfig struct {
	Es    bool   `toml:"es"`
	Index string `toml:"index"`
	Dir   string `toml:"dir"`
}

type elasticConfig struct {
	URL         string `toml:"url"`
	Index       string `toml:"index"`
	Username    string `toml:"username"`
	Password    string `toml:"password"`
	Shards      int    `toml:"shards"`
	Replicas    int    `toml:"replicas"`
	Sniff       bool   `toml:"sniff"`
	HealthCheck bool   `toml:"health"`
	InfoLog     string `toml:"info_log"`
	ErrorLog    string `toml:"error_log"`
	TraceLog    string `toml:"trace_log"`
}

func startEsLog() {
	var elasticConfig elasticConfig
	_ = app.Config().Bind("application", "elasticsearch", &elasticConfig)
	elasticConfig.Index = config.Index
	var retryTimes = 0

retry:
	client, err := elastic.NewClientFromConfig(&esconfig.Config{
		URL:         elasticConfig.URL,
		Index:       elasticConfig.Index,
		Username:    elasticConfig.Username,
		Password:    elasticConfig.Password,
		Shards:      elasticConfig.Shards,
		Replicas:    elasticConfig.Replicas,
		Sniff:       &elasticConfig.Sniff,
		Healthcheck: &elasticConfig.HealthCheck,
		Infolog:     elasticConfig.InfoLog,
		Errorlog:    elasticConfig.ErrorLog,
		Tracelog:    elasticConfig.TraceLog,
	})
	if err != nil {
		retryTimes++
		Logger.Info("try to reconnect: elasticsearch")
		if retryTimes < 3 {
			time.Sleep(time.Duration(retryTimes) * 5 * time.Second)
			goto retry
		}
		Logger.Fatalln(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "development"
	}

	hook, err := NewBulkProcessorElasticHookWithFunc(client, hostname, logrus.DebugLevel, func() string {
		return fmt.Sprintf("%s-%s", config.Index, time.Now().Format("2006-01-02"))
	})
	if err != nil {
		Terminal.Debug(err)
	}
	es := logrus.New()
	es.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	es.Hooks.Add(hook)
	// 这个操作太影响性能,release不启用
	//es.SetReportCaller(gin.Mode() != gin.ReleaseMode)
	es.SetNoLock()
	if gin.Mode() != gin.ReleaseMode {
		Terminal.Level = logrus.DebugLevel
	}
	Logger = es
}

// Start 启动日志服务
func Start() {
	_ = app.Config().Bind("application", "log", &config)
	Terminal.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	// 这个操作太影响性能,release不启用
	//Terminal.SetReportCaller(gin.Mode() != gin.ReleaseMode)
	Terminal.SetNoLock()
	if config.Es {
		startEsLog()
	} else {
		infoWriter, _ := rotatelogs.New(config.Dir+"/info_%Y%m%d.log",
			rotatelogs.WithMaxAge(time.Duration(86400)*time.Second), rotatelogs.WithRotationTime(time.Duration(86400)*time.Second))
		errorWriter, _ := rotatelogs.New(config.Dir+"/error_%Y%m%d.log",
			rotatelogs.WithMaxAge(time.Duration(86400)*time.Second), rotatelogs.WithRotationTime(time.Duration(86400)*time.Second))
		debugWriter, _ := rotatelogs.New(config.Dir+"/debug_%Y%m%d.log",
			rotatelogs.WithMaxAge(time.Duration(86400)*time.Second), rotatelogs.WithRotationTime(time.Duration(86400)*time.Second))
		warnWriter, _ := rotatelogs.New(config.Dir+"/warn_%Y%m%d.log",
			rotatelogs.WithMaxAge(time.Duration(86400)*time.Second), rotatelogs.WithRotationTime(time.Duration(86400)*time.Second))

		if gin.Mode() != gin.ReleaseMode {
			Terminal.Level = logrus.DebugLevel
		}

		Terminal.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel:  infoWriter,
				logrus.ErrorLevel: errorWriter,
				logrus.DebugLevel: debugWriter,
				logrus.WarnLevel:  warnWriter,
			}, &logrus.JSONFormatter{}))
		Terminal.AddHook(NewContextHook())
	}

	app.Register("logger", Logger)
}

// ContextHook for log the call context
type contextHook struct {
	Field  string
	Skip   int
	levels []logrus.Level
}

// NewContextHook use to make an hook
// 根据上面的推断, 我们递归深度可以设置到5即可.
func NewContextHook(levels ...logrus.Level) logrus.Hook {
	hook := contextHook{
		Field:  "line",
		Skip:   5,
		levels: levels,
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}
	return &hook
}

// Levels implement levels
func (hook contextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire implement fire
func (hook contextHook) Fire(entry *logrus.Entry) error {
	entry.Data[hook.Field] = findCaller(hook.Skip)
	return nil
}

// 对caller进行递归查询, 直到找到非logrus包产生的第一个调用.
// 因为filename我获取到了上层目录名, 因此所有logrus包的调用的文件名都是 logrus/...
// 因此通过排除logrus开头的文件名, 就可以排除所有logrus包的自己的函数调用
func findCaller(skip int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// 这里其实可以获取函数名称的: fnName := runtime.FuncForPC(pc).Name()
// 但是我觉得有 文件名和行号就够定位问题, 因此忽略了caller返回的第一个值:pc
// 在标准库log里面我们可以选择记录文件的全路径或者文件名, 但是在使用过程成并发最合适的,
// 因为文件的全路径往往很长, 而文件名在多个包中往往有重复, 因此这里选择多取一层, 取到文件所在的上层目录那层.
func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	//fmt.Println(file)
	//fmt.Println(line)
	if !ok {
		return "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}
