
package controllers

import (
	//"net/url"
	//"strconv"
	"strings"
	"time"
	"fmt"

	//"github.com/astaxie/beego"
	"github.com/chenqi123/unionwarm/models"


)

type FalconGphostController struct {
	BaseController
}

func (this *FalconGphostController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *FalconGphostController)Index() {


	var fields []string
	var sortby []string
	var order []string
	var falcongroupquery map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	falcongroupquery["Create_user"] = this.CurrentUser.Name
	falcongroups, _:=models.GetFalcongroupAll(falcongroupquery, fields, sortby, order, offset, limit)
	

	var falcontplquery map[string]string = make(map[string]string)
	falcontplquery["Create_user"] = this.CurrentUser.Name

	falcontpls, _:=models.GetFalcontplAll(falcontplquery, fields, sortby, order, offset, limit)
    
	var falcongrptplquery map[string]string = make(map[string]string)
    falcongrptpl,_:=models.GetFalcongrptplAll(falcongrptplquery, fields, sortby, order, offset, limit)


    var falconclusterquery map[string]string = make(map[string]string)
	falconclusterquery["Creator"] = this.CurrentUser.Name

	falconclusters, _:=models.GetFalconclusterAll(falconclusterquery, fields, sortby, order, offset, limit)
    

    this.Data["falcongroups"] = falcongroups
    this.Data["falcontpls"] = falcontpls
    this.Data["falcongrptpl"] = falcongrptpl
    this.Data["falconclusters"] = falconclusters

 	this.TplName = "falcon/hgindex.html"
}



func (this *FalconGphostController) NewGroup() {
	this.TplName = "falcon/newgroup.html"
}


func (this *FalconGphostController) AddHosts() {
		
		grp_id,_:=this.GetInt("grp_id")
		hn:=strings.TrimSpace(this.GetString("hosts"))
		fmt.Println("###cqcq####",hn,grp_id)
		hnl:=strings.Split(hn, "\n")
		out := make(map[string]interface{})
		//var Id int64


		gh:=make(map[string]int)
		var fields []string
		var sortby []string
		var order []string
		var falconhgquery map[string]string = make(map[string]string)
		var limit int64 = 0
		var offset int64 = 0


		fields=append(fields,"Grp_id")
		fields=append(fields,"Host_id")

		hgs, _:=models.GetFalcongrphostAll(falconhgquery, fields, sortby, order, offset, limit)
		for _,v:=range hgs {
				g:=v.(map[string]interface{})["Grp_id"]
				h:=v.(map[string]interface{})["Host_id"]
				gh_key:=fmt.Sprintf("%d_%d",g,h)
				gh[gh_key]=1
		}
		
				

		for _,v:= range hnl {
			falconhost := new(models.Falconhost)
			falconhost.Hostname = v
			if models.ExistFalconhostByName(falconhost.Hostname) {
				fmt.Println("dup Hostname",falconhost.Hostname)
			} else {
				if _, err := models.FalconhostAdd(falconhost); err != nil {
					out["errInfo"] = err.Error()
					out["success"] = false
					out["errCode"] = "0006"
					this.jsonResult(out)
				} 
			}
			hid,_:=models.GetFalconhostidbyHn(v)
			falcongphost := new(models.Falcongphost)
			falcongphost.Grp_id = grp_id
			falcongphost.Host_id = hid

			gh_key1:=fmt.Sprintf("%d_%d",grp_id,hid)

			if _, ok := gh[gh_key1]; !ok {
				if _, err := models.Falcongphostbind(falcongphost); err != nil {
					out["errInfo"] = err.Error()
					out["success"] = false
					out["errCode"] = "0006"
					this.jsonResult(out)
				} 
			}


		}
		
		out["success"] = true
		this.jsonResult(out)

}



func (this *FalconGphostController) AddGroup() {
	
		falcongroup := new(models.Falcongroup)

		falcongroup.Grp_name = strings.TrimSpace(this.GetString("Grp_name"))
		falcongroup.Create_user = strings.TrimSpace(this.GetString("Create_user"))
		falcongroup.Come_from = 1	

		out := make(map[string]interface{})

		if models.ExistFalcongroupByName(falcongroup.Grp_name) {
			out["errInfo"] = "同名的模版名称已经存在！"
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		}

		var Id int64
		var err error
		if Id, err = models.FalcongroupAdd(falcongroup); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		} 

		out["success"] = true
		out["message"] = Id 
		this.jsonResult(out)

}


