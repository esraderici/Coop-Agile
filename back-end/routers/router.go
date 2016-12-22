package routers

import (
	"Coop-Agile/back-end/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin", &controllers.PanelController{},"Get:AdminPanelStatic")
	beego.Router("/admin/logout", &controllers.PanelController{},"Get:Logout")

	//index pages
	beego.Router("/", &controllers.IndexController{},"Get:Index")
	beego.Router("/index", &controllers.IndexController{},"Get:Index")

	//news pages
	beego.Router("/admin/news", &controllers.NewsController{},"Get:Newspage")
	beego.Router("/admin/news", &controllers.NewsController{},"Post:NewsAdd")
	beego.Router("/admin/publish/?:id", &controllers.NewsController{},"Get:PublishNews")
	beego.Router("/admin/delete/?:id", &controllers.NewsController{},"Get:DeleteNews")
	//Users Page
	beego.Router("/admin/user", &controllers.PrivateUserController{},"Get:Userlist")

	//Login page Static file
	beego.Router("/login", &controllers.UserController{}, "Get:LoginStaticPage")
	beego.Router("/login", &controllers.UserController{}, "Post:LoginControl")


	beego.Router("/user/register", &controllers.UserController{}, "Get:RegisterStaticPage")
	beego.Router("/user/register", &controllers.UserController{}, "Post:RegisterControl")
}
