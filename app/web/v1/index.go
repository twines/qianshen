/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/8/28 9:25
*/
package webV1

import (
	"github.com/gin-gonic/gin"
	"qianshen/pkg/response"
	"qianshen/pkg/util"
)

func Index(c *gin.Context) {
	if admin, ok := c.Get("user"); ok {
		c.JSON(200, response.Success(admin))
	}
}

func Login(c *gin.Context) {
	if token, err := util.GenerateToken("muyu", "123456", 2, "web"); err == nil {
		c.JSON(200, response.Success(token))
	} else {
		c.JSON(200, response.Error("error"))
	}

}
