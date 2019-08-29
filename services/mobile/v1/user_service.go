/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/8/29 15:56
*/
package userService

import "qianshen/models"

type UserService struct {
}

func (us *UserService) GetUserByName(user models.User) {
	models.DB().First(&user, "user_name=?", user.UserName)
}
