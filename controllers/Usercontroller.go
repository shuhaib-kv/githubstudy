package controllers

import (
	"fmt"
	"gin/gorm/initializers"
	"gin/gorm/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/tawesoft/golib/v2/dialog"
)

func Userloggedin(g *gin.Context) bool {
	session, err := Store.Get(g.Request, "user")
	if session.Values["email"] == nil {
		return false
	}
	fmt.Println(err)
	return true
}

var Store = sessions.NewCookieStore([]byte("saffu"))
var P interface{}

func init() {

}
func HomePage(c *gin.Context) {
	session, _ := Store.Get(c.Request, "user")
	String := session.Values["email"]
	h := String.(string)
	fmt.Println(h)
	splt := strings.TrimSuffix(h, "@gmail.com")
	ok := Userloggedin(c)
	if !ok {
		c.Redirect(303, "/signup")
		return
	}
	c.HTML(200, "homepage.html", gin.H{
		"name": splt,
	})

}

func Signuppage(c *gin.Context) {
	c.HTML(200, "signup.html", nil)
}

func Signuppagepost(c *gin.Context) {

	// fmt.Fprintf(c.Writer, "hai")
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println("error parsing form")
	}
	Names := c.PostForm("name")

	Emails := c.PostForm("email")
	// Name:=c.PostForm("name")
	password := c.PostForm("password")
	user := models.User{
		Name:     Names,
		Email:    Emails,
		Password: password,
	}
	initializers.DB.Create(&user)
	c.Redirect(301, "/login")
	// fmt.Println(c.Request.PostForm)
}
func LoginGet(c *gin.Context) {

	c.HTML(200, "login.html", nil)
}
func LoginPost(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println("error parsing form")
	}

	Emails := c.PostForm("email")
	password := c.PostForm("password")
	var user models.User
	var status bool
	initializers.DB.Raw("select email,password,block_status FROM users where email=?", Emails).Scan(&user)
	if user.Email == Emails && user.Password == password && user.Block_status == false {
		status = true
		fmt.Println(user.Email, user.Password, Emails, password)
	}
	if !status {
		dialog.Alert("wrong username and password")
		c.Redirect(303, "/login")
		return
	}
	sessions, err := Store.Get(c.Request, "user")
	sessions.Values["email"] = Emails

	P = sessions.Values["email"]
	sessions.Save(c.Request, c.Writer)
	fmt.Println(err)

	c.Redirect(301, "/")

}
func Logout(c *gin.Context) {
	sessions, err := Store.Get(c.Request, "user")
	fmt.Println("hai")
	fmt.Println(sessions.Values["email"])
	sessions.Options.MaxAge = -1
	sessions.Save(c.Request, c.Writer)
	fmt.Println(err)
	c.Redirect(303, "/signup")
}

func Test(c *gin.Context) {

	var user []models.User
	initializers.DB.Raw("SELECT * FROM users ").Scan(&user)
	// email:=user.Email
	// password:=user.Password
	// P.Header="hellooooi"

	c.HTML(200, "test.html", gin.H{
		"user": user,
	})
	// c.Render(200,P)
}
