package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"unionwarm/g"
	"time"
	"github.com/astaxie/beego"
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Im    string `json:"im"`
	Phone string `json:"phone"`
	Role    int       `json:"role"`
}



func GetUser(sig string) (*User, error) {
	
	key := fmt.Sprintf("u:%s", sig)
	u := g.Cache.Get(key)
	if u != nil {
		uobj := u.(User)
		return &uobj, nil
	}
	//UicInternal:="http://10.78.218.100:1234"
	UicInternal := beego.AppConfig.String("UicInternal")

	uri := fmt.Sprintf("%s/sso/user/%s", UicInternal, sig)
	req := httplib.Get(uri)
	req.Param("token", "")
	resp, err := req.Response()

	//fmt.Println("#####resp",resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	type TmpStruct struct {
		User *User `json:"user"`
	}
	var t TmpStruct

	err = decoder.Decode(&t)
	if err != nil {
		return nil, err
	}

	// don't worry cache expired. we just use username which can not modify
	g.Cache.Put(key, *t.User, time.Duration(360000))

	return t.User, nil
}
