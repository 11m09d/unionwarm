package models

import (
	"errors"
	"fmt"
	"reflect"
	//"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Falconhost struct { 
	Id             		int       `orm:"column(id);auto"`
	Hostname        	string    `orm:"column(hostname);size(255);default("")"`
	Ip              	string    `orm:"column(ip);size(16);default("")"`
	Agent_version       string    `orm:"column(agent_version);size(16);default("")"`
	Plugin_version      string    `orm:"column(plugin_version);size(128);default("")"`	
	Maintain_begin      int       `orm:"column(maintain_begin);default(0)"`
	Maintain_end        int       `orm:"column(maintain_end);default(0)"`
	Update_at           time.Time `orm:"column(update_at);type(timestamp);auto_now_add"`
}

func (t *Falconhost) TableName() string {
	return "host"
}

func init() {
	orm.RegisterModel(new(Falconhost))
}


func FalconhostAdd(m *Falconhost) (id int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	id, err = o.Insert(m)
	return
}


func ExistFalconhostByName(name string) bool {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	return o.QueryTable("host").Filter("Hostname", name).Exist()
}



func GetFalconhostidbyHn(hn string) (hid int, err error) {
	var h Falconhost
	o := orm.NewOrm()
	o.Using("falcon_portal")
	if err = o.QueryTable("host").Filter("Hostname", hn).One(&h); err != nil {
		return
	}
	return h.Id,nil
}


func GetFalconhostnamebyid(hid int) (hn string, err error) {
	var h Falconhost
	o := orm.NewOrm()
	o.Using("falcon_portal")
	if err = o.QueryTable("host").Filter("Id", hid).One(&h); err != nil {
		return
	}
	return h.Hostname,nil
}


func GetFalconhostAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	qs := o.QueryTable(new(Falconhost))
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


	
	var l []Falconhost
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



func UpdateMaintainTime(m *Falconhost) (err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falconhost{Id: m.Id}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m,"Maintain_begin","Maintain_end"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}