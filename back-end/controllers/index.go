package controllers

import "Coop-Agile/back-end/models"

type IndexController struct {
	PublicController
}



func(this *IndexController) Index() {
	newslist ,err := models.GetAcceptedNewsBy(0,0)
	var response []ResponseStruct
	if err != nil {

	}else{
		for _, News := range newslist{
			usr := models.GetUserById(News.Userid)
			response = append(response,ResponseStruct{News:*News,User:usr})
		}
	}
	this.Data["liste"] = response
	this.TplName = "index.tpl"

}
