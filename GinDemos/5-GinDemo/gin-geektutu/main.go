package main
/*
这是一个跟练gin框架的demo
https://geektutu.com/post/quick-go-gin.html
 */
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// H is a shortcut for map[string]interface{}
// type H map[string]interface{}

type student struct {
	Name string
	Age  int8
}


func main() {
	r := gin.Default() // 实例为 WSGI 应用程序
	route(r)

	//r.Run() // 默认 0.0.0.0:8080
	r.Run(":9999")



}
func route(r *gin.Engine){
	// curl "http://localhost:9999/"
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "Hello, Geektutu")
	//})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Who are you?")
	})

	// 匹配 /user/geektutu
	// curl “http://localhost:9999/user/geektutu”
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name") // 解析获取路径某位置名
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 匹配users?name=xxx&role=xxx，role可选
	// curl “http://localhost:9999/users?name=Tom&role=student”
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name") // 获取路径下请求参数
		role := c.DefaultQuery("role", "teacher") // 带默认值的参数查询
		c.String(http.StatusOK, "%s is a %s", name, role)
	})


	// POST
	// Linux: curl http://localhost:9999/form -X POST -d 'username=geektutu&password=1234'
	// windows: curl http://localhost:9999/form -X POST -d "username=geektutu&password=1234"
	// windows下，curl中的单引号要变双引号，json数据中的双引号要加 \ 转义
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// GET 和 POST 混合
	// curl "http://localhost:9999/posts?id=9876&page=7"  -X POST -d "username=geektutu&password=1234"
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	// map 参数
	// curl -g "http://localhost:9999/post?ids[Jack]=001&ids[Tom]=002" -X POST -d ”names[a]=Sam&names[b]=David“
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// 重定向
	// curl -i ”http://localhost:9999/redirect“
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
	/*
	HTTP/1.1 301 Moved Permanently
	Content-Type: text/html; charset=utf-8
	Location: /
	Date: Thu, 08 Aug 2019 17:22:14 GMT
	Content-Length: 36
	 */


	// curl "http://localhost:9999/goindex"
	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/" // 重定向到 /
		//c.Request.URL.RawQuery = "name=john" // 还可以重置获取的 get 参数等
		r.HandleContext(c)
	})



	// group routes 分组路由，将同样的前缀分到一起
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	// group: v1
	// curl "http://localhost:9999/v1/posts"
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	// group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	//-------------------上传文件
	// 上传单个文件
	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

	// 多个文件
	r.POST("/upload2", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})

	//---------------模板渲染
	// curl "http://localhost:9999/arr"
	r.LoadHTMLGlob("templates/*")
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	/*
	Gin默认使用模板Go语言标准库的模板text/template和html/template，
	语法与标准库一致，支持各种复杂场景的渲染。
	 */

	//-----------------------中间件
	// Logger、Recovery 帮助我们打印日志输出和 painc 处理。
	// func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes

	//作用于全局
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	// r.Use(gin.Logger(),gin.Recovery()) // 可连续设置多个中间件

	//// 作用于单个路由（中间件执行是有先后顺序的）
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	//
	//// 作用于某个组
	//authorized := r.Group("/")
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//}

	//----------------具体中间件
	// gin.BasicAuth 基本认证中间件
	//// 首页必须输入密码才能显示（但 c.String 不知道为什么不行）
	//r.Use(gin.BasicAuth(gin.Accounts{
	//	"admin": "123456",
	//}))
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, "首页")
	//})

	// 对分组路由设置权限
	adminGroup := r.Group("/adminG")
	adminGroup.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))
	adminGroup.GET("/index", func(c *gin.Context) {
		c.JSON(200, "后台首页")
	})


	//自定义中间件 costTime 使用
	r.Use()
	r.GET("/costTime", func(c *gin.Context) {
		c.JSON(200, "首页")
	},costTime())
}

// -------------------自定义中间件
// 中间件其实就是 gin.HandlerFunc 类型函数
// 通过自定义中间件,我们可以很方便地拦截请求，
// 来做一些我们需要做的事情，比如日志记录、授权校验、各种过滤等等。
func costTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求前获取当前时间
		nowTime := time.Now()

		//请求处理
		c.Next()
		/*
			这个是执行后续中间件请求处理的意思（含没有执行的中间件和我们定义的GET方法处理），
			这样我们才能获取执行的耗时。也就是在c.Next方法前后分别记录时间，就可以得出耗时。
		*/

		//处理后获取消耗时间
		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}
}
