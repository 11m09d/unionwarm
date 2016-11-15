package models

import (
	"errors"
	//"fmt"
	"reflect"
	//"strconv"
	"strings"
	//"time"
	"database/sql"

	"github.com/astaxie/beego/orm"
)

type Falcongphost struct { 
	Grp_id           int      `orm:"column(grp_id)"`
	Host_id			 int 	  `orm:"column(host_id)"`
	Id               int       `orm:"column(id);auto"`
}

func (t *Falcongphost) TableName() string {
	return "grp_host"
}

func init() {
	orm.RegisterModel(new(Falcongphost))
}

type RawSeter interface {
	Exec() (sql.Result, error)
	//QueryRow(...interface{}) error
	//QueryRows(...interface{}) (int64, error)
	//SetArgs(...interface{}) RawSeter
	//Values(*[]Params, ...string) (int64, error)
	//ValuesList(*[]ParamsList, ...string) (int64, error)
	//ValuesFlat(*ParamsList, ...string) (int64, error)
	//RowsToMap(*Params, string, string) (int64, error)
	//RowsToStruct(interface{}, string, string) (int64, error)
	//Prepare() (RawPreparer, error)
}



func Falcongphostbind(m *Falcongphost)(id int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	//var r RawSeter
	//r = o.Raw("INSERT INTO `grp_host` (`grp_id`,`host_id`) VALUES (m.Grp_id,m.Host_id)")
	//r.Exec()
	id, err = o.Insert(m)
	return
}


func GetFalcongphostbindCount(id int) (cnt int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	
	cnt, err = o.QueryTable("grp_host").Filter("Grp_id", id).Count()
	return
}




/*
func FalcongphostDelbind(grp_id,id int) (err error) {
	
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falcongphost{Grp_id:grp_id,Host_id: id}
	
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Falcongphost{Grp_id:grp_id,Host_id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return


}
*/

func FalcongphostDelbind(grp_id,id int) (num int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	num, err = o.QueryTable("grp_host").Filter("Grp_id__in", grp_id).Filter("Host_id__in",id).Delete()
	return
}


func GetFalcongrphostAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	qs := o.QueryTable(new(Falcongphost))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		//fmt.Println("###k,v",k,v)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	//fmt.Println("##sortby",len(sortby))


	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}


	
	var l []Falcongphost
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}




