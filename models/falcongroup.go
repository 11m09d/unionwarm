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

type Falcongroup struct { 
	Id             		int       `orm:"column(id);auto"`
	Grp_name    		string    `orm:"column(grp_name);size(255);default("")"`
	Create_user         string    `orm:"column(create_user);size(64);default("")"`
	Create_at           time.Time `orm:"column(create_at);type(timestamp);auto_now_add"`
	Come_from           int8      `orm:"column(come_from);default(0)"`
}

func (t *Falcongroup) TableName() string {
	return "grp"
}

func init() {
	orm.RegisterModel(new(Falcongroup))
}



func FalcongroupAdd(m *Falcongroup) (id int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	id, err = o.Insert(m)
	return
}

func ExistFalcongroupByName(name string) bool {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	return o.QueryTable("grp").Filter("Grp_name", name).Exist()
}


func DelGroup(id int) (err error) {
	
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falcongroup{Id: id}
	
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Falcongroup{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func GetFalcongrpnamebygid(gid int) (gn string, err error) {
	var h Falcongroup
	o := orm.NewOrm()
	o.Using("falcon_portal")
	if err = o.QueryTable("grp").Filter("Id", gid).One(&h); err != nil {
		return
	}
	return h.Grp_name,nil
}



func GetFalcongroupAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	qs := o.QueryTable(new(Falcongroup))
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


	
	var l []Falcongroup
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
