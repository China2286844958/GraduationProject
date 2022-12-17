package tools

import (
	"fmt"
	"log"
	"os/exec"
)

/**
@Title 服务器启动完成后，自动跳转指定页面
@Author 薛智敏
@CreateTime 2022年8月2日22:32:10
*/

//
//  AutoHref
//  @Description:
//  @param href 地址
//

func AutoHref(href string) {
	//cmd := exec.Command("explorer", "http://localhost:8848/User/login")
	cmd := exec.Command("explorer", href)
	err := cmd.Start()
	if err != nil {
		fmt.Println("自动跳转捕捉的错误:" + err.Error())
	}
}

//
//  RestartMySqlServer 重启mysql服务器
//  @Description:
//

func RestartMySqlServer() error {
	cmd := exec.Command("net", "start", "mysql80")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to call cmd.Run(): %v", err)
		return err
	}
	return nil
}
