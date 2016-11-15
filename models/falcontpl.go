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

type Falcontpl struct { 
	Id              int       `orm:"column(id);auto"`
	Tpl_name        string    `orm:"column(tpl_name);size(255);default("")"`
	Parent_id       int       `orm:"column(parent_id);default(0)"`
	Action_id       int       `orm:"column(action_id);default(0)"`
	Create_user     string    `orm:"column(create_user);size(64);default("")"`
	Create_at       time.Time `orm:"column(create_at);type(timestamp);auto_now_add"`
}

func (t *Falcontpl) TableName() string {
	return "tpl"
}

func init() {
	orm.RegisterModel(new(Falcontpl))
}


func GetFalcontplCount() (cnt int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	cnt, err = o.QueryTable("tpl").Count()
	return
}


func GetFalcontplAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	qs := o.QueryTable(new(Falcontpl))
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


	
	var l []Falcontpl
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



func FalcontplAdd(m *Falcontpl) (id int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	id, err = o.Insert(m)
	return
}



func GetFalcontplById(id int) (v *Falcontpl, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v = &Falcontpl{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


func ExistFalcontplByName(name string) bool {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	return o.QueryTable("tpl").Filter("Tpl_name", name).Exist()
}


func DelTpl(id int) (err error) {
	
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falcontpl{Id: id}
	
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Falcontpl{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func UpdateParentid(m *Falcontpl) (err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falcontpl{Id: m.Id}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m,"Parent_id"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}


func GetFalconactidById(id int) (action_id int, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := &Falcontpl{Id: id}
	if err = o.Read(v); err == nil {
		return v.Action_id, nil
	}
	return 0, err
}


func UpdateActionid(m *Falcontpl) (err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falcontpl{Id: m.Id}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m,"Action_id"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

