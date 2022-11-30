package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"html/template"
	"os"
	"strings"
	time2 "time"
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

func genMigrateCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("migrate latest have one params")
		return
	}
	table := args[0]
	migrateName := args[0]
	if strings.Contains(table, "_") {
		words := strings.Split(table, "_")
		for k, v := range words {
			words[k] = strings.ToUpper(string(v[0])) + v[1:]
		}
		migrateName = strings.Join(words, "")
	}
	module, _ := cmd.Flags().GetString("m")
	migrateFile := fmt.Sprintf("../internal/%s/database/migrate/%s.go", module, table)
	fp, err := os.OpenFile(migrateFile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	time := time2.Now().Format("200601021504")
	migrateName = fmt.Sprintf("%s_%s", migrateName, time)
	tpl, err := template.ParseFiles("tpl/migrate/migrate.tpl")
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	data["Name"] = migrateName

	tpl.Execute(fp, data)
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

func Conv2CamlStyle(name string) string {
	if strings.Contains(name, "_") {
		words := strings.Split(name, "_")
		for k, v := range words {
			words[k] = strings.ToUpper(string(v[0])) + v[1:]
		}
		name = strings.Join(words, "")
	} else {
		name = strings.ToUpper(string(name[0])) + name[1:]
	}
	return name
}

func genController(cmd *cobra.Command, args []string) {
	ctrName := args[0]
	Cval := ctrName
	ctrName = Conv2CamlStyle(ctrName)
	module, _ := cmd.Flags().GetString("m")
	migrateFile := fmt.Sprintf("../internal/%s/controller/%s.go", module, Cval)
	fp, err := os.OpenFile(migrateFile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	tpl, err := template.ParseFiles("tpl/controller/controller.tpl")
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	data["CVal"] = Cval
	data["CStruct"] = ctrName
	tpl.Execute(fp, data)
}

func genService(cmd *cobra.Command, args []string) {
	ctrName := args[0]
	Cval := ctrName
	ctrName = Conv2CamlStyle(ctrName)
	module, _ := cmd.Flags().GetString("m")
	migrateFile := fmt.Sprintf("../internal/%s/service/%s.go", module, Cval)
	fp, err := os.OpenFile(migrateFile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	tpl, err := template.ParseFiles("tpl/service/controller.tpl")
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	data["CVal"] = Cval
	data["CStruct"] = ctrName
	tpl.Execute(fp, data)
}

func genModel(cmd *cobra.Command, args []string) {
	ctrName := args[0]
	Cval := ctrName
	ctrName = Conv2CamlStyle(ctrName)
	module, _ := cmd.Flags().GetString("m")
	migrateFile := fmt.Sprintf("../internal/%s/model/%s.go", module, Cval)
	fp, err := os.OpenFile(migrateFile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	tpl, err := template.ParseFiles("tpl/model/controller.tpl")
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	data["CVal"] = Cval
	data["CStruct"] = ctrName
	tpl.Execute(fp, data)
}

func genControllerCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("controller gen latest have one params")
		return
	}
	genController(cmd, args)
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
	genController(cmd, args)
	genService(cmd, args)
	genModel(cmd, args)
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
    migrate    create migration file 
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
	Use:   "model",
	Short: "gen model",
	Long:  `这条命令可以用来生成model`,
	Run:   testCmdFunc,
}

var genCtrlCmd = &cobra.Command{
	Use:   "ctrl",
	Short: "gen controller",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: genControllerCmdFunc,
}

var curdCmd = &cobra.Command{
	Use:   "curd",
	Short: "gen curd",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: testCmdFunc,
}

var daoCmd = &cobra.Command{
	Use:   "dao",
	Short: "gen dao",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: testCmdFunc,
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "gen view",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: testCmdFunc,
}

var genMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "gen migrate",
	Long:  `create migration file`,
	Run:   genMigrateCmdFunc,
}

func init() {
	genCmd.AddCommand(daoCmd)
	genCmd.AddCommand(curdCmd)
	genCmd.AddCommand(modelCmd)
	genMigrateCmd.Flags().String("m", "admin", "change a module")
	genCmd.AddCommand(genMigrateCmd)
	genCtrlCmd.Flags().String("m", "admin", "change a module")
	genCmd.AddCommand(genCtrlCmd)
}
