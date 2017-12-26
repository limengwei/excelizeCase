package routers

import (
	"hellogui/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainCtrl{})

}
