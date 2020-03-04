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
	"log"
	"os"
)

func main() {
	var (
		mode = 0
		err  error
	)
	fmt.Println("服务器更新工具版本:0.0.2")
	configUpload := ""
	config := ""
	//读取ssh相关的配置文件
	if len(os.Args) > 1 {
		config = os.Args[1]
		if len(os.Args) > 2 {
			configUpload = os.Args[2]
		}
	}

	if config == "" {
		config = "config/config.toml"
	}
	if configUpload == "" {
		configUpload = "config/upload.toml"
	}
	sshConfig, choose, err := ReadSSHConfig(config)
	fmt.Println("1.全覆盖---->意味着无差别覆盖本地路径的数据覆盖到远程服务器上")
	fmt.Println("请输入匹配模式(1.全覆盖，2:只覆盖config文件):")
	//for {
	_, err = fmt.Scan(&mode)
	if err != nil {
		fmt.Println("输入错误:->", err)
	}

	flag := 0
	//用来表示所选择的匹配模式
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

	_, err = toml.DecodeFile(configUpload, &strings)
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
	//用于处理进行填充和过滤所需要的远程服务器中
	for k, v := range strings {
		if k != "Local" {
			tag := true
			m := v.(map[string]interface{})
			val, ok := m["svn"]
			if ok {
				if val != choose {
					tag = false
				}
			}
			if tag {
				remotes = append(remotes, v.(map[string]interface{}))
			}

		}
	}
	//发送文件到远程服务器中
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

/**
读取ssh配置,需要输入配置的路径
*/
func ReadSSHConfig(path string) (model.SshConfig, string, error) {
	sshAllConfig := make(map[string]interface{})
	var configObj model.SshConfig
	choose := ""
	err := utils.ReadConfigByMap(path, &sshAllConfig)
	if err != nil {
		fmt.Println("读取配置文件时发生错误:", err)
		log.Fatal()
	}
	//如果  配置数量大于1 需要 用户自己选择配置
	if len(sshAllConfig) > 1 {
		fmt.Println("请选择一个SSH:")
		for k, v := range sshAllConfig {
			fmt.Printf("输入%s,该服务器指向-->%s\n", k, v.(map[string]interface{})["Annota"])
		}
		_, err := fmt.Scan(&choose)
		v, ok := sshAllConfig[choose]
		if ok {
			configObj = fillSshConfigObj(v.(map[string]interface{}))
		} else {
			log.Fatal("输入的ssh配置未找到!")
		}
		if err != nil {
			fmt.Println("输入的时候发生错误:", err)
			log.Fatal()
		}
	} else {
		//配置数量小于2的时候则默认选择该配置
		for _, v := range sshAllConfig {
			configObj = fillSshConfigObj(v.(map[string]interface{}))
		}
	}
	return configObj, choose, nil
}

func fillSshConfigObj(obj map[string]interface{}) model.SshConfig {
	sshConfig := model.SshConfig{}
	sshConfig.HostName = obj["HostName"].(string)
	sshConfig.Port = int16(obj["Port"].(int64))
	sshConfig.UserName = obj["UserName"].(string)
	sshConfig.Password = obj["Password"].(string)
	return sshConfig
}
