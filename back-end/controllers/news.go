package controllers

//private page


type NewsController struct {
	PrivateController
}


//todo show list of newspage
func(this *PrivateController) Newspage() {
	this.TplName = "newslist.tpl"
}
//todo show static page (form)
func(this *PrivateController) NewsAddStatic() {
	this.TplName = "newsaddpage.tpl"
}

//todo:insert news to database.
func(this *PrivateController) NewsAdd() {
	this.TplName = "admin.tpl"
}