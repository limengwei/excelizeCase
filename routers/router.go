package routers

import (
	"helloexcel/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainCtrl{})
	beego.Router("/gis", &controllers.MainCtrl{}, "*:Gis")

}
