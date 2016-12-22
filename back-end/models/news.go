package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

var missingAtribute = errors.New("Haber eklerken eksik bilgi var")


type News struct {
	Id      int `orm:"auto;serial"`
	Title   string `orm:"size(30);null"`
	Content string `orm:"size(20);null"json:"-"`
	Userid  int `orm:"size(20);null"`
	Rate    int `orm:"size(2);null"`
	Status  int `orm:"size(2);null"`
}

func init() {
	orm.RegisterModel(new(News))
}

func GetAcceptedNewsBy(Limit,offset int) ([]*News,error) {
	var news []*News
	o := orm.NewOrm();
	qs := o.QueryTable("news")
	_,err := qs.Filter("status",1).Limit(Limit,offset).All(&news)
	return news,err
}

func GetOneNews(Id int) News {
	var news News
	o := orm.NewOrm();
	qs := o.QueryTable("news")
	qs.Filter("id",Id).One(&news)
	return news
}

func GetAllNewsBy(Limit,offset int) ([]*News,error) {
	var news []*News
	o := orm.NewOrm();
	qs := o.QueryTable("news")
	_,err := qs.Limit(Limit,offset).All(&news)
	return news,err
}

func (this *News) Delete() {
	o := orm.NewOrm();
	o.Delete(this)
}

//to save news
func (this *News) Save() error {
	if (this.Content == "") || (this.Title == "" ){
		return missingAtribute
	}

	o := orm.NewOrm();
	o.Insert(this)

	return nil
}


//update
func (this *News) Update() error {
	if (this.Content == "") || (this.Title == "" ){
		return missingAtribute
	}


	o := orm.NewOrm();
	o.Update(this)

	return nil

}


