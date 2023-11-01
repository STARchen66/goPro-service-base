package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// Env 存储全局选项 --env 的值
var Env string

// RegisterGlobalFlags 注册全局选项（flag）
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env file, example: --env=testing will use .env.testing file")
}

// RegisterDefaultCmd 在用户不输入运行参数时，使用默认参数运行
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {

	arg := os.Args[1:]
	//从用户输入中获取第一个参数
	cmd, _, err := rootCmd.Find(arg)
	var firstArg string
	if len(arg) > 0 {
		firstArg = arg[0]
	} else {
		firstArg = ""
	}

	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, arg...)
		rootCmd.SetArgs(args)
	}
}
