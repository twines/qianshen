/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/8/29 15:52
*/
package mobileV1

import (
	"github.com/gin-gonic/gin"
	"qianshen/models"
	"qianshen/services/mobile/v1"
)

var (
	us = &userService.UserService{}
)

func Login(c *gin.Context) {
	var user = models.User{}
	_ = c.ShouldBind(&user)
	//us.GetUserByName(user)
	//if user.ID <= 0 {
	//	c.JSON(200, response.Error("用户不存在"))
	//} else {
	//	if md := util.EncodeMD5(c.PostForm("password")); md != user.Password {
	//		c.JSON(200, response.Error("用户不存在"))
	//	} else {
	//		if token, err := util.GenerateToken(user, "mobile"); err == nil {
	//			c.JSON(200, response.Success(token))
	//		} else {
	//			c.JSON(200, response.Error("error"))
	//		}
	//
	//	}
	//}
}
