package auth

import (
	"encoding/base64"
	"fmt"

	"net/http"
	"sre-sms-server/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthUserKey is the cookie name for user credential in basic auth.
const AuthUserKey = "user"

// Accounts defines a key/value for user/pass list of authorized logins.
type Accounts map[string]string

type authPair struct {
	value string
	user  string
}

type authPairs []authPair

func (a authPairs) searchCredential(authValue string) (string, bool) {
	if authValue == "" {
		return "", false
	}
	for _, pair := range a {
		if pair.value == authValue {
			return pair.user, true
		}
	}
	return "", false
}

func assert1(guard bool, text string) {
	if !guard {
		panic(text)
	}
}

var dbAuthPairs authPairs

// BasicAuth returns a Basic HTTP Authorization middleware. use dbAuthPairs to verify user and password.
func BasicAuth() gin.HandlerFunc {
	realm := "Basic realm=" + strconv.Quote("Authorization Required")
	return func(c *gin.Context) {
		// Search user in the slice of allowed credentials

		user, found := dbAuthPairs.searchCredential(c.Request.Header.Get("Authorization"))
		if !found {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
		// c.MustGet(gin.AuthUserKey).
		c.Set(AuthUserKey, user)
	}
}

func processAccounts(accounts Accounts) authPairs {
	assert1(len(accounts) > 0, "Empty list of authorized credentials")
	pairs := make(authPairs, 0, len(accounts))
	for user, password := range accounts {
		assert1(user != "", "User can not be empty")
		value := authorizationHeader(user, password)
		pairs = append(pairs, authPair{
			value: value,
			user:  user,
		})
	}
	return pairs
}

func authorizationHeader(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
}

// LoadAuthData to load user,password and generate authpairs for BasicAuth
// todo 后面需要区分第一次运行和计划任务运行, 计划运行的时候 不能panic, 允许出错.
var ApiUsers map[string]db.Apiuser

func LoadAuthData() {
	dbAccounts := Accounts{}
	mapApiusers := make(map[string]db.Apiuser)
	apiusers, err := db.GetApiUsers()
	if err != nil {
		panic(err)
	}
	for _, apiuser := range apiusers {
		if apiuser.Project.Enable == false {
			continue
		}
		dbAccounts[apiuser.Username] = apiuser.Password
		mapApiusers[apiuser.Username] = apiuser
	}
	fmt.Println(dbAccounts)
	pairs := processAccounts(dbAccounts)
	dbAuthPairs = pairs
	ApiUsers = mapApiusers
	fmt.Println(ApiUsers)
}
