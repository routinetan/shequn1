package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func genModuleCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("test cmd function execute.")

	if len(args) > 0 {
		i := 0
		for i = 0; i < len(args); i++ {

			fmt.Printf("  args[%d]:%s\r\n", i, args[i])

		}

	}
}

func genModelCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("test cmd function execute.")
	
	if len(args) > 0 {
		i := 0
		for i = 0; i < len(args); i++ {

			fmt.Printf("  args[%d]:%s\r\n", i, args[i])

		}

	}
}

func genServiceCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("test cmd function execute.")

	if len(args) > 0 {
		i := 0
		for i = 0; i < len(args); i++ {

			fmt.Printf("  args[%d]:%s\r\n", i, args[i])

		}

	}
}

func genDAOCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("test cmd function execute.")

	if len(args) > 0 {
		i := 0
		for i = 0; i < len(args); i++ {

			fmt.Printf("  args[%d]:%s\r\n", i, args[i])

		}

	}
}

func genCrudCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("test cmd function execute.")

	if len(args) > 0 {
		i := 0
		for i = 0; i < len(args); i++ {

			fmt.Printf("  args[%d]:%s\r\n", i, args[i])

		}

	}
}

func genViewCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("test cmd function execute.")

	if len(args) > 0 {
		i := 0
		for i = 0; i < len(args); i++ {

			fmt.Printf("  args[%d]:%s\r\n", i, args[i])

		}

	}
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "automatically generate go files for ORM models...",
	Long: `
USAGE
    qcli gen COMMAND [ARGUMENT] [OPTION]
COMMAND
    app        生成一个新的模块
    curd       install or update qcli to system in default...
    controller automatically generate go files for ORM models...
    service    extra features for go modules...
    dao        running go codes with hot-compiled-like feature...
    model      create and initialize an empty qcli project...
    view       show more information about a specified command
    migrateJob    create migration file 
    api        packing any file/directory to a resource file, or a go file...
    vueview    packing any file/directory to a resource file, or a go file...
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

var modelCmd = &cobra.Command{
	Use:   "gen model",
	Short: "gen model",
	Long:  `这条命令可以用来生成model`,
	Run:   testCmdFunc,
}

var curdCmd = &cobra.Command{
	Use:   "gen curd",
	Short: "gen curd",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: testCmdFunc,
}

var daoCmd = &cobra.Command{
	Use:   "gen dao",
	Short: "gen dao",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: testCmdFunc,
}

var viewCmd = &cobra.Command{
	Use:   "gen view",
	Short: "gen view",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: testCmdFunc,
}

var genMigrateCmd = &cobra.Command{
	Use:   "migrateJob",
	Short: "gen view",
	Long:  `create migration file`,
	Run:   testCmdFunc,
}

func init() {
	genCmd.AddCommand(daoCmd)
	genCmd.AddCommand(curdCmd)
	genCmd.AddCommand(modelCmd)
	genCmd.AddCommand(genMigrateCmd)
}
