
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

type FalconTplController struct {
	BaseController
}

func (this *FalconTplController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}
 
func (this *FalconTplController)Index() {

	var fields []string
	var sortby []string
	var order []string
	var falcontplquery map[string]string = make(map[string]string)
	var falconuicquery map[string]string = make(map[string]string)
	var falconactionquery map[string]string = make(map[string]string)
	
	var limit int64 = 0
	var offset int64 = 0

	falcontplquery["Create_user"] = this.CurrentUser.Name
	//fmt.Println("######test",this.CurrentUser)
	falcontpls, _:=models.GetFalcontplAll(falcontplquery, fields, sortby, order, offset, limit)
	falconuics, _:=models.GetFalconuicAll(falconuicquery, fields, sortby, order, offset, limit)
	falconactions, _:=models.GetFalconactionAll(falconactionquery, fields, sortby, order, offset, limit)

	/*
	out := make(map[string]interface{})
    i,err := models.GetFalcontplCount()
    is:=fmt.Sprintf("%d",i)
    if err == nil {
        out["success"] = true
		out["message"] = is
		this.jsonResult(out)
     }
     */
    this.Data["falcontpls"] = falcontpls
    this.Data["falconuics"] = falconuics
    this.Data["falconactions"] = falconactions
    


    
	//fmt.Println(this.CurrentUser)
	//fmt.Println("##Role",this.CurrentUser.Role)
	/*
	if this.CurrentUser.Role == 0 {
		this.TplName = "cmdb/zjzImport_guest.html"
	}	else if this.CurrentUser.Role == 1 {
		this.TplName = "cmdb/zjzImport_owner.html"
	} else {
		this.TplName = "cmdb/zjzImport.html"
	}
	*/

    
 	this.TplName = "falcon/index.html"
}

func (this *FalconTplController) NewTpl() {
	this.TplName = "falcon/newtpl.html"
}



func (this *FalconTplController) UcintfIndex() {
	this.TplName = "falcon/ucintf.html"
}
func (this *FalconTplController) UcintfAdd() {
	this.TplName = "falcon/ucintfadd.html"
}

func (this *FalconTplController) NewStrategy() {
	this.TplName = "falcon/newstrategy.html"
}



func (this *FalconTplController) AddTpl() {
	if this.isPost() {
		falcontpl := new(models.Falcontpl)

		falcontpl.Tpl_name = strings.TrimSpace(this.GetString("Tpl_name"))
		falcontpl.Create_user = strings.TrimSpace(this.GetString("Create_user"))
		falcontpl.Create_at = time.Now()
		

		out := make(map[string]interface{})

		if models.ExistFalcontplByName(falcontpl.Tpl_name) {
			out["errInfo"] = "同名的模版名称已经存在！"
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		}


		var Id int64
		var err error
		if Id, err = models.FalcontplAdd(falcontpl); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		} 

		out["success"] = true
		out["message"] = Id 
		this.jsonResult(out)
	}
}

func (this *FalconTplController) DelTpl() {
	out := make(map[string]interface{})
	Id, _ := this.GetInt("Id")
	if err := models.DelTpl(Id); err == nil {
		out["success"] = true
		this.jsonResult(out)
	} else {
		out["success"] = false
		out["errInfo"] = err.Error()
		this.jsonResult(out)
	}
}

func (this *FalconTplController) DelStrategy() {
	out := make(map[string]interface{})
	Id, _ := this.GetInt("Id")
	if err := models.DelStrategy(Id); err == nil {
		out["success"] = true
		this.jsonResult(out)
	} else {
		out["success"] = false
		out["errInfo"] = err.Error()
		this.jsonResult(out)
	}
}


func (this *FalconTplController) ShowStrategy() {
	out := make(map[string]interface{})
	
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	Tpl_id,_:=this.GetInt("Id")
	query["Tpl_id"] = fmt.Sprintf("%d",Tpl_id)
	
	strategyall, _:=models.GetFalconstrategyAll(query, fields, sortby, order, offset, limit)
	
	out["data"] = strategyall
	this.jsonResult(out)

}


