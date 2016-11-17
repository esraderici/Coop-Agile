package routers

import (
	"Coop-Agile/back-end/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin", &controllers.PanelController{},"Get:AdminPanelStatic")
	beego.Router("/admin/logout", &controllers.PanelController{},"Get:Logout")

	//Login page Static file
	beego.Router("/login", &controllers.UserController{}, "Get:LoginStaticPage")
	beego.Router("/login", &controllers.UserController{}, "Post:LoginControl")


	beego.Router("/user/register", &controllers.UserController{}, "Get:RegisterStaticPage")
	beego.Router("/user/register", &controllers.UserController{}, "Post:RegisterControl")
}
