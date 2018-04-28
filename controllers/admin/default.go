package admin

import (
	//"github.com/astaxie/beego"
)

type MainController struct {
	CommonController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "admin/index.tpl"

}
