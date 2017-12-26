package routers

import (
	"helloexcel/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainCtrl{})

}
