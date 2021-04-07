package main

import (
	_ "github.com/Miavega/crud_test_api/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func main() {
	sqlconn, _ := beego.AppConfig.String("sqlconn")
	orm.RegisterDataBase("default", "postgres", sqlconn)
	// orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@127.0.0.1/postgres?sslmode=disable" )
	orm.Debug = true
	// orm.RegisterDataBase("default", "postgres", "postgres://"+
	// 	beego.AppConfig.String("PGuser")+":"+
	// 	beego.AppConfig.String("PGpass")+"@"+
	// 	beego.AppConfig.String("PGhost")+":"+
	// 	beego.AppConfig.String("PGport")+"/"+
	// 	beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+
	// 	beego.AppConfig.String("PGschema")+"")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

