package routers

import (
	"sadmin_beego/controllers"
	"github.com/astaxie/beego"
	"sadmin_beego/controllers/admin"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &admin.MainController{})
}
