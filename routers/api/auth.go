package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"niurenshuo/models"
	"niurenshuo/pkg/e"
	"niurenshuo/pkg/logging"
	"niurenshuo/pkg/util"
)

type Auth struct {
	Identifier   string `valid:"Required; MaxSize(50)"` //标识:手机 邮箱 用户名或第三方的唯一标识
	Credential   string `valid:"Required; MaxSize(50)"` //密码凭证:密码或TOKEN
	IdentityType string `valid:"Required; MaxSize(50)"` //登录类型：手机 邮箱 用户名 或第三方
}

func GetAuth(c *gin.Context) {
	identifier := c.Query("username")
	credential := c.Query("password")
	identityType := c.Query("login_type")
	valid := validation.Validation{}
	a := Auth{Identifier: identifier, Credential: credential, IdentityType: identityType}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(identifier, credential, identityType)
		if isExist {
			token, err := util.GenerateToken(identifier, credential)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
