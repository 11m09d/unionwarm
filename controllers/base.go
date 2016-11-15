package controllers

import (
	//"net/url"
	//"strconv"
	"strings"
	"time"
	"fmt"

	"github.com/astaxie/beego"

	"unionwarm/models"
	//"github.com/shwinpiocess/cc/utils"
	"github.com/astaxie/beego/httplib"
)

type BaseController struct {
	beego.Controller

	userId         int
	userName       string
	
	requestPath string
	today string
	
	firstApp bool
	firstSet bool
	firtModule bool

	//cq0824
	CurrentUser *models.User
	
	//defaultApp *models.App
	//defaultCmdbApp *models.CmdbApp
}

func (this *BaseController) Prepare() {
	this.requestPath = this.Ctx.Input.URL()
	this.today = time.Now().Format("20060102")
	
	//this.auth()
	/*
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	query["owner_id"] = strconv.Itoa(this.userId)
	query["default"] = strconv.FormatBool(false)


	//fmt.Println("##control##base##query:##",query,"##fields",fields,"##sortby",sortby,"##order",order,"##offset",offset,"##limit",limit)
	//apps, _ := models.GetAllApp(query, fields, sortby, order, offset, limit)
	//fmt.Println("###apps",apps)


	var cmdbquery map[string]string = make(map[string]string)
	var cmdbhostquery map[string]interface{} = make(map[string]interface{})

	cmdbapps, _:=models.GetCmdbAllApp(cmdbquery, fields, sortby, order, offset, limit)
	cmdbhosts,_:=models.GetCmdbAllHost(cmdbhostquery,fields, sortby, order, offset, limit)
	//fmt.Println("###cmdbapps",cmdbapps)
	cnt,_:= models.GetCmdbAppCount()
	apps_cnt := fmt.Sprintf("%d", cnt-1)

	cnt1,_:=models.GetCmdbHostAllCount()
	fmt.Println("##cnt1-->",cnt1)
	hosts_cnt:=fmt.Sprintf("%d",cnt1)

	cnt2,_:=models.GetCmdbHostCount(0,"ApplicationID")
	free_hosts_cnt:=fmt.Sprintf("%d",cnt2)	

	var cmdbhostquery_all map[string]interface{} = make(map[string]interface{})
	var fields_all []string
	fields_all=append(fields_all,"Cpunum")
	fields_all=append(fields_all,"Memsize")
	fields_all=append(fields_all,"Disksize")
	fields_all=append(fields_all,"Cputype")
	fields_all=append(fields_all,"OSName")
	fields_all=append(fields_all,"Product")
	fields_all=append(fields_all,"Cabinets")
	
	cmdbhosts_all,_:=models.GetCmdbAllHost(cmdbhostquery_all,fields_all, sortby, order, offset, limit)

	//0816
	cpusum:=0
	var memsum float64
	var disksum float64
	cputype_map:=make(map[string]int)
	osname_map:=make(map[string]int)
	product_map:=make(map[string]int)
	cabinets_map:=make(map[string]int)

	for _,ahost := range cmdbhosts_all {
		cpusum=cpusum+ahost.(map[string]interface{})["Cpunum"].(int)

		m:=fmt.Sprintf("%s",ahost.(map[string]interface{})["Memsize"])
		m1:=strings.Replace(m," GB","",-1)
		m2:=strings.Replace(m1,"GB","",-1)
		if m2=="" {
			m2="0"
		}
		m3,err:=strconv.ParseFloat(m2,64)
		if err!=nil {
			fmt.Println(err)
		}
		memsum=memsum+m3


		d:=fmt.Sprintf("%s",ahost.(map[string]interface{})["Disksize"])
		d1:=strings.Replace(d," GB","",-1)
		d2:=strings.Replace(d1,"GB","",-1)
		if d2=="" {
			d2="0"
		} 
		d3,err:=strconv.ParseFloat(d2,64)
		if err!=nil {
			fmt.Println(err)
		}
		disksum=disksum+d3


		cputype:=fmt.Sprintf("%s",ahost.(map[string]interface{})["Cputype"])
		if cputype_val,ok:=cputype_map[cputype]; ok {
			cputype_map[cputype]=cputype_val+1
		} else {
			cputype_map[cputype]=1
		}

		osname:=fmt.Sprintf("%s",ahost.(map[string]interface{})["OSName"])
		if osname_val,ok:=osname_map[osname]; ok {
			osname_map[osname]=osname_val+1
		} else if len(osname)==0 {
			fmt.Println(osname)
		} else {
			osname_map[osname]=1
		}

		product:=fmt.Sprintf("%s",ahost.(map[string]interface{})["Product"])
		if product_val,ok:=product_map[product]; ok {
			product_map[product]=product_val+1
		} else {
			product_map[product]=1
		}

		cabinets:=fmt.Sprintf("%s",ahost.(map[string]interface{})["Cabinets"])
		if cabinets_val,ok:=cabinets_map[cabinets]; ok {
			cabinets_map[cabinets]=cabinets_val+1
		} else {
			cabinets_map[cabinets]=1
		}




	}

	memsum_web:=fmt.Sprintf("%.0f",memsum)
	disksum_web:=fmt.Sprintf("%.0f",disksum/1024)



	var cmdbquery_app map[string]string = make(map[string]string)
	var fields_app []string
	fields_app=append(fields_app,"Owner")
	

	cmdbapps_app, _:=models.GetCmdbAllApp(cmdbquery_app, fields_app, sortby, order, offset, limit)
	owner_map:=make(map[string]bool)
	for _,one_app := range cmdbapps_app {
		owner:=fmt.Sprintf("%s",one_app.(map[string]interface{})["Owner"])
		if _,ok:=owner_map[owner]; ok {
			continue
		} else {
			owner_map[owner]=true
		}

	}

	module_cnt,_:= models.GetCmdbModuleCount()
	


	if len(cmdbapps) > 0 {
		defaultAppId := this.Ctx.GetCookie("defaultAppId")
		//fmt.Println("####defaultAPPID:###",defaultAppId)
		
		if defaultAppId, err := strconv.Atoi(defaultAppId); err == nil {
			if app, err := models.GetCmdbAppById(defaultAppId); err != nil {
				defaultCmdbApp := cmdbapps[0].(models.CmdbApp)
				this.defaultCmdbApp = &defaultCmdbApp
				//fmt.Println("###itoa defaultappid:",strconv.Itoa(this.defaultApp.Id))
				this.Ctx.SetCookie("defaultAppId", strconv.Itoa(this.defaultCmdbApp.Id))
				this.Ctx.SetCookie("defaultAppName", url.QueryEscape(this.defaultCmdbApp.ApplicationName))
			} else {
				this.defaultCmdbApp = app
			}
		} else {
			defaultCmdbApp := cmdbapps[0].(models.CmdbApp)
			this.defaultCmdbApp = &defaultCmdbApp
			//fmt.Println("##2#itoa defaultappid:",strconv.Itoa(this.defaultApp.Id))

			this.Ctx.SetCookie("defaultAppId", strconv.Itoa(this.defaultCmdbApp.Id))
			this.Ctx.SetCookie("defaultAppName", url.QueryEscape(this.defaultCmdbApp.ApplicationName))
		}
	} else {
		this.firstApp = true
		this.firstSet = true
		this.firtModule = true
	}

	

	this.Data["defaultApp"] = this.defaultCmdbApp
	this.Data["apps"] = cmdbapps
	this.Data["cmdbhosts"] = cmdbhosts
	this.Data["appscnt"] = apps_cnt
	this.Data["hostscnt"] =hosts_cnt
	this.Data["freehostscnt"]= free_hosts_cnt
	this.Data["firstApp"] = this.firstApp
	this.Data["firstSet"] = this.firstSet
	this.Data["firstModule"] = this.firtModule
	this.Data["cpusum"] = cpusum
	this.Data["memsum"] = memsum_web
	this.Data["disksum"] = disksum_web
	this.Data["ownersum"] = len(owner_map)
	this.Data["cputypesum"] = len(cputype_map)
	this.Data["osnamesum"] = len(osname_map)
	this.Data["productsum"] = len(product_map)
	this.Data["cabinetssum"] = len(cabinets_map)
	this.Data["module_cnt"] = module_cnt

	*/

	
	
	//if !strings.Contains(this.requestPath, "/zjz_work_cmdb/cmdb/api/action") && !strings.Contains(this.requestPath, "/zjz_work_cmdb/cmdb/api/eptohid") {
	if !strings.Contains(this.requestPath, "/zjz_work_cmdb/cmdb/api/") {
	
		this.CheckLogin()	
		this.Data["user"] = this.CurrentUser.Name
		this.Data["requestPath"] = this.requestPath
		this.Data["today"] = this.today

	} else {
		this.Data["requestPath"] = this.requestPath
		this.Data["today"] = this.today

	}
	/*
	this.CheckLogin()	
	this.Data["user"] = this.CurrentUser.Name
	this.Data["requestPath"] = this.requestPath
	this.Data["today"] = this.today
	*/

	
}

