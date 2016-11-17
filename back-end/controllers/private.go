package controllers

import (
	"github.com/astaxie/beego"
)

type PrivateController struct {
	beego.Controller
}


func (controller *PrivateController) Prepare() {
	v := controller.GetSession("login")
	if v == nil {
		controller.Redirect("/login", 304)
	}

//only for private area
}