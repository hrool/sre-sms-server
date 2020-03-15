package db

type Project struct {
	ID           uint
	Name         string
	Enable       bool
	Note         string
	TimeoutRetry bool
	Timeout      uint
	FailRetry    bool
	StrongValid  bool
}

type Apiuser struct {
	ID        uint
	ProjectID uint
	Project   Project
	Username  string
	Password  string
	Enable    bool
	Note      string
}

type Sms struct {
	ID        uint
	ApiuserID uint
	Apiuser   Apiuser
	Mock      bool
	SmsType   string
	To        string
	Subject   string
	Content   string
}

func GetApiUsers() ([]Apiuser, error) {
	var apiusers []Apiuser
	err := Conn.Where(&Apiuser{Enable: true}).Preload("Project").Find(&apiusers).Error
	return apiusers, err
}
