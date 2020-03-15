package api

import (
	"fmt"
	"net/http"
	"sre-sms-server/middleware/auth"
	"sre-sms-server/serializer"
	"sre-sms-server/tasks"

	"github.com/gin-gonic/gin"
)

func SmsSend(c *gin.Context) {
	var sms serializer.Sms
	authUser := c.MustGet(gin.AuthUserKey).(string)
	apiuser := auth.ApiUsers[authUser]
	if err := c.ShouldBind(&sms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 因为validator需要根据apiuser的配置来确定, 所以这里不能直接用go-playground的validator
	// valid phone number
	fmt.Println(sms)
	valid, possible := serializer.ValidNumber(sms.To)
	if possible == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "To is not a Possible PhoneNumber!, To 不可能是电话号码"})
		return
	}
	if apiuser.Project.StrongValid == true && valid == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "To is not a Valid PhoneNumber! To 所在区域/运营商的存在性校验失败"})
		return
	}
	tasks.SmsCreate(authUser, sms.To, sms.Content, sms.Subject, sms.SmsType, sms.Mock)

	c.JSON(200, gin.H{"message": "pong"})
}
