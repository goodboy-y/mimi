/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"mimi/internal/route"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start http server",
	Long:  `Start http server.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		r := gin.Default()
		r.Static("/assets/", "./assets/")
		r.LoadHTMLGlob("ui/*")
		//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
		r.Use(Cors()) // 跨域设置
		route.Init(r)
		r.Run(":" + port)
	},
}

// 开启跨域函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringP("port", "p", "8222", "http服务端口")
}