func (this *FalconTplController) AddStrategy() {
	out := make(map[string]interface{})

	sid,_:=this.GetInt("sid")
	if sid==0 {
		falconstrategy := new(models.Falconstrategy)
	//falconstrategy.Id,_ = this.GetInt("sid")
	falconstrategy.Metric = strings.TrimSpace(this.GetString("metric"))
	falconstrategy.Tags = strings.TrimSpace(this.GetString("tags"))
	falconstrategy.Max_step,_ = this.GetInt("max_step")
	falconstrategy.Priority,_ = this.GetInt8("priority")
	falconstrategy.Func = strings.TrimSpace(this.GetString("func"))
	falconstrategy.Op = strings.TrimSpace(this.GetString("op"))
	falconstrategy.Right_vaue = strings.TrimSpace(this.GetString("right_value"))
	falconstrategy.Note = strings.TrimSpace(this.GetString("note"))
	falconstrategy.Run_begin = strings.TrimSpace(this.GetString("run_begin"))
	falconstrategy.Run_end = strings.TrimSpace(this.GetString("run_end"))
	falconstrategy.Tpl_id,_ = this.GetInt("tpl_id")

	var Id int64
	var err error
	if Id, err = models.FalconstrategyAdd(falconstrategy); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
	} 

	out["success"] = true
	out["message"] = Id 
	

	} else {
		falconstrategy := new(models.Falconstrategy)
	falconstrategy.Id,_ = this.GetInt("sid")
	falconstrategy.Metric = strings.TrimSpace(this.GetString("metric"))
	falconstrategy.Tags = strings.TrimSpace(this.GetString("tags"))
	falconstrategy.Max_step,_ = this.GetInt("max_step")
	falconstrategy.Priority,_ = this.GetInt8("priority")
	falconstrategy.Func = strings.TrimSpace(this.GetString("func"))
	falconstrategy.Op = strings.TrimSpace(this.GetString("op"))
	falconstrategy.Right_vaue = strings.TrimSpace(this.GetString("right_value"))
	falconstrategy.Note = strings.TrimSpace(this.GetString("note"))
	falconstrategy.Run_begin = strings.TrimSpace(this.GetString("run_begin"))
	falconstrategy.Run_end = strings.TrimSpace(this.GetString("run_end"))
	//falconstrategy.Tpl_id,_ = this.GetInt("tpl_id")

	var err error
	if err = models.FalconstrategyUpdate(falconstrategy); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
	} 

	out["success"] = true
	

	}
	this.jsonResult(out)
	
}


func (this *FalconTplController) ShowGrpTpl() {
	out := make(map[string]interface{})
	var ml []interface{}

	var fields []string
	var fields1 []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0

	Tpl_id,_:=this.GetInt("Id")
	query["Tpl_id"] = fmt.Sprintf("%d",Tpl_id)
	fields=append(fields,"Grp_id")

	grpids, _:=models.GetFalcongrptplAll(query, fields, sortby, order, offset, limit)

	

	for _,v:=range grpids {
		var queryid = make(map[string]string)
		grpid:=v.(map[string]interface{})["Grp_id"]
		queryid["Id"]=fmt.Sprintf("%d",grpid)
		grp,_:=models.GetFalcongroupAll(queryid, fields1, sortby, order, offset, limit)
		ml=append(ml,grp)
	}
	
	out["data"] = ml
	this.jsonResult(out)


}




func (this *FalconTplController) ShowTplGrp() {
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
	fields=append(fields,"Tpl_id")

	tplids, _:=models.GetFalcongrptplAll(query, fields, sortby, order, offset, limit)

	

	for _,v:=range tplids {
		var queryid = make(map[string]string)
		tplid:=v.(map[string]interface{})["Tpl_id"]
		queryid["Id"]=fmt.Sprintf("%d",tplid)
		tpl,_:=models.GetFalcontplAll(queryid, fields1, sortby, order, offset, limit)
		ml=append(ml,tpl)
	}
	
	out["data"] = ml
	this.jsonResult(out)


}


