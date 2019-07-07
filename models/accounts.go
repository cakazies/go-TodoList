package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	util "github.com/local/TaskListGo/util"
)

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"` //skipping this field does not work in mysql
}

func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return util.MetaMsg(false, "Email address format is incorrect"), false
	}

	if len(account.Password) < 6 {
		return util.MetaMsg(false, "Password is minimum 6 character"), false
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return util.MetaMsg(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return util.MetaMsg(false, "Email address already in use by another user."), false
	}

	return util.MetaMsg(true, "Requirement passed"), true
}

// func (account *Account) CreateAccount() map[string]interface{} {
// 	if rsp, status := account.Validate(); !status {
// 		return rsp
// 	}

// 	//TODO
// 	//1. generate hashedPassword from user plaintext password
// 	//use bcrypt.GenerateFromPassword

// 	//2. create account of new user

// 	//3. return success response
// }

// func (account *Account) Login() map[string]interface{} {
// 	//TODO
// 	//1. get account from DB with specified email
// 	//return if email is not found

// 	//2. compared the registered password hash with given password hash
// 	//use bcrypt.CompareHashAndPassword
// 	//return is there is mismatch

// 	//3. return success response
// }
