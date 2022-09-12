// Package cmd /* Copyright © 2022 liuyu <liuyucupid@qq.com>
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mimi/internal/pkg/core"
	"mimi/internal/pkg/log"
	"os"
)

var version = "1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mimi",
	Short: "分析关键词的代理服务",
	Long:  `分析关键词的代理服务`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		key, _ := cmd.Flags().GetString("key")
		logLevel, _ := cmd.Flags().GetString("log")
		persist, _ := cmd.Flags().GetBool("persist")
		log.LogLevel(logLevel)
		ipt := core.NewIntercept(port, key, persist)
		// 启动服务
		ipt.RunServer()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = version
	rootCmd.Flags().BoolP("version", "v", false, "版本")
	rootCmd.Flags().StringP("port", "p", "8111", "端口")
	rootCmd.Flags().StringP("key", "k", "idCard", "关键词，支持正则")
	rootCmd.Flags().String("log", "info", "日志等级:debug,info,warn,error")
	rootCmd.Flags().Bool("persist", false, "是否持久化")
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本号",
	Long:  `版本号`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
