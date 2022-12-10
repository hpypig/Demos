package main
/*
这是一个自由测试gin的demo
包含：静态资源测试
 */
import (
	"github.com/gin-gonic/gin"
)
//
//var db = make(map[string]string)
//
//func setupRouter() *gin.Engine {
//	// Disable Console Color
//	// gin.DisableConsoleColor()
//	r := gin.Default()
//
//	// Ping test
//	r.GET("/ping", func(c *gin.Context) {
//		c.String(http.StatusOK, "pong")
//	})
//
//	// Get user value
//	r.GET("/user/:name", func(c *gin.Context) {
//		user := c.Params.ByName("name")
//		value, ok := db[user]
//		if ok {
//			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
//		} else {
//			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
//		}
//	})
//
//	// Authorized group (uses gin.BasicAuth() middleware)
//	// Same than:
//	// authorized := r.Group("/")
//	// authorized.Use(gin.BasicAuth(gin.Credentials{
//	//	  "foo":  "bar",
//	//	  "manu": "123",
//	//}))
//	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
//		"foo":  "bar", // user:foo password:bar
//		"manu": "123", // user:manu password:123
//	}))
//
//	/* example curl for /admin with basicauth header
//	   Zm9vOmJhcg== is base64("foo:bar")
//
//		curl -X POST \
//	  	http://localhost:8080/admin \
//	  	-H 'authorization: Basic Zm9vOmJhcg==' \
//	  	-H 'content-type: application/json' \
//	  	-d '{"value":"bar"}'
//	*/
//	authorized.POST("admin", func(c *gin.Context) {
//		user := c.MustGet(gin.AuthUserKey).(string)
//
//		// Parse JSON
//		var json struct {
//			Value string `json:"value" binding:"required"`
//		}
//
//		if c.Bind(&json) == nil {
//			db[user] = json.Value
//			c.JSON(http.StatusOK, gin.H{"status": "ok"})
//		}
//	})
//
//	return r
//}

func main() {
	//r := setupRouter()
	// Listen and Server in 0.0.0.0:8080

	r := gin.Default()

	// 最后*所表示的文件里不能有文件夹 **表示文件夹、*表示文件
	r.LoadHTMLGlob("dist/**/**/*")    // 添加入口index.html
	// 下两行有问题，*包含的文件里有文件夹，不能包含文件夹
	r.LoadHTMLGlob("dist/**/*.html")
	r.LoadHTMLGlob("dist/*.html")

	r.Static("/css", "./css") 	// 添加资源路径
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(200, "test.html", gin.H{})
	//})



	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "test.html", gin.H{})
	})
	//r.GET("/3", func(c *gin.Context) {
	//	c.HTML(200, "test3.html", gin.H{})
	//})

	//r.GET("/", func(c *gin.Context) {
	//	c.String(401, "test.html")
	//})

	r.Run(":9991")
}
