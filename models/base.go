package models

import (
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")

	if dbport == "" {
		dbport = "3306"
	}

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)


	//falcon by cq
	dbhost1 := beego.AppConfig.String("db1.host")
	dbport1 := beego.AppConfig.String("db1.port")
	dbuser1 := beego.AppConfig.String("db1.user")
	dbpassword1 := beego.AppConfig.String("db1.password")
	dbname1 := beego.AppConfig.String("db1.name")
	timezone1 := beego.AppConfig.String("db1.timezone")

	if dbport1 == "" {
		dbport1 = "3306"
	}

	dsn1 := dbuser1 + ":" + dbpassword1 + "@tcp(" + dbhost1 + ":" + dbport1 + ")/" + dbname1 + "?charset=utf8"
	if timezone1 != "" {
		dsn1 = dsn1 + "&loc=" + url.QueryEscape(timezone1)
	}
	orm.RegisterDataBase("falcon", "mysql", dsn1)



	//falcon new template by cq
	dbhost2 := beego.AppConfig.String("db02.host")
	dbport2 := beego.AppConfig.String("db02.port")
	dbuser2 := beego.AppConfig.String("db02.user")
	dbpassword2 := beego.AppConfig.String("db02.password")
	dbname2 := beego.AppConfig.String("db02.name")
	timezone2 := beego.AppConfig.String("db02.timezone")

	if dbport2 == "" {
		dbport2 = "3306"
	}

	dsn2 := dbuser2 + ":" + dbpassword2 + "@tcp(" + dbhost2 + ":" + dbport2 + ")/" + dbname2 + "?charset=utf8"
	if timezone2 != "" {
		dsn2 = dsn2 + "&loc=" + url.QueryEscape(timezone2)
	}
	orm.RegisterDataBase("falcon_portal", "mysql", dsn2)

	//falcon uic new by cq
	dbhost3 := beego.AppConfig.String("db03.host")
	dbport3 := beego.AppConfig.String("db03.port")
	dbuser3 := beego.AppConfig.String("db03.user")
	dbpassword3 := beego.AppConfig.String("db03.password")
	dbname3 := beego.AppConfig.String("db03.name")
	timezone3 := beego.AppConfig.String("db03.timezone")

	if dbport3 == "" {
		dbport3 = "3306"
	}

	dsn3 := dbuser3 + ":" + dbpassword3 + "@tcp(" + dbhost3 + ":" + dbport3 + ")/" + dbname3 + "?charset=utf8"
	if timezone3 != "" {
		dsn3 = dsn3 + "&loc=" + url.QueryEscape(timezone3)
	}
	orm.RegisterDataBase("falcon_uic", "mysql", dsn3)




		if beego.AppConfig.String("runmode") == "dev" {
			orm.Debug = true
		}
}
