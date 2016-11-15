package g

import (
	"github.com/astaxie/beego/cache"
	//"fmt"
)

var (
	Cache cache.Cache
)

func init() {
	Cache, _= cache.NewCache("memory", `{"interval":60}`)
	
}