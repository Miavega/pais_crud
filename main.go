package main

import (
	_ "github.com/miavega/pais_crud/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/customerror"
	_ "github.com/lib/pq"
)

func main() {
  orm.Debug = true
  orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
  if beego.BConfig.RunMode == "dev" {
  	beego.BConfig.WebConfig.DirectoryIndex = true
  	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
  }
  beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
    AllowHeaders: []string{"Origin", "x-requested-with",
      "content-type",
      "accept",
      "origin",
      "authorization",
      "x-csrftoken"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
  }))
  logs.SetLogger(logs.AdapterFile, `{"filename":"/var/log/beego/pais_crud/pais_crud.log"}`)
  beego.ErrorController(&customerror.CustomErrorController{})

  beego.Run()
}

