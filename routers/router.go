package routers

import (
	"unionwarm/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.FalconTplController{},"*:Index")
    beego.Router("/zjz_work_cmdb/cmdb/logout", &controllers.LoginController{}, "get:Logout")
	
    //cq 0926
	beego.Router("/zjz_work_cmdb/cmdb/falcontpl", &controllers.FalconTplController{}, "*:Index")

	//cq 0927
	beego.Router("/zjz_work_cmdb/cmdb/falconnewtpl", &controllers.FalconTplController{}, "*:NewTpl")
	beego.Router("/zjz_work_cmdb/cmdb/falconaddtpl", &controllers.FalconTplController{}, "*:AddTpl")
	beego.Router("/zjz_work_cmdb/cmdb/falcondeltpl", &controllers.FalconTplController{}, "*:DelTpl")
	beego.Router("/zjz_work_cmdb/cmdb/falconnewstrategy", &controllers.FalconTplController{}, "*:NewStrategy")

	//cq 0928
	beego.Router("/zjz_work_cmdb/cmdb/falconshowstrategy", &controllers.FalconTplController{}, "*:ShowStrategy")
	beego.Router("/zjz_work_cmdb/cmdb/falconaddstrategy", &controllers.FalconTplController{}, "*:AddStrategy")

	//cq 0929
	//beego.Router("/zjz_work_cmdb/cmdb/falconaddhost", &controllers.FalconGphostController{}, "*:AddHost")
	beego.Router("/zjz_work_cmdb/cmdb/falconnewgrp", &controllers.FalconGphostController{}, "*:NewGroup")
	beego.Router("/zjz_work_cmdb/cmdb/falconaddgrp", &controllers.FalconGphostController{}, "*:AddGroup")
	beego.Router("/zjz_work_cmdb/cmdb/falcondelgrp", &controllers.FalconGphostController{}, "*:DelGroup")
	beego.Router("/zjz_work_cmdb/cmdb/falconbindgphost", &controllers.FalconGphostController{}, "*:BindGpHost")
	beego.Router("/zjz_work_cmdb/cmdb/falconhg", &controllers.FalconGphostController{}, "*:Index")
	beego.Router("/zjz_work_cmdb/cmdb/falcongphost", &controllers.FalconGphostController{}, "*:ShowGrpHost")


	//cq1001
	beego.Router("/zjz_work_cmdb/cmdb/falconaddhosts", &controllers.FalconGphostController{}, "*:AddHosts")

	//cq1004
	beego.Router("/zjz_work_cmdb/cmdb/falconhgdel", &controllers.FalconGphostController{}, "*:DelbindHosts")

	//cq1005
	beego.Router("/zjz_work_cmdb/cmdb/falconshowgrptpl", &controllers.FalconTplController{}, "*:ShowGrpTpl")

	//cq1008
	beego.Router("/zjz_work_cmdb/cmdb/falcongptpl", &controllers.FalconTplController{}, "*:ShowTplGrp")
	beego.Router("/zjz_work_cmdb/cmdb/falconquerytpl", &controllers.FalconTplController{}, "*:QueryTpl")
	beego.Router("/zjz_work_cmdb/cmdb/falconbindgrptpl", &controllers.FalconTplController{}, "*:BindGrpTpl")
	beego.Router("/zjz_work_cmdb/cmdb/falcongrptpldel", &controllers.FalconTplController{}, "*:DelGrpTpl")


	//cq1009
	beego.Router("/zjz_work_cmdb/cmdb/falconparenttpl", &controllers.FalconTplController{}, "*:UpdateParenttpl")
	beego.Router("/zjz_work_cmdb/cmdb/falcongetactioncount", &controllers.FalconTplController{}, "*:GetActionCount")
	beego.Router("/zjz_work_cmdb/cmdb/falcongetuiccount", &controllers.FalconTplController{}, "*:GetUicCount")

	//cq1010
	beego.Router("/zjz_work_cmdb/cmdb/falcongetactiondetail", &controllers.FalconTplController{}, "*:GetActionDetail")
	beego.Router("/zjz_work_cmdb/cmdb/falconactionmodify", &controllers.FalconTplController{}, "*:ModifyAction")

	//cq1013
	beego.Router("/zjz_work_cmdb/cmdb/api/action/:id([0-9]+)", &controllers.FalconTplController{}, "*:GetApiAction") 
	//cq1014 temp
	//beego.Router("/zjz_work_cmdb/cmdb/falconagg", &controllers.FalconTplController{}, "*:ClusterIndex") 

	//cq1017 
	beego.Router("/zjz_work_cmdb/cmdb/falcondelstrategy", &controllers.FalconTplController{}, "*:DelStrategy")
	beego.Router("/zjz_work_cmdb/cmdb/falconhostmaintain", &controllers.FalconGphostController{}, "*:UpdateMaintainTime")

	//cq1018
	beego.Router("/zjz_work_cmdb/cmdb/falconucintf", &controllers.FalconTplController{}, "*:UcintfIndex")


	//cq1025
	beego.Router("/zjz_work_cmdb/cmdb/falconupdatestrategy", &controllers.FalconTplController{}, "*:UpdateStrategy")


	//cq1031
	beego.Router("/zjz_work_cmdb/cmdb/template/view/:id([0-9]+)", &controllers.FalconTplController{}, "*:GetTplDetail")

	//cq1103
	beego.Router("/zjz_work_cmdb/cmdb/falconagg", &controllers.FalconClusterController{}, "*:Index") 
	beego.Router("/zjz_work_cmdb/cmdb/falconnewcls", &controllers.FalconClusterController{}, "*:NewCls") 
	beego.Router("/zjz_work_cmdb/cmdb/falconaddcls", &controllers.FalconClusterController{}, "*:AddCls") 
	beego.Router("/zjz_work_cmdb/cmdb/falcondelcls", &controllers.FalconClusterController{}, "*:DelCls") 
	beego.Router("/zjz_work_cmdb/cmdbapp/falconeditcls", &controllers.FalconClusterController{}, "*:EditCls") 

	//cq1109
	beego.Router("/zjz_work_cmdb/cmdb/api/eptohid/:id([0-9]+)", &controllers.FalconGphostController{}, "*:GetEndpointToHostid") 
	beego.Router("/zjz_work_cmdb/cmdb/api/getgrpname/:id([0-9]+)", &controllers.FalconGphostController{}, "*:GetApiGrpName") 
	beego.Router("/zjz_work_cmdb/cmdb/api/addapihosts", &controllers.FalconGphostController{}, "*:AddApiHosts") 

	//cq1115
	beego.Router("/zjz_work_cmdb/cmdb/falconucintf_addhosts", &controllers.FalconTplController{}, "*:UcintfAdd")


}
