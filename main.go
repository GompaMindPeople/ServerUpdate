/*
@Time : 2019/7/22 13:24
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package main

import (
	"ServerUpdate/model"
	"ServerUpdate/service"
	"ServerUpdate/utils"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

func main() {
	var (
		mode = 0
		err  error
	)
	fmt.Println("服务器更新工具版本:0.0.1")
	fmt.Println("1.全覆盖---->意味着无差别覆盖本地路径的数据覆盖到远程服务器上")
	fmt.Println("请输入匹配模式(1.全覆盖，2:只覆盖config文件):")
	//for {
	_, err = fmt.Scan(&mode)
	if err != nil {
		fmt.Println("输入错误:->", err)
	}
	flag := 0
	switch mode {
	case 1:
		break
	case 2:
		flag = 2
		break
	default:
		fmt.Println("无该选项")
		os.Exit(1)
	}
	//}

	sshConfig := model.SshConfig{}
	err = utils.ReadConfigByObj("config/config.toml", &sshConfig)
	if err != nil {
		fmt.Println("读取配置文件时发生错误:", err)
		return
	}
	//获取 ssh 的实例
	bean, err := utils.GetSshBean(sshConfig.UserName, sshConfig.Password, sshConfig.HostName, sshConfig.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer bean.Client.Close()
	//strings := model.TomlConfig{}
	//获取 sftp 实例
	sftpClient, err := bean.GetSftpConnect()
	defer sftpClient.Close()

	strings := make(map[string]interface{})

	_, err = toml.DecodeFile("config/upload.toml", &strings)
	if err != nil {
		fmt.Println("读取上传配置文件文件时发生错误:", err)
		return
	}
	local := strings["Local"].(map[string]interface{})
	//2表示只覆盖  config文件
	if flag == 2 {
		for k := range local {
			if k != "Data" {
				delete(local, k)
			}
		}
	}
	remotes := make([]map[string]interface{}, 0)
	for k, v := range strings {
		if k != "Local" {
			remotes = append(remotes, v.(map[string]interface{}))
		}
	}

	for key, value := range local {
		for _, remote := range remotes {
			//请求远程服务器并准备上传文件
			service.UploadLocalByRemote(sftpClient, value.(string), remote[key].(string))
		}
	}

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("输入任意键退出!")
	_, _ = fmt.Scan(&mode)
}