func (this *FalconGphostController) DelGroup() {
	out := make(map[string]interface{})
	Id, _ := this.GetInt("Id")
	if cnt, err := models.GetClusterCount(Id); err == nil {
		if cnt==0 {
			if err1 := models.DelGroup(Id); err1 == nil {
				out["success"] = true
				this.jsonResult(out)
			} else {
				out["success"] = false
				out["errInfo"] = err.Error()
				this.jsonResult(out)
			}
		} else {
			out["success"] = false
			out["errInfo"] = "请先清理自定义集群监控"
			this.jsonResult(out)
		}
		
	} else {
		out["success"] = false
		out["errInfo"] = err.Error()
		this.jsonResult(out)
	}
	
}


func (this *FalconGphostController) ShowGrpHost() {
	out := make(map[string]interface{})
	var ml []interface{}

	var fields []string
	var fields1 []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	Grp_id,_:=this.GetInt("Id")
	query["Grp_id"] = fmt.Sprintf("%d",Grp_id)
	fields=append(fields,"Host_id")

	hostids, _:=models.GetFalcongrphostAll(query, fields, sortby, order, offset, limit)

	

	for _,v:=range hostids {
		var queryid = make(map[string]string)
		hostid:=v.(map[string]interface{})["Host_id"]
		queryid["Id"]=fmt.Sprintf("%d",hostid)
		host,_:=models.GetFalconhostAll(queryid, fields1, sortby, order, offset, limit)
		ml=append(ml,host)
	}
	
	out["data"] = ml
	this.jsonResult(out)


}

func (this *FalconGphostController) BindGpHost() {
	
		falcongphost := new(models.Falcongphost)

		falcongphost.Grp_id = 4
		falcongphost.Host_id = 321299
	
		out := make(map[string]interface{})
	
		var Id int64
		var err error
		if Id, err = models.Falcongphostbind(falcongphost); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		} 

		out["success"] = true
		out["message"] = Id 
		this.jsonResult(out)


}



func (this *FalconGphostController) DelbindHosts() {

		grp_id,_:=this.GetInt("Grp_id")
		host_id,_:=this.GetInt("Id")
		
		fmt.Println("####--->grp_id,host_id",grp_id,host_id)
		//falcongphost := new(models.Falcongphost)

		//falcongphost.Grp_id = grp_id
		//falcongphost.Host_id = host_id
	
		out := make(map[string]interface{})
	
		var Id int64
		var err error
		if  Id,err = models.FalcongphostDelbind(grp_id,host_id); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		} 

		out["success"] = true
		out["message"] = Id 
		this.jsonResult(out)


}



func (this *FalconGphostController) UpdateMaintainTime() {

	out := make(map[string]interface{})
		
	falconhost := new(models.Falconhost)
	Id,_:=this.GetInt("Id")
	host_begin := strings.TrimSpace(this.GetString("host_begin"))+":00"
	host_end := strings.TrimSpace(this.GetString("host_end"))+":00"
		//fmt.Println("##host time##",host_begin,host_end)
	Maintain_begin, _ := time.Parse("2006-01-02 15:04:05", host_begin)
	Maintain_end, _ := time.Parse("2006-01-02 15:04:05", host_end)
	mb:=Maintain_begin.Unix()
	me:=Maintain_end.Unix()


		//fmt.Println("##time maintain##",Maintain_begin.Unix(),Maintain_end.Unix())
	
	falconhost.Id = Id
	falconhost.Maintain_begin= int(mb)
	falconhost.Maintain_end = int(me)
	fmt.Println("##",falconhost.Maintain_begin,falconhost.Maintain_end)


	var err error
	if err = models.UpdateMaintainTime(falconhost); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
	}
		

	var ml []interface{}

	var fields []string
	var fields1 []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	Grp_id,_:=this.GetInt("grp_id")
	query["Grp_id"] = fmt.Sprintf("%d",Grp_id)
	fields=append(fields,"Host_id")

	hostids, _:=models.GetFalcongrphostAll(query, fields, sortby, order, offset, limit)

	

	for _,v:=range hostids {
		var queryid = make(map[string]string)
		hostid:=v.(map[string]interface{})["Host_id"]
		queryid["Id"]=fmt.Sprintf("%d",hostid)
		host,_:=models.GetFalconhostAll(queryid, fields1, sortby, order, offset, limit)
		ml=append(ml,host)
	}
	
	out["data"] = ml
	this.jsonResult(out)

}


