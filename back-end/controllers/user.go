package controllers

import (
	"Coop-Agile/back-end/models"
)


//login page
//login control
//register page
//register control

type UserController struct {
	PublicController
}
//login page static
func (this *UserController) LoginStaticPage() {
	if this.GetSession("error") != "" {
		this.Data["error"] = this.GetSession("error")
		this.SetSession("error","")
	}

	this.TplName = "login.tpl"

	return
}

func (this *UserController) LoginControl() {
	//get variable
	//find user
	username := this.GetString("username")
	password := this.GetString("password")

	user,err := models.GetUserByUsernameandPassword(username,password)
	this.Data["error"] = ""
	if err != nil {
		this.SetSession("error",err.Error())
		this.Redirect("/login",302)
		return
	}else{
		this.SetSession("login","k3y")
		this.SetSession("userid",user.Id)
		this.SetSession("username",user.Username)
		this.SetSession("type",user.Type)
		this.SetSession("firstname",user.Firstname)
		this.SetSession("lastname",user.Lastname)
		this.SetSession("email",user.Email)
	}
	this.Redirect("/admin",302)
	return

}

//register page static
func (this *UserController) RegisterStaticPage() {
	if this.GetSession("error") != "" {
		this.Data["error"] = this.GetSession("error")
		this.SetSession("error","")
	}
	this.TplName = "register.tpl"
	return

}

func (this *UserController) RegisterControl() {
	//Register verileri ile ilgili kontrol yapilacak

	username := this.GetString("username")
	password := this.GetString("password")
	lastname := this.GetString("lastname")
	email := this.GetString("email")
	firstname := this.GetString("firstname")

	user := models.User{Username:username,Password:password,Lastname:lastname,Email:email,Firstname:firstname}
	err := user.Save()

	if err != nil {
		this.SetSession("error",err.Error())
		this.Redirect("/user/register",302)
	}
		this.Redirect("/login",302)



}
