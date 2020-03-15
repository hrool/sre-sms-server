package serializer

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
)

type Sms struct {
	Mock    bool   `form:"mock" json:"mock"`
	SmsType string `form:"sms_type" json:"sms_type"`
	To      string `form:"to" json:"to" binding:"required"`
	Subject string `form:"Subject" json:"Subject"`
	Content string `form:"Content" json:"Content" binding:"required"`
}

func ValidNumber(to string) (bool, bool) {
	num, err := phonenumbers.Parse(to, "XX")
	if err != nil {
		fmt.Println(err)
		return false, false
	}
	return phonenumbers.IsValidNumber(num), phonenumbers.IsPossibleNumber(num)
}

func ValidSign(content string) bool {
	fmt.Println(content)
	return true
}
