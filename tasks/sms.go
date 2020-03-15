package tasks

func SmsCreate(auth_user string, to string, content string, subject string, sms_type string, mock bool) {
	println(auth_user, to, content, subject, sms_type, mock)
}
