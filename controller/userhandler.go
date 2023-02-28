package controller

import (
	"cmdb/dao"
	"net/http"
	"text/template"
)

// 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//查询是否有这个用户
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.Uid > 0 {
		t := template.Must(template.ParseFiles(("views/pages/user/login_success.html")))
		t.Execute(w, user)
	} else {
		//用户名或密码不正确
		t := template.Must(template.ParseFiles(("views/pages/user/login.html")))
		t.Execute(w, "用户名或密码不正确")
	}
}

// 处理用户注册的函数
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//用户名可用，将用户保存在数据库中（还缺少用户名重复的判断）
	user, _ := dao.CheckUserName(username)
	if user.Uid > 0 {
		//用户已经存储
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户已存在")
	} else {
		//用户名可用，将用户保存在数据库中
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles(("views/pages/user/regist_success.html")))
		t.Execute(w, "")
	}
}
