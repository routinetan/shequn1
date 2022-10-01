package cmd

import "github.com/spf13/cobra"

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "running go codes with hot-compiled-like feature..",
	Run:   testCmdFunc,
}
