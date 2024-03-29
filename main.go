package main

import (
	_ "github.com/LiWenGu/payServer/cache"
	_ "github.com/LiWenGu/payServer/filter"
	_ "github.com/LiWenGu/payServer/models"
	_ "github.com/LiWenGu/payServer/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