func (this *FalconTplController) QueryTpl() {

	//query1:= strings.TrimSpace(this.GetString("query"))
	//fmt.Println("###...query1###",query1)
	var ml []interface{}
	
	var fields []string
	var sortby []string
	var order []string
	var falcontplquery map[string]string = make(map[string]string)
	var limit int64 = 0
	var offset int64 = 0
	//falcontplquery["Tpl_name"]=query1

	fields=append(fields,"Id")
	fields=append(fields,"Tpl_name")
	fields=append(fields,"Parent_id")
	fields=append(fields,"Action_id")
	fields=append(fields,"Create_user")
	

	falcontpls, _:=models.GetFalcontplAll(falcontplquery, fields, sortby, order, offset, limit)
	
	for _,v:=range falcontpls {
		m:=make(map[string]interface{})
		name:=fmt.Sprintf("%s",v.(map[string]interface{})["Tpl_name"])
		id:=v.(map[string]interface{})["Id"]
		parent_id:=v.(map[string]interface{})["Parent_id"]
		action_id:=v.(map[string]interface{})["Action_id"]
		create_user:=fmt.Sprintf("%s",v.(map[string]interface{})["Create_user"])
		m["name"]=name
		m["id"]=id
		m["parent_id"]=parent_id
		m["action_id"]=action_id
		m["create_user"]=create_user
		ml=append(ml,m)
	}
	

	
	out := make(map[string]interface{})
	/*
    i,err := models.GetFalcontplCount()
    is:=fmt.Sprintf("%d",i)
    if err == nil {
        out["success"] = true
		out["message"] = is
		this.jsonResult(out)
     }
    */
    out["data"]=ml
    this.jsonResult(out)
    
}


func (this *FalconTplController) BindGrpTpl() {

		gt:=make(map[string]int)
		var fields []string
		var sortby []string
		var order []string
		var falconhgquery map[string]string = make(map[string]string)
		var limit int64 = 0
		var offset int64 = 0


		fields=append(fields,"Grp_id")
		fields=append(fields,"Tpl_id")

		gts, _:=models.GetFalcongrptplAll(falconhgquery, fields, sortby, order, offset, limit)
		for _,v:=range gts {
				g:=v.(map[string]interface{})["Grp_id"]
				t:=v.(map[string]interface{})["Tpl_id"]
				gt_key:=fmt.Sprintf("%d_%d",g,t)
				gt[gt_key]=1
		}

		falcongptpl := new(models.Falcongptpl)
		grp_id,_:=this.GetInt("grp_id")
		tpl_id,_:=this.GetInt("tpl_id")

		falcongptpl.Grp_id = grp_id
		falcongptpl.Tpl_id = tpl_id
	
		out := make(map[string]interface{})
	
		var Id int64
		var err error

		gt_key1:=fmt.Sprintf("%d_%d",grp_id,tpl_id)

		if _, ok := gt[gt_key1]; !ok {
			if Id, err = models.Falcongrptplbind(falcongptpl); err != nil {
				out["errInfo"] = err.Error()
				out["success"] = false
				out["errCode"] = "0006"
				this.jsonResult(out)
			}
			out["success"] = true
			out["message"] = Id 
		} 

		
		this.jsonResult(out)


}


func (this *FalconTplController) DelGrpTpl() {

		grp_id,_:=this.GetInt("Grp_id")
		tpl_id,_:=this.GetInt("Tpl_id")
		
		fmt.Println("####--->grp_id,tpl_id",grp_id,tpl_id)
		//falcongphost := new(models.Falcongphost)

		//falcongphost.Grp_id = grp_id
		//falcongphost.Host_id = host_id
	
		out := make(map[string]interface{})
	
		var Id int64
		var err error
		if  Id,err = models.FalcongrptplDelbind(grp_id,tpl_id); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
		} 

		out["success"] = true
		out["message"] = Id 
		this.jsonResult(out)


}






func (this *FalconTplController) UpdateParenttpl() {


		falcontpl := new(models.Falcontpl)
		parent_tpl_id,_:=this.GetInt("parent_tpl_id")
		tpl_id,_:=this.GetInt("tpl_id")

		
		falcontpl.Id = tpl_id
		falcontpl.Parent_id = parent_tpl_id
		
	
		out := make(map[string]interface{})
		var err error

		if err = models.UpdateParentid(falcontpl); err != nil {
				out["errInfo"] = err.Error()
				out["success"] = false
				out["errCode"] = "0006"
				this.jsonResult(out)
		}
		
		out["success"] = true
		this.jsonResult(out)
}



