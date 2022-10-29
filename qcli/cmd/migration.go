package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

type MigrateHead struct {
	ID         int
	CurVersion int
	LastTime   string
}

type MigrateList struct {
	ID       int
	Name     string
	Batch    int
	Author   string
	UpdataAt string
}

var migrationResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset the database migrateJob",
	Long: `
USAGE
    qcli migration:reset [ARGUMENT] [OPTION]
OPTION
    -y         all yes for all command without prompt ask 
    -?,-h      show this help or detail for specified command
    -v,-i      show version information
ADDITIONAL
    Use 'qcli help COMMAND' or 'qcli COMMAND -h' for detail about a command, which has '...' 
    in the tail of their comments.
`,
	Run: testCmdFunc,
}

var migrationRollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "rollback database to last version",
	Long: `
USAGE
    qcli migration:COMMAND [ARGUMENT] [OPTION]
COMMAND
    rollback        回滚到上一次
    init            迁移初始化
    reset           清楚数据库所有表,除了保留migrate迁移表
    refresh         重置迁移表，撤销所有迁移后重新运行
    status          迁移任务处理的状态
OPTION
    -y         all yes for all command without prompt ask 
    -?,-h      show this help or detail for specified command
    -v,-i      show version information
ADDITIONAL
    Use 'qcli help COMMAND' or 'qcli COMMAND -h' for detail about a command, which has '...' 
    in the tail of their comments.
`,
	Run: testCmdFunc,
}

var migrationRefreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "rollback the database to first-version and execute to the last migration",
	Long: `
USAGE
    qcli migration:COMMAND [ARGUMENT] [OPTION]
COMMAND
    rollback        回滚到上一次
    init            迁移初始化
    reset           清楚数据库所有表,除了保留migrate迁移表
    refresh         重置迁移表，撤销所有迁移后重新运行
    status          迁移任务处理的状态
OPTION
    -y         all yes for all command without prompt ask 
    -?,-h      show this help or detail for specified command
    -v,-i      show version information
ADDITIONAL
    Use 'qcli help COMMAND' or 'qcli COMMAND -h' for detail about a command, which has '...' 
    in the tail of their comments.
`,
	Run: testCmdFunc,
}

func migrationRunCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("test entitys function execute.")

	//if len(args) > 3 {
	//	fmt.Println(cmd.Long)
	//	return
	//}
	fmt.Println("entitys")
	//读取版本号
	//遍历文件夹，获取未执行的文件

	//执行迁移

	// 1st example: list files
	//pid, err := os.StartProcess("go", []string{"generate -x pre/maing.go"}, procAttr)
	//if err != nil {
	//	fmt.Printf("Error %v starting process!", err) //
	//	os.Exit(1)
	//}
	//fmt.Printf("The process id is %v", pid)
	genCmd := exec.Command("go", "generate", "-x", "pre/main.go")
	genCmd.Dir = "E:/goproject/shequn1/qcli"
	err := genCmd.Run()
	genCmd.Stdout = os.Stdout
	genCmd.Stderr = os.Stderr
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	//fmt.Printf("combined out:\n%s\n", string(out))
	//生成并写入迁移执行之路
}

var migrateCmd = &cobra.Command{
	Use:   "migration",
	Short: "execute databases migrateJob file,eg:change,reset,rollback",
	Long: `
USAGE
    qcli migration:COMMAND [ARGUMENT] [OPTION]
COMMAND
    rollback        回滚到上一次
OPTION
    -?,-h      show this help or detail for specified command
ADDITIONAL
    Use 'qcli migration:COMMAND' or 'qcli migration -h' for detail about a command, which has '...' 
    in the tail of their comments.
`,
	Run: migrationRunCmdFunc,
}

func init() {
	migrateCmd.AddCommand(migrationRollbackCmd)
}
