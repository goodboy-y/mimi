package route

import (
	"github.com/gin-gonic/gin"
	"mimi/internal/persist"
	"net/http"
)

func Init(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

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
