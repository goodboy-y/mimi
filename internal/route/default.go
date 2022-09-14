package route

import (
	"github.com/gin-gonic/gin"
	"mimi/internal/persist"
)

func Init(r *gin.Engine) {
	r.GET("/requestInfo", func(c *gin.Context) {
		//value := c.Query("pageSize")
		//value := c.Query("pageNo")
		search := persist.Search(10, 0)
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    search,
		})
	})
}
