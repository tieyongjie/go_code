package main

import "fmt"

type Pay struct {
	Account string
	Pwd     string
	Money   float32
}

func (pay *Pay) SaveMoney(num float32, pwd string) {
	if pwd != pay.Pwd {
		fmt.Println("密码输入错误")
	} else {
		pay.Money += num
		fmt.Println("存款成功~,您的余额为:", pay.Money)
	}
}

func (pay *Pay) TakeMoney(num float32, pwd string) {
	if pwd != pay.Pwd {
		fmt.Println("密码输入错误")
	} else {
		if pay.Money >= num {
			pay.Money -= num
			fmt.Println("取款成功~,您的余额为:", pay.Money)
		} else {
			fmt.Println("余额不足,取款失败...")
		}
	}
}

func (pay *Pay) LookMoney(pwd string) {
	if pwd != pay.Pwd {
		fmt.Println("密码输入错误")
	} else {
		fmt.Println("您的余额为:", pay.Money)
	}
}

func main() {
	pay := Pay{"yongjie", "123456", 100}
	var action int
	var count float32
	var pwd string
Loop:
	for {
		fmt.Print("请输入要执行的操作：\n1：存款\n2：取款\n3：查询\n4：退出\n")
		//对变量取地址
		fmt.Scanln(&action)

		switch action {
		case 1:
			fmt.Print("请输入要存的金额：")
			fmt.Scanln(&count)
			fmt.Print("请输入密码：")
			fmt.Scanln(&pwd)
			pay.SaveMoney(count, pwd)
		case 2:
			fmt.Print("请输入要取的金额：")
			fmt.Scanln(&count)
			fmt.Print("请输入密码：")
			fmt.Scanln(&pwd)
			pay.TakeMoney(count, pwd)
		case 3:
			fmt.Print("请输入密码：")
			fmt.Scanln(&pwd)
			pay.LookMoney(pwd)
		case 4:
			break Loop
		}
	}
}
