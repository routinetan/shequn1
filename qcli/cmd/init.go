package cmd

import "github.com/spf13/cobra"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create and initialize an empty qcli project...",
	Run:   testCmdFunc,
}
