package main

import (
	_ "github.com/chenqi123/unionwarm/routers"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/chenqi123/unionwarm/models"
	"fmt"
)

func id_host_cnt(in int)(out string){
	cnt,_:=models.GetFalcongphostbindCount(in)
	out=fmt.Sprintf("%d",cnt)	
    return out
}

func id_strategy_cnt(in int)(out string){
	cnt,_:=models.GetFalconstrategyCountbytplid(in)
	out=fmt.Sprintf("%d",cnt)	
    return out
}

func main() {
	beego.SetStaticPath("/zjz_work_cmdb/static","static")
	orm.RunSyncdb("default", false, true)
	beego.AddFuncMap("idhostcnt",id_host_cnt)
	beego.AddFuncMap("idstrategycnt",id_strategy_cnt)

	beego.Run()
}

