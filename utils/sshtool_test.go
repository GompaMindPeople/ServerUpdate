/*
@Time : 2019/7/22 16:15 
@Author : 一条小咸鱼
@File : 
@Software: GoLand
*/
package utils

import (
	"ServerUpdate/model"
	"fmt"
	"testing"
)

func TestUploadFileRemote(t *testing.T) {
	sshConfig := model.SshConfig{}
	err := ReadConfigByObj("../config/config.toml", &sshConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	bean, err := GetSshBean(sshConfig.UserName, sshConfig.Password, sshConfig.HostName, sshConfig.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	sftpClient, err := bean.GetSftpConnect()
	if err != nil {
		fmt.Println(err)
		return
	}
	UploadFileRemote(sftpClient,"F:\\UpdateCfg.ini","/home/1.txt")

}