func (this *BaseController) getClientIP() string {
	return strings.Split(this.Ctx.Request.RemoteAddr, ":")[0]
}

func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

func (this *BaseController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) auth()(uname, auth interface{}){
	uname = this.GetSession("uname")
	auth = this.GetSession("auth")
	if uname == nil {
		this.Redirect("/zjz_work_cmdb/cmdb/login", 302)
		return
	}
	return
}

//cq0824

func (this *BaseController) hasBusinessCookie() bool {
	if this.CurrentUser != nil {
		return true
	}

	sigInCookie := this.Ctx.GetCookie("sig")
	if sigInCookie != "" {
		u, err := models.GetUser(sigInCookie)
		if err != nil || u == nil {
			return false
		}

		this.CurrentUser = u
		
		return true
	}

	return false
}


//var UicInternal="http://10.78.218.100:1234"
var UicInternal = beego.AppConfig.String("UicInternal")
var Redirect_url = beego.AppConfig.String("Redirect_url")



func (this *BaseController) CheckLogin() {
	
	if this.hasBusinessCookie() {
		this.Data["CurrentUser"] = this.CurrentUser
		return
	}

	sig, err := GetSig()
	if err != nil {
		//this.ServeErrJSON(fmt.Sprintf("curl get sig fail: %v", err))
		panic(err)
		return
	}

	this.Ctx.SetCookie("sig", sig, 2592000, "/")
	this.Redirect(fmt.Sprintf("%s/auth/login?sig=%s&callback=http://%s:%d%s", UicInternal, sig, this.Ctx.Input.Host(), this.Ctx.Input.Port(), this.Ctx.Input.URI()), 302)
}


func GetSig() (sig string, err error) {
	uri := fmt.Sprintf("%s/sso/sig", UicInternal)
	req := httplib.Get(uri)
	sig, err = req.String()
	return
}


















/*
func (this *BaseController) auth() {
	arrs := strings.Split(this.Ctx.GetCookie("auth"), "|")
	if len(arrs) == 2 {
		idstr, password := arrs[0], arrs[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := models.GetUserById(userId)
			if err == nil && password == utils.Md5([]byte(this.getClientIP()+"|"+user.Password+user.Salt)) {
				this.userId = user.Id
				this.userName = user.UserName
			}
		}
	}
	
	if this.userId == 0 && this.requestPath != beego.URLFor("MainController.Login") && this.requestPath != beego.URLFor("MainController.Logout") {
		this.redirect(beego.URLFor("MainController.Login"))
	}
	
}
*/
