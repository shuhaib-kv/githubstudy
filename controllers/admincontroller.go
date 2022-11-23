package controllers

import (
	"fmt"
	"gin/gorm/initializers"
	"gin/gorm/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tawesoft/golib/v2/dialog"
)

var Password = "safwan"
// var Hash, err = HashPassword(Password)
var AdminDB = map[string]string{
	"password": "safwan",
	"email":    "safwan@gmail.com",
}

func AdminLogin(c *gin.Context) { // admin login page
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	c.Writer.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	c.Writer.Header().Set("Expires", "0")
	// ok := IsAdminloggeedin(c)
	// if ok {
	// 	c.Redirect(303, "/admin/home")
	// } else {

		c.HTML(200, "adminlogin.html", nil)
	

}
func AdminLoginPost(c *gin.Context) { // admin login page post

	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	c.Writer.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	c.Writer.Header().Set("Expires", "0")
fmt.Println("hai)
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println("error parsing form")
	}
	email := c.PostForm("email")
	password := c.PostForm("password")
	// match := bcrypt.CompareHashAndPassword([]byte(Hash), []byte(password))

	if AdminDB["email"] == email && AdminDB["password"] == password && c.Request.Method == "POST" {
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		c.Writer.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		c.Writer.Header().Set("Expires", "0")

		session, err := Store.Get(c.Request, "admin")
		session.Values["id"] = "email"
		fmt.Println(session.Values["email"])
		session.Save(c.Request, c.Writer)
		fmt.Println(err)

		c.Redirect(303, "/admin/home")
	} else {
		dialog.Alert("wrong user name and password")
		c.Redirect(303, "/admin")
	}

}

var Status models.Page

func AdminHome(c *gin.Context) { // admin Homepage
	Status.Status = true



	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	c.Writer.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	c.Writer.Header().Set("Expires", "0")
	var user []models.User
	initializers.DB.Raw("SELECT * FROM users ORDER BY id ASC").Scan(&user)


	c.HTML(200, "adminhomepage.html", gin.H{
		"user":   user,
		"status": Status.Status,
	})
}
func AdminLogout(c *gin.Context) { // adminLogout page
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	c.Writer.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	c.Writer.Header().Set("Expires", "0")
	session, err := Store.Get(c.Request, "admin")
	session.Options.MaxAge = -1
	fmt.Println(session.Values["email"])
	session.Save(c.Request, c.Writer)
	fmt.Println(err)
	c.Redirect(303, "/admin")
	// c.HTML(200, "adminlogin.html", nil)
}

// /
func IsAdminloggeedin(c *gin.Context) bool {
	session, err := Store.Get(c.Request, "admin")
	if session.Values["email"] == nil {
		return false
	}
	fmt.Println(err)
	return true
}

func Block(c *gin.Context) {
	fmt.Println("hello")
	params := c.Param("id")
	fmt.Println(params)
	page, _ := strconv.Atoi(params)
	fmt.Println("poooy mwone")
	var users models.User
	initializers.DB.Raw("update users SET block_status=true WHERE id=?", page).Scan(&users)
	c.Redirect(303, "/admin/home")
}
func Unblock(c *gin.Context) {

	params := c.Param("id")
	fmt.Printf("%T", params)
	page, _ := strconv.Atoi(params)
	fmt.Println("poooy mwone")
	var users models.User
	initializers.DB.Raw("update users SET block_status=false WHERE id=?", page).Scan(&users)
	c.Redirect(303, "/admin/home")
}
