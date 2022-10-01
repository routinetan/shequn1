package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// 构建根 command 命令。前面我们介绍它还可以有子命令，这个command里没有构建子命令
var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "A brief description of your application",
	Long: `
USAGE
    qcli COMMAND [ARGUMENT] [OPTION]
COMMAND
    env        show current Golang environment variables
    get        install or update qcli to system in default...
    gen        automatically generate go files for ORM models...
    migration  execute databases migrateJob file,eg:change,reset,rollback
    mod        extra features for go modules...
    run        running go codes with hot-compiled-like feature...
    init       create and initialize an empty qcli project...
    help       show more information about a specified command
    pack       packing any file/directory to a resource file, or a go file...
    build      cross-building go project for lots of platforms...
    docker     create a docker image for current qcli project...
    swagger    swagger feature for current project...
    update     update current qcli binary to latest one (might need root/admin permission)
    install    install qcli binary to system (might need root/admin permission)
    version    show current binary version info
OPTION
    -y         all yes for all command without prompt ask 
    -?,-h      show this help or detail for specified command
    -v,-i      show version information
ADDITIONAL
    Use 'qcli help COMMAND' or 'qcli COMMAND -h' for detail about a command, which has '...' 
    in the tail of their comments.
`}

func testCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("test cmd function execute.")

	if len(args) > 0 {
		i := 0
		for i = 0; i < len(args); i++ {

			fmt.Printf("  args[%d]:%s\r\n", i, args[i])

		}

	}
}

// 执行 rootCmd 命令并检测错误
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// 加载运行初始化配置
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(migrateCmd)
	// rootCmd，命令行下读取配置文件，持久化的 flag，全局的配置文件
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.firstappname.yaml)")
	// local flag，本地化的配置
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 初始化配置的一些设置
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) // viper 设置配置文件
	} else { // 上面没有指定配置文件，下面就读取 home 下的 .firstappname.yaml文件
		// 配置文件参数设置
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".firstappname")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil { // 读取配置文件
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
