package controllers

import (
	
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	//"github.com/shwinpiocess/cc/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Logout() {
	sig := this.Ctx.GetCookie("sig")
	if sig == "" {
		this.Ctx.WriteString("u'r not login")
		return
	}

	err := Logout(sig)
	if err != nil {
		this.Ctx.WriteString("logout from uic fail")
		return
	}

	this.Ctx.SetCookie("sig", "", 0, "/")
	//this.Redirect("/", 302)
	this.Redirect(Redirect_url,302)
}

//var UicInternal="http://10.78.218.100:1234"

func Logout(sig string) error {
	uri := fmt.Sprintf("%s/sso/logout/%s", UicInternal, sig)
	req := httplib.Get(uri)
	//cq0824 req.Param("token", Token)
	req.Param("token", "")
	_, err := req.String()
	return err
}


