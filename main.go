package main

import (
	"one/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("template/*.html")
	router.Static("/static", "./static")
	handler.ConnectPostgresDB()
	router.GET("/", handler.HomePage)
	router.GET("/signup", handler.SignupPage)
	router.POST("/signuppost", handler.SignupPost)
	router.GET("/login", handler.LoginPage)
	router.POST("/loginpost", handler.LoginPost)
	router.GET("/logout", handler.Logout)

	router.GET("/adminloginpage", handler.Adminloginpage)
	router.POST("/adminlogin", handler.AdminloginPost)
	router.GET("/admin", handler.Adminpanel)
	router.GET("/adminlogout", handler.AdminLogout)

	router.GET("/admin/search", handler.Search)
	router.POST("/admin/delete/:id", handler.DeleteUser)
	router.GET("/admin/edituser/:id", handler.EditUser)
	router.POST("/admin/updateuser/:id", handler.UpdateUser)
	router.GET("/createuser", handler.CreateUserPage)
	router.POST("/adduser", handler.AddNewUser)
	router.Run(":8081")

}
