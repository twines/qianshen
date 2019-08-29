package jwt

import (
	"github.com/gin-gonic/gin"
	"qianshen/pkg/response"
	"qianshen/pkg/util"
	"strings"
)

func Web() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := parseToken(c.GetHeader("token")); len(token) != 2 {
			c.JSON(200, response.Error("issuer error", 4001))
			c.Abort()
		} else {
			if user, err := util.ParseToken(token[1]); err == nil {
				if strings.ToUpper(token[0]) != user.Issuer {
					c.JSON(200, response.Error("token error", 4001))
					c.Abort()
				}
				c.Set("user", user)
			} else {
				c.JSON(200, response.Error("token error", 4001))
				c.Abort()
			}
			c.Next()
		}

	}
}
func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := parseToken(c.GetHeader("token")); len(token) != 2 {
			c.JSON(200, response.Error("issuer error", 4001))
			c.Abort()
		} else {
			if admin, err := util.ParseToken(token[1]); err == nil {
				if strings.ToUpper(token[0]) != admin.Issuer {
					c.JSON(200, response.Error("issuer error", 4001))
					c.Abort()
				}
				c.Set("admin", admin)
			} else {
				c.JSON(200, response.Error("token error", 4001))
				c.Abort()
			}
			c.Next()
		}
	}
}
func Api() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := parseToken(c.GetHeader("token")); len(token) != 2 {
			c.JSON(200, response.Error("issuer error", 4001))
			c.Abort()
		} else {
			if user, err := util.ParseToken(token[1]); err == nil {
				if strings.ToUpper(token[0]) != user.Issuer {
					c.JSON(200, response.Error("token error", 4001))
					c.Abort()
				}
				c.Set("user", user)
			} else {
				c.JSON(200, response.Error("token error", 4001))
				c.Abort()
			}
			c.Next()
		}
	}
}
func parseToken(token string) []string {
	return strings.Split(token, " ")
}
