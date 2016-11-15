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

type Falconcluster struct { 
	Id             		int         `orm:"column(id);auto"`
	Grp_id				int	  		`orm:"column(grp_id);default(0)"`
	Numerator    		string      `orm:"column(numerator);size(10240);default("")"`
	Denominator    		string      `orm:"column(denominator);size(10240);default("")"`
	Endpoint    		string      `orm:"column(endpoint);size(255);default("")"`
	Metric    			string      `orm:"column(metric);size(255);default("")"`
	Tags    			string      `orm:"column(tags);size(255);default("")"`
	Ds_type    			string      `orm:"column(ds_type);size(255);default("")"`
	Step				int	        `orm:"column(step);default(0)"`
	Last_update         time.Time   `orm:"column(last_update);type(timestamp);auto_now_add"`
	Creator    			string      `orm:"column(creator);size(255);default("")"`
}

func (t *Falconcluster) TableName() string {
	return "cluster"
}

func init() {
	orm.RegisterModel(new(Falconcluster))
}



func GetFalconclusterAll(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	qs := o.QueryTable(new(Falconcluster))
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


	
	var l []Falconcluster
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


func FalconclusterAdd(m *Falconcluster) (id int64, err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	id, err = o.Insert(m)
	return
}

func DelCls(id int) (err error) {
	
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falconcluster{Id: id}
	
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Falconcluster{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func EditCls(m *Falconcluster) (err error) {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	v := Falconcluster{Id: m.Id}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m,"Numerator","Denominator","Endpoint","Metric","Tags"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func ExistFalconclusterByName(endpoint,metric,tags string) bool {
	o := orm.NewOrm()
	o.Using("falcon_portal")
	return o.QueryTable("cluster").Filter("Endpoint", endpoint).Filter("Metric", metric).Filter("Tags", tags).Exist()
}

