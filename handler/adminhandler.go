package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Admindata struct {
// 	Email    string
// 	Password string
// }

const adminEmail = "admin@gmail.com"
const adminPassword = "hello"

func Adminpanel(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	admincookie, err := c.Cookie("AdminCookie")
	if err != nil || admincookie == "" {
		fmt.Println(err)
		c.Redirect(303, "/adminloginpage")

		return
	} else {

		var user []Signupdata
		Db.Find(&user)
		c.HTML(http.StatusOK, "adminpanel.html", gin.H{"Users": user})
	}

	// var users []Signupdata
	// db.Find(&users)

}

func Adminloginpage(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	admincookie, err := c.Cookie("AdminCookie")
	if err == nil && admincookie != "" {
		c.HTML(http.StatusOK, "adminpanel.html", nil)
		return
	}
	c.HTML(http.StatusOK, "adminloginPage.html", nil)
}
func AdminloginPost(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")

	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println(email, password)
	if email == "" {
		c.HTML(http.StatusUnauthorized, "adminloginPage.html", "invalid entry")
		fmt.Println("email is not given")
		return
	} else if password == "" {
		c.HTML(http.StatusUnauthorized, "adminloginPage.html", "invalid entry")
		fmt.Println("password is not given")
		return
	}

	// result := Db.Where("email=?", email).Where("password=?", password).First(&user)
	// if result.Error != nil || result.RowsAffected == 0 {
	// 	c.HTML(303, "adminloginPage.html", "Email not found")
	// }
	if email == adminEmail && password == adminPassword {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("AdminCookie", email, 3600, "", "", false, true)
		// Db.Where("email=?", email).Where("password=?", password).First(&user)
		var user []Signupdata

		Db.Find(&user)

		// c.Redirect(302, "/admin")
		c.HTML(http.StatusOK, "adminpanel.html", gin.H{"Users": user})
		return
	} else {
		c.HTML(http.StatusOK, "adminloginPage.html", nil)
		fmt.Println("invalid credentials")
		return
	}
	// c.HTML(http.StatusOK, "adminpanel.html", nil)
}
func AdminLogout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminCookie", "", 0, "", "", true, true)
	c.Redirect(303, "/admin")
}
func Search(c *gin.Context) {
	var user []Signupdata
	Db.Find(&user)
	searchQuery := c.DefaultQuery("query", "")
	if searchQuery != "" {
		Db.Where("name ILIKE ?", "%"+searchQuery+"%").Find(&user)

	} else {
		Db.Find(&user)
	}
	c.HTML(200, "adminpanel.html", gin.H{"Users": user})
}
func DeleteUser(c *gin.Context) {
	var user Signupdata
	userId := c.Param("id")
	if err := Db.Where("id = ?", userId).Delete(&user).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	c.Redirect(303, "/admin")
}

func EditUser(c *gin.Context) {
	var user Signupdata

	userId := c.Param("id")
	if err := Db.Where("id= ?", userId).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	c.HTML(200, "edituser.html", gin.H{"Users": user})
}
func UpdateUser(c *gin.Context) {
	var user Signupdata
	userId := c.Param("id")
	if err := Db.Where("id= ?", userId).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	Db.Save(&user)
	c.Redirect(303, "/admin")
}
func CreateUserPage(c *gin.Context) {
	c.HTML(200, "createuser.html", nil)
}
func AddNewUser(c *gin.Context) {
	var user Signupdata
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	Db.Save(&user)
	c.Redirect(303, "/admin")
}