func (this *FalconGphostController) GetEndpointToHostid() {
	
		
		out := make(map[string]interface{})
		var ml =make([]string,0)

		var fields []string
		var sortby []string
		var order []string
		var query = make(map[string]string)
		var limit int64 = 0
		var offset int64 = 0

		Grp_id,_:=this.GetInt(":id")
		query["Grp_id"] = fmt.Sprintf("%d",Grp_id)
		fields=append(fields,"Host_id")

		hostids, _:=models.GetFalcongrphostAll(query, fields, sortby, order, offset, limit)

		for _,v:=range hostids {
			
			hostid:=v.(map[string]interface{})["Host_id"]
			hn,_:=models.GetFalconhostnamebyid(hostid.(int))
			str:=fmt.Sprintf(hn+"//"+fmt.Sprintf("%d",Grp_id))
			//queryid["Id"]=fmt.Sprintf("%d",hostid)
			
			ml=append(ml,str)
		}
	



		out["success"] = true
		out["data"] = ml
		this.jsonResult(out)


}

func (this *FalconGphostController) GetApiGrpName() {
	out := make(map[string]interface{})
	Grp_id,_:=this.GetInt(":id")

	gn,_:=models.GetFalcongrpnamebygid(Grp_id)

	t:=make(map[string]string)
	t[fmt.Sprintf("%d",Grp_id)]=gn
	out["data"]=t
	this.jsonResult(out)
}


func (this *FalconGphostController) AddApiHosts() {
		
		grp_id,_:=this.GetInt("grp_id")
		hn:=strings.TrimSpace(this.GetString("hosts"))
		//fmt.Println("###cqcq####",hn,grp_id)
		hnl:=strings.Split(hn, "\n")
		out := make(map[string]interface{})
		//var Id int64

		
		gh:=make(map[string]int)
		var fields []string
		var sortby []string
		var order []string
		var falconhgquery map[string]string = make(map[string]string)
		var limit int64 = 0
		var offset int64 = 0


		fields=append(fields,"Grp_id")
		fields=append(fields,"Host_id")

		hgs, _:=models.GetFalcongrphostAll(falconhgquery, fields, sortby, order, offset, limit)
		for _,v:=range hgs {
				g:=v.(map[string]interface{})["Grp_id"]
				h:=v.(map[string]interface{})["Host_id"]
				gh_key:=fmt.Sprintf("%d_%d",g,h)
				gh[gh_key]=1
		}
		
				

		for _,v:= range hnl {
			falconhost := new(models.Falconhost)
			falconhost.Hostname = v
			if models.ExistFalconhostByName(falconhost.Hostname) {
				fmt.Println("dup Hostname",falconhost.Hostname)
			} else {
				if _, err := models.FalconhostAdd(falconhost); err != nil {
					out["errInfo"] = err.Error()
					out["success"] = false
					out["errCode"] = "0006"
					this.jsonResult(out)
				} 
			}
			hid,_:=models.GetFalconhostidbyHn(v)
			falcongphost := new(models.Falcongphost)
			falcongphost.Grp_id = grp_id
			falcongphost.Host_id = hid

			gh_key1:=fmt.Sprintf("%d_%d",grp_id,hid)

			if _, ok := gh[gh_key1]; !ok {
				if _, err := models.Falcongphostbind(falcongphost); err != nil {
					out["errInfo"] = err.Error()
					out["success"] = false
					out["errCode"] = "0006"
					this.jsonResult(out)
				} 
			}


		}
		
		
		out["success"] = true
		this.jsonResult(out)

}





