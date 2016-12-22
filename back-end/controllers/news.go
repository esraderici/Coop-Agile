package controllers

import (
	"Coop-Agile/back-end/models"
	"strconv"
	"fmt"
)

//private page


type NewsController struct {
	PrivateController
}


//todo show list of newspage
func(this *NewsController) Newspage() {
	this.Data["firstname"] = this.GetSession("firstname")
	this.Data["lastname"] = this.GetSession("lastname")
	this.Data["email"] = this.GetSession("email")
	this.TplName = "newslist.tpl"
}
//todo show static page (form)
func(this *NewsController) NewsAdd() {

	title := this.GetString("titleTxt")
	content := this.GetString("haberIcerigitxt")
	User := this.GetSession("userid").(int)

	NewNews :=models.News{Title:title,Content:content,Userid:User,Status:0}
	NewNews.Save()

	this.Redirect("/admin", 303)
}

//todo:insert news to database.
func(this *NewsController) PublishNews() {
	newsId,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	news := models.GetOneNews(newsId)
	news.Status = 1
	news.Update()

	this.Redirect("/admin", 303)
}

func(this *NewsController) DeleteNews() {
	newsId,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))


	news := models.GetOneNews(newsId)
	news.Delete()

	fmt.Println()
	fmt.Printf("id : %v News : %v ",newsId,news)
	fmt.Println()

	this.Redirect("/admin", 303)
}