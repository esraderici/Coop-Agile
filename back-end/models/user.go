package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

var MissingAttribute = errors.New("Kullanici adi veya sifresi eksik.")
var UserNotFound = errors.New("Boyle Bir kullanici adi ve sifre ile eslesme olmadi.")
var UnknownError = errors.New("Bilinmeyen bir hata olustu.")
var ThereisalreadyHaveUser = errors.New("Boyle bir user daha onceden olusturulmus")

type User struct {
	Id         int `orm:"auto;serial"`
	Username   string `orm:"size(30);"`
	Firstname   string `orm:"size(30);null"`
	Lastname   string `orm:"size(30);null"`
	Password   string `orm:"size(20);"json:"-"`
	Email   string `orm:"size(40);null"`
	Type   int `orm:"size(2);null"`
}

func init() {
	orm.RegisterModel(new(User))
}

func GetUserByUsernameandPassword(Username,Password string) (User,error) {
	if (Username == "" || Password == ""){
		return User{},MissingAttribute
	}
	o := orm.NewOrm();

	User := User{}

	err := o.QueryTable("user").Filter("username",Username).Filter("password",Password).One(&User)

	if err == orm.ErrNoRows  {

		return User,UserNotFound
	}
	if err !=nil {
		return User,UnknownError
	}

	return  User,nil
}
func GetUserById(id int) User {
	var usr User
	o := orm.NewOrm();
	o.QueryTable("user").Filter("id",id).One(&usr)
	return usr

}
func GetAllUser() []User {

	var usrlist []User
	o := orm.NewOrm();
	o.QueryTable("user").All(&usrlist)

	return usrlist

}
func (this *User) Save() (error){

	if (this.Username == "" || this.Password == ""){
		return MissingAttribute
	}

	o := orm.NewOrm();
	created,_,err := o.ReadOrCreate(this,"username")

	if !created {
		return ThereisalreadyHaveUser
	}

	if err != nil {
		return UnknownError
	}
	return nil

}