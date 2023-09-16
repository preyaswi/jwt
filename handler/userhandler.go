package handler

import (
	"fmt"
	"net/http"
	"one/jwt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Signupdata struct {
	Id          uint `gorm:"primary key"`
	Name        string
	PhoneNumber string
	Password    string
	Email       string `gorm:"unique"`
}

func ConnectPostgresDB() {
	var err error
	dsn := "user=postgres dbname=postgres password=preya host=localhost port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Db.AutoMigrate(&Signupdata{})
	if err != nil {
		panic("failed to automigrate the table ")
	}
}
func HomePage(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	cookie, err := c.Cookie("Cookie")
	if err == nil && cookie != "" {
		c.HTML(http.StatusOK, "homepage.html", nil)
		return
	} else {
		c.Redirect(303, "/login")
	}

}
func SignupPage(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	c.HTML(http.StatusFound, "signupPage.html", nil)

}

func SignupPost(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	var data Signupdata
	data.Name = c.Request.FormValue("firstname")
	data.Password = c.Request.FormValue("password")
	data.PhoneNumber = c.Request.FormValue("phonenumber")
	ConfirmPassword := c.Request.FormValue("confirmpassword")
	data.Email = c.Request.FormValue("email")
	if data.Name == "" {

		c.HTML(http.StatusUnauthorized, "signupPage.html", "invalid entry")
		return
	}

	if data.Email == "" {

		c.HTML(http.StatusUnauthorized, "signupPage.html", "invalid entry")

		return
	}
	if data.Password == "" {

		c.HTML(http.StatusUnauthorized, "signupPage.html", "invalid entry")

		return
	}
	if data.PhoneNumber == "" {

		c.HTML(http.StatusUnauthorized, "signupPage.html", "invalid entry")

		return
	}
	if ConfirmPassword != data.Password {

		c.HTML(http.StatusUnauthorized, "signupPage.html", "invalid entry")

		return
	}

	Db.Create(&data)
	fmt.Println("data inserted")
	c.Redirect(303, "/login")

}
func LoginPage(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	cookie, err := c.Cookie("Cookie")
	if err == nil && cookie != "" {
		c.HTML(http.StatusOK, "homepage.html", nil)
		return
	}
	c.HTML(http.StatusOK, "loginPage.html", nil)
}
func LoginPost(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	email := c.Request.FormValue("emailLogin")
	password := c.Request.FormValue("passwordLogin")
	if email == "" || password == "" {
		c.HTML(http.StatusSeeOther, "loginPage.html", "invalid entry")
		fmt.Println("email is not given")
		return
	}
	var user Signupdata
	if err := Db.Where("email=?", email).First(&user).Error; err != nil {
		c.HTML(http.StatusSeeOther, "loginPage.html", "user not found")
		fmt.Println("user not found")
		return
	}
	if user.Password != password {
		c.HTML(http.StatusSeeOther, "loginPage.html", "user not found")
		fmt.Println("invalid credentials")
		return
	}
	token, err := jwt.GenerateToken()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generate token",
		})
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Cookie", token, 3600, "", "", false, true)
	c.Redirect(302, "/")

}
func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Cookie", "", 0, "", "", true, true)

	c.Redirect(303, "/")

}
