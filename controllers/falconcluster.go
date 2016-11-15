package controllers

import (
	//"net/url"
	//"strconv"
	"strings"
	//"time"
	//"fmt"

	//"github.com/astaxie/beego"
	"unionwarm/models"


)

type FalconClusterController struct {
	BaseController
}

func (this *FalconClusterController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *FalconClusterController)Index() {


	var fields []string
	var sortby []string
	var order []string
	var falcongroupquery map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	falcongroupquery["Create_user"] = this.CurrentUser.Name
	falcongroups, _:=models.GetFalcongroupAll(falcongroupquery, fields, sortby, order, offset, limit)
	

	var falconclusterquery map[string]string = make(map[string]string)
	falconclusterquery["Creator"] = this.CurrentUser.Name

	falconclusters, _:=models.GetFalconclusterAll(falconclusterquery, fields, sortby, order, offset, limit)
    
	this.Data["falcongroups"] = falcongroups
    this.Data["falconclusters"] = falconclusters
   
 	this.TplName = "falcon/clsindex.html"
}

func (this *FalconClusterController) NewCls() {
	var fields []string
	var sortby []string
	var order []string
	var falcongroupquery map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	falcongroupquery["Create_user"] = this.CurrentUser.Name
	falcongroups, _:=models.GetFalcongroupAll(falcongroupquery, fields, sortby, order, offset, limit)
	this.Data["falcongroups"] = falcongroups
	this.TplName = "falcon/newcls.html"
}


func (this *FalconClusterController) AddCls() {
	grp_id,_:=this.GetInt("Grp_id")
	numerator:=strings.TrimSpace(this.GetString("Numerator"))
	denominator:=strings.TrimSpace(this.GetString("Denominator"))
	endpoint:=strings.TrimSpace(this.GetString("Endpoint"))
	metric:=strings.TrimSpace(this.GetString("Metric"))
	tags:=strings.TrimSpace(this.GetString("Tags"))
	ds_type:="GAUGE"
	step:=60
	creator:=this.CurrentUser.Name

	out := make(map[string]interface{})

	falconcluster := new(models.Falconcluster)
	falconcluster.Grp_id = grp_id
	falconcluster.Numerator = numerator
	falconcluster.Denominator = denominator
	falconcluster.Endpoint = endpoint
	falconcluster.Metric = metric
	falconcluster.Tags = tags
	falconcluster.Ds_type = ds_type
	falconcluster.Step = step
	falconcluster.Creator = creator

	if models.ExistFalconclusterByName(falconcluster.Endpoint,falconcluster.Metric,falconcluster.Tags) {
				out["success"] = false
				out["errCode"] = "0006"
				out["errInfo"] = "设备名称、指标、标签重复"
				this.jsonResult(out)
	} else {
		if _, err := models.FalconclusterAdd(falconcluster); err != nil {
		out["errInfo"] = err.Error()
		out["success"] = false
		out["errCode"] = "0006"
		this.jsonResult(out)
		}
	}

	out["success"] = true
	this.jsonResult(out)
}

func (this *FalconClusterController) DelCls() {
	out := make(map[string]interface{})
	Id, _ := this.GetInt("Id")
	if err := models.DelCls(Id); err == nil {
		out["success"] = true
		this.jsonResult(out)
	} else {
		out["success"] = false
		out["errInfo"] = err.Error()
		this.jsonResult(out)
	}
}


func (this *FalconClusterController) EditCls() {
	out := make(map[string]interface{})
	algorithm:=strings.TrimSpace(this.GetString("algorithm"))
	algorithm_arr:=strings.Split(algorithm,"//")
	falconcluster := new(models.Falconcluster)
	falconcluster.Id,_ = this.GetInt("id")
	falconcluster.Numerator = algorithm_arr[0]
	falconcluster.Denominator = algorithm_arr[1]
	falconcluster.Endpoint = strings.TrimSpace(this.GetString("endpoint"))
	falconcluster.Metric = strings.TrimSpace(this.GetString("metric"))
	falconcluster.Tags = strings.TrimSpace(this.GetString("tags"))

	if models.ExistFalconclusterByName(falconcluster.Endpoint,falconcluster.Metric,falconcluster.Tags) {
				out["success"] = false
				out["errCode"] = "0006"
				out["errInfo"] = "设备名称、指标、标签重复"
				this.jsonResult(out)
	} else {
		var err error
		if err = models.EditCls(falconcluster); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		} 
	}
	
	out["success"] = true
	this.jsonResult(out)

}


