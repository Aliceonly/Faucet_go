package main

import (
	"fmt"
	"test/controller"
	"net/http"
	"os/exec"
)

func main() {
	//处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//解析html
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//去首页
	http.HandleFunc("/main", controller.GetHome)
	//登录接口
	http.HandleFunc("/login", controller.Login)
	//邮箱登录接口
	http.HandleFunc("/Elogin", controller.ELogin)
	//注销接口
	http.HandleFunc("/logout", controller.Logout)
	//注册接口
	http.HandleFunc("/regist", controller.Regist)
	//超级管理员注册接口
	http.HandleFunc("/Adregist", controller.RegistAsAdministrator)
	/*
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	http.HandleFunc("/getUserInfo", controller.GetUserInfo)
	http.HandleFunc("/deleteUser", controller.DeleteUser)
	http.HandleFunc("/toUpdateUser", controller.ToUpdateUser)
	http.HandleFunc("/updateOraddUser", controller.UpdateOrAddUser)
	*/

	fmt.Println("后端已启动，地址为","http://127.0.0.1:8080/main")

	cmd := exec.Command("explorer", "http://127.0.0.1:8080/main")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}

	http.ListenAndServe(":8080", nil)
}
