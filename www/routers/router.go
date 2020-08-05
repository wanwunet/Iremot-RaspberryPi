package routers

import (
	"Iremot-RaspberryPi/www/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
