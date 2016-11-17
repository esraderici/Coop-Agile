package controllers

import "github.com/astaxie/beego"

type PublicController struct {
	beego.Controller
}


func (controller *PublicController) Prepare() {
	//only for private area
}
