package service

import (
	"ginchat/models"

	"net/http"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	ind.Execute(c.Writer, "index")
}

// 用户注册
func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	ind.Execute(c.Writer, "register")
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles(
		"views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/main.html",
	)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	user := models.UserBasic{}

	userId, _ := strconv.Atoi(c.Query("userId"))
	user.ID = uint(userId)

	token := c.Query("token")
	user.Identity = token

	ind.Execute(c.Writer, user)
}

func Chat(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
