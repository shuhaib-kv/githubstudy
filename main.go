package main

import (
	"gin/gorm/controllers"
	"gin/gorm/initializers"

	"github.com/gin-gonic/gin"
)

var R = gin.Default()

func init() {
	R.LoadHTMLGlob("templates/*.html")
	initializers.Loadenv()
	initializers.Connectodb()
}

func main() {
	// R.POST("/Posts", controllers.Postcreate)
	R.GET("/", controllers.HomePage)
	R.GET("/signup", controllers.Signuppage)
	R.POST("/signup", controllers.Signuppagepost)
	R.GET("/login", controllers.LoginGet)
	R.POST("/login", controllers.LoginPost)
	R.GET("/logout", controllers.Logout)
	R.GET("/test", controllers.Test)
	admin := R.Group("admin")
	{
		admin.GET("/", controllers.AdminLogin)
		admin.POST("/", controllers.AdminLoginPost)
		admin.GET("/home", controllers.AdminHome)
		admin.GET("/signout", controllers.AdminLogout)
		admin.GET("/block/:id", controllers.Block)
		admin.GET("/unblock/:id", controllers.Unblock)
	}

	R.Run() // listen and serve on 0.0.0.0:8080
	// /home/{{.ID}}
}
