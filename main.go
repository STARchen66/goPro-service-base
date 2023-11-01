package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"goImPro-service/cmd"
	"goImPro-service/pkg/console"
	"goImPro-service/src/config"
	"goImPro-service/src/service/bootstrap"
	"os"
)

func init() {
	config.InitConfig("config.yaml")
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "im",
		Short: "mc-im-system",
		Long: console.ToGreen("A Fast and Flexible Im System build with love by mc in Go." +
			"Complete documentation is available at https://hugo.spf13.com"),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			bootstrap.LoadConfiguration()
		},
	}

	//扩展命令可选项
	rootCmd.AddCommand(
		cmd.AppCmdServe)

	cmd.RegisterDefaultCmd(rootCmd, cmd.AppCmdServe)
	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("命令启动失败 %v: %s", os.Args, err.Error()))
	}
}
