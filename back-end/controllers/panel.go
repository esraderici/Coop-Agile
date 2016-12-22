package controllers

import (
	"Coop-Agile/back-end/models"
	"fmt"
)

type PanelController struct {
	PrivateController
}

type ResponseStruct struct {
	News models.News
	User models.User
}

func (this *PanelController) AdminPanelStatic()  {
	this.Data["firstname"] = this.GetSession("firstname")
	this.Data["lastname"] = this.GetSession("lastname")
	this.Data["email"] = this.GetSession("email")

	NewsList,err := models.GetAllNewsBy(0,0)

	var response []ResponseStruct
	if err != nil {

	}else{
		for _, News := range NewsList{
			usr := models.GetUserById(News.Userid)
			response = append(response,ResponseStruct{News:*News,User:usr})
		}
	}
	this.Data["liste"] = response
	fmt.Println(response)
	this.TplName = "admin.tpl"
}

func (this *PanelController) Logout() {
	this.DelSession("login")
	this.DelSession("userid")
	this.DelSession("username")
	this.DelSession("type")
	this.DelSession("firstname")
	this.DelSession("lastname")
	this.DelSession("email")
	this.Redirect("/login",302)
}