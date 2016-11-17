package controllers


type PanelController struct {
	PrivateController
}

func (this *PanelController) AdminPanelStatic()  {
	this.Data["firstname"] = this.GetSession("firstname")
	this.Data["lastname"] = this.GetSession("lastname")
	this.Data["email"] = this.GetSession("email")
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