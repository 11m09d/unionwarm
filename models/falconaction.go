package models

import (
	"errors"
	"fmt"
	"reflect"
	//"strconv"
	"strings"
	//"time"

	"github.com/astaxie/beego/orm"
)

type Falconaction struct { 
	Id             			int       `orm:"column(id);auto"`
	Uic    					string    `orm:"column(uic);size(255);default("")"`
	Url         			string    `orm:"column(url);size(255);default("")"`
	Callback				int8	  `orm:"column(callback);default(0)"`
	Before_callback_sms		int8	  `orm:"column(before_callback_sms);default(0)"`
	Before_callback_mail	int8	  `orm:"column(before_callback_mail);default(0)"`
	After_callback_sms		int8	  `orm:"column(after_callback_sms);default(0)"`
	After_callback_mail		int8	  `orm:"column(after_callback_mail);default(0)"`
}

func (t *Falconaction) TableName() string {
	return "action"
}

func init() {
	orm.RegisterModel(new(Falconaction))
}


func GetFalconActionCount() (cnt int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	cnt, err = o.QueryTable("action").Count()
	return
}


func GetFalconactionAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	qs := o.QueryTable(new(Falconaction))
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


	
	var l []Falconaction
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


func FalconactionAdd(m *Falconaction) (id int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	id, err = o.Insert(m)
	return
}


func UpdateAction(m *Falconaction) (err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falconaction{Id: m.Id}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m,"Uic","Url","Callback","Before_callback_sms","Before_callback_mail","After_callback_sms","After_callback_mail"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}


func GetFalconActionById(id int) (v *Falconaction, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v = &Falconaction{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}