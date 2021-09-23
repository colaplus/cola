package account

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mercury/common"
	"mercury/dal/db"
	"mercury/id_gen"
	"mercury/middleware/account"
	"mercury/util"
)

func LoginHandle(c *gin.Context) {

	account.ProcessRequest(c)
	var userInfo common.UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	if len(userInfo.Password) == 0 || len(userInfo.Username) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	err = db.Login(&userInfo)
	if err == db.ErrUserNotExists {
		util.ResponseError(c, util.ErrCodeUserNotExist)
		return
	}
	if err == db.ErrUserPasswordWrong {
		util.ResponseError(c, util.ErrCodeUserPasswordWrong)
		return
	}

	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	//用户登录成功，那么我们设置user_id到用户的session当中
	account.SetUserId(userInfo.UserId, c)

	account.ProcessResponse(c)
	util.ResponseSuccess(c, nil)
}

func RegisterHandle(c *gin.Context) {
	var userInfo common.UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	if len(userInfo.Email) == 0 || len(userInfo.Password) == 0 || len(userInfo.Username) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	if userInfo.Sex != common.UserSexMan && userInfo.Sex != common.UserSexWomen {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	userInfo.UserId, err = id_gen.GetId()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	err = db.Register(&userInfo)
	if err == db.ErrUserExists {
		util.ResponseError(c, util.ErrCodeUserExist)
		return
	}
	if err != nil {
		fmt.Printf("err : %#v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	util.ResponseSuccess(c, nil)
	return
}
