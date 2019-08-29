/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/8/28 9:25
*/
package adminV1

import (
	"github.com/gin-gonic/gin"
	"qianshen/pkg/response"
)

func Index(c *gin.Context) {
	if admin, ok := c.Get("admin"); ok {
		c.JSON(200, response.Success(admin))
	}
}

func Login(c *gin.Context) {

}
