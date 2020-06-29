package main

import (
	"github.com/astaxie/beego"
	_ "k8-web-terminal/controllers"
	_ "k8-web-terminal/routers"
	_ "k8-web-terminal/util/logger"
)

func main() {
	beego.SetStaticPath("/assets", "./static/assets")
	beego.SetStaticPath("/public", "./static")
	beego.Run()
}
