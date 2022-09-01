// Package cmd /* Copyright © 2022 liuyu <liuyucupid@qq.com>
package cmd

import (
	"fmt"
	"log"
	gt "mimi/pkg"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mimi",
	Short: "代理",
	Long:  `代理服务器.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		key, _ := cmd.Flags().GetString("key")
		ipt := &gt.Intercept{
			Ip: "0.0.0.0:" + port,
			HttpPackageFunc: func(pack *gt.HttpPackage) {
				if strings.Index(pack.ContentType, "json") != -1 && strings.Contains(pack.Json(), key) {
					log.Printf("json: url = %s\r\n ", pack.Url.String())
				}
				if strings.Index(pack.ContentType, "html") != -1 && strings.Contains(pack.Html(), key) {
					log.Printf("html: url = %s\r\n ", pack.Url.String())
				}
			},
		}
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
	rootCmd.Flags().BoolP("version", "v", false, "")
	rootCmd.Flags().StringP("port", "p", "8111", "")
	rootCmd.Flags().StringP("key", "k", "idCard", "")
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
