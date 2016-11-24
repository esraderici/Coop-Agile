package routers

import (
	"Coop-Agile/back-end/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin", &controllers.PanelController{},"Get:AdminPanelStatic")
	beego.Router("/admin/logout", &controllers.PanelController{},"Get:Logout")

	//news pages
	beego.Router("/admin/news", &controllers.NewsController{},"Get:Newspage")
	beego.Router("/admin/news/add", &controllers.NewsController{},"Get:NewsAddStatic")
	beego.Router("/admin/news", &controllers.NewsController{},"Post:NewsAdd")

	//Login page Static file
	beego.Router("/login", &controllers.UserController{}, "Get:LoginStaticPage")
	beego.Router("/login", &controllers.UserController{}, "Post:LoginControl")


	beego.Router("/user/register", &controllers.UserController{}, "Get:RegisterStaticPage")
	beego.Router("/user/register", &controllers.UserController{}, "Post:RegisterControl")
}
