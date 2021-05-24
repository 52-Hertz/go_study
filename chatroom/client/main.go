package main

import "fmt"

// 用户id, 用户密码
var userId int
var userPwd string

func main() {
	// 登录界面
	var key int
	var loop bool = true
	for loop {
		fmt.Println("------------欢迎登录多人聊天系统------------")
		fmt.Println("\t\t\t 1. 登录聊天")
		fmt.Println("\t\t\t 2. 注册用户")
		fmt.Println("\t\t\t 3. 退出系统")
		fmt.Print("\t\t\t 请选择[1-3]：")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("\t\t\t 登录聊天")
			loop = false
		case 2:
			fmt.Println("\t\t\t 注册用户")
			loop = false
		case 3:
			fmt.Println("\t\t\t 退出系统")
			loop = false
		case 4:
			fmt.Println("\t\t\t 你登录有误，请重新登录...")
		}
	}

	//登录界面
	if 1 == key {
		fmt.Print("请输入用户Id: ")
		fmt.Scanf("%d\n", &userId)
		fmt.Print("请输入用户密码: ")
		fmt.Scanf("%s\n", &userPwd)

		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录错误")
		}
	} else if 2 == key{
		fmt.Println("用户注册逻辑.......")
	}
}