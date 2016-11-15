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

type Falconstrategy struct { 
	Id              int       `orm:"column(id);auto"`
	Metric          string    `orm:"column(metric);size(128);default("")"`
	Tags            string    `orm:"column(tags);size(256);default("")"`
	Max_step        int       `orm:"column(max_step);default(1)"`
	Priority        int8      `orm:"column(priority);default(0)"`
	Func            string    `orm:"column(func);size(16);default("all(#1)")"`
	Op              string    `orm:"column(op);size(8);default("")"`
	Right_vaue      string    `orm:"column(right_value);size(64)"`
	Note            string    `orm:"column(note);size(128);default("")"`
	Run_begin       string    `orm:"column(run_begin);size(16);default("")"`
	Run_end         string    `orm:"column(run_end);size(16);default("")"`
	Tpl_id          int       `orm:"column(tpl_id);default(0)"`
}

func (t *Falconstrategy) TableName() string {
	return "strategy"
}

func init() {
	orm.RegisterModel(new(Falconstrategy))
}


func GetFalconstrategyCount() (cnt int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	cnt, err = o.QueryTable("strategy").Count()
	return
}

func GetFalconstrategyCountbytplid(id int) (cnt int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	
	cnt, err = o.QueryTable("strategy").Filter("Tpl_id", id).Count()
	return
}


func GetFalconstrategyAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	qs := o.QueryTable(new(Falconstrategy))
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


	
	var l []Falconstrategy
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


func FalconstrategyAdd(m *Falconstrategy) (id int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	id, err = o.Insert(m)
	return
}

func DelStrategy(id int) (err error) {
	
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falconstrategy{Id: id}
	
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Falconstrategy{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func FalconstrategyUpdate(m *Falconstrategy) (err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falconstrategy{Id: m.Id}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m,"Metric","Tags","Max_step","Priority","Func","Op","Right_vaue","Note","Run_begin","Run_end"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
