package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.New()
	//r.LoadHTMLFiles("./pic/test.html")
	r.LoadHTMLGlob("./pic/*.html")    // 添加入口index.html
	r.Static("/static", "./pic") 	// 从 /static 路径请求的图片，去 ./pic 找
	r.Static("/static2", "./") 	// 添加资源路径
	r.GET("/",func(c *gin.Context){
		c.HTML(200, "test.html", gin.H{"jsonField":1})
		//c.JSON(http.StatusOK, map[string]interface{}{"jsonField":1})
	})
	r.Run(":9096")
}