func (this *FalconTplController) GetActionCount() {
		out := make(map[string]interface{})
		var id int64
		var err error
		if id,err = models.GetFalconActionCount(); err != nil {
				out["errInfo"] = err.Error()
				out["success"] = false
				out["errCode"] = "0006"
				this.jsonResult(out)
		}
		
		out["success"] = true
		out["message"] = id
		this.jsonResult(out)
}


func (this *FalconTplController) GetUicCount() {
		out := make(map[string]interface{})
		var id int64
		var err error
		if id,err = models.GetFalconUicCount(); err != nil {
				out["errInfo"] = err.Error()
				out["success"] = false
				out["errCode"] = "0006"
				this.jsonResult(out)
		}
		
		out["success"] = true
		out["message"] = id
		this.jsonResult(out)
}

func (this *FalconTplController) GetActionDetail() {
		out := make(map[string]interface{})
		tpl_id,_:=this.GetInt("tpl_id")

		var id int
		var err error
		if id,err = models.GetFalconactidById(tpl_id); err != nil {
				out["errInfo"] = err.Error()
				out["success"] = false
				out["errCode"] = "0006"
				this.jsonResult(out)
		}

		if id!=0 {
			var fields []string
			var sortby []string
			var order []string
			var falconactionquery map[string]string = make(map[string]string)
			var limit int64 = 0
			var offset int64 = 0
			falconactionquery["Id"] = fmt.Sprintf("%d",id)
			actions, _:=models.GetFalconactionAll(falconactionquery, fields, sortby, order, offset, limit)
			out["success"] = true
			out["message"] = actions
		} else {
			out["success"] = true
			out["message"] = id
		}
		
		this.jsonResult(out)
}


func (this *FalconTplController) ModifyAction() {
		out := make(map[string]interface{})
		
		tpl_id,_:=this.GetInt("tpl_id")
		action_id,_:=this.GetInt("action_id")
		uic := strings.TrimSpace(this.GetString("uic"))
		callback_url := strings.TrimSpace(this.GetString("callback_url"))
		callback,_:=this.GetInt("callback")
		after_callback_mail,_:=this.GetInt("after_callback_mail")
		after_callback_sms,_:=this.GetInt("after_callback_sms")
		before_callback_mail,_:=this.GetInt("before_callback_mail")
		before_callback_sms,_:=this.GetInt("before_callback_sms")


		if action_id==0 {
			falconaction := new(models.Falconaction)
			falconaction.Uic = uic
			falconaction.Url = callback_url
			falconaction.Callback = int8(callback)
			falconaction.Before_callback_sms = int8(before_callback_sms)
			falconaction.Before_callback_mail = int8(before_callback_mail)
			falconaction.After_callback_sms = int8(after_callback_sms)
			falconaction.After_callback_mail = int8(after_callback_mail)

			var Id int64
			var err error
			if Id, err = models.FalconactionAdd(falconaction); err != nil {
				out["errInfo"] = err.Error()
				out["success"] = false
				out["errCode"] = "0006"
				this.jsonResult(out)
			} 


			falcontpl := new(models.Falcontpl)
			falcontpl.Id = tpl_id
			falcontpl.Action_id = int(Id)

			if err = models.UpdateActionid(falcontpl); err != nil {
				out["errInfo"] = err.Error()
				out["success"] = false
				out["errCode"] = "0006"
				this.jsonResult(out)
			}
		} else {
			falconaction := new(models.Falconaction)
			falconaction.Id = action_id
			falconaction.Uic = uic
			falconaction.Url = callback_url
			falconaction.Callback = int8(callback)
			falconaction.Before_callback_sms = int8(before_callback_sms)
			falconaction.Before_callback_mail = int8(before_callback_mail)
			falconaction.After_callback_sms = int8(after_callback_sms)
			falconaction.After_callback_mail = int8(after_callback_mail)

			
			var err error
			if err = models.UpdateAction(falconaction); err != nil {
				out["errInfo"] = err.Error()
				out["success"] = false
				out["errCode"] = "0006"
				this.jsonResult(out)
			} 


		}
		
		out["success"] = true
		this.jsonResult(out)


}


func (this *FalconTplController) GetTplDetail() {
		out := make(map[string]interface{})
		//tpl_id,_:=this.GetInt(":id")
		tpl_id,_:=this.GetInt(":id")
		//fmt.Println("##act_id#",act_id)
		tp,err:=models.GetFalcontplById(tpl_id)
		if err!=nil {
			out["error"]=err
			this.jsonResult(out)
		}
		
		if tp.Action_id != 0 {
			action,_:=models.GetFalconActionById(tp.Action_id)	
			out["action"]=action
		}

		var fields []string
		var sortby []string
		var order []string
		var query = make(map[string]string)
		var limit int64 = 0
		var offset int64 = 0


		if tp.Parent_id !=0 {
			parent,_:=models.GetFalcontplById(tp.Parent_id)	

			query["Tpl_id"] = fmt.Sprintf("%d",tp.Parent_id)
			p_strategyall, _:=models.GetFalconstrategyAll(query, fields, sortby, order, offset, limit)
			
			out["Parent"]=parent
			out["Parent_strategy"]=p_strategyall
		}

		
		if tpl_id !=0 {
			query["Tpl_id"] = fmt.Sprintf("%d",tpl_id)
			strategyall, _:=models.GetFalconstrategyAll(query, fields, sortby, order, offset, limit)
			out["strategy"]=strategyall
		}
		
		out["tpl"]=tp
		this.jsonResult(out)
}



func (this *FalconTplController) GetApiAction() {
		out := make(map[string]interface{})
		//tpl_id,_:=this.GetInt(":id")
		act_id,_:=this.GetInt(":id")
		fmt.Println("##act_id#",act_id)

		/*
		var id int
		var err error
		if id,err = models.GetFalconactidById(tpl_id); err != nil {
				out["errInfo"] = err.Error()
				out["success"] = false
				out["errCode"] = "0006"
				this.jsonResult(out)
		}
		fmt.Println("#id",id)
		*/

		//if id!=0 {
		if act_id!=0 {
			var fields []string
			var sortby []string
			var order []string
			var falconactionquery map[string]string = make(map[string]string)
			var limit int64 = 0
			var offset int64 = 0
			//falconactionquery["Id"] = fmt.Sprintf("%d",id)
			falconactionquery["Id"] = fmt.Sprintf("%d",act_id)
			actions, _:=models.GetFalconactionAll(falconactionquery, fields, sortby, order, offset, limit)
			out["msg"] = ""
			out["data"] = actions[0]
		} else {
			out["msg"] = ""
			//out["data"] = id
			out["data"]=act_id
		}
		
		this.jsonResult(out)
}


func (this *FalconTplController) UpdateStrategy() {
	out := make(map[string]interface{})

	falconstrategy := new(models.Falconstrategy)
	falconstrategy.Id,_ = this.GetInt("sid")
	falconstrategy.Metric = strings.TrimSpace(this.GetString("metric"))
	falconstrategy.Tags = strings.TrimSpace(this.GetString("tags"))
	falconstrategy.Max_step,_ = this.GetInt("max_step")
	falconstrategy.Priority,_ = this.GetInt8("priority")
	falconstrategy.Func = strings.TrimSpace(this.GetString("func"))
	falconstrategy.Op = strings.TrimSpace(this.GetString("op"))
	falconstrategy.Right_vaue = strings.TrimSpace(this.GetString("right_value"))
	falconstrategy.Note = strings.TrimSpace(this.GetString("note"))
	falconstrategy.Run_begin = strings.TrimSpace(this.GetString("run_begin"))
	falconstrategy.Run_end = strings.TrimSpace(this.GetString("run_end"))
	//falconstrategy.Tpl_id,_ = this.GetInt("tpl_id")

	var err error
	if err = models.FalconstrategyUpdate(falconstrategy); err != nil {
			out["errInfo"] = err.Error()
			out["success"] = false
			out["errCode"] = "0006"
			this.jsonResult(out)
	} 

	out["success"] = true
	this.jsonResult(out)

}

