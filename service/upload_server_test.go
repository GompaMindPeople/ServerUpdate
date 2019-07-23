/*
@Time : 2019/7/23 11:21 
@Author : 一条小咸鱼
@File : 
@Software: GoLand
*/
package service

import (
	"ServerUpdate/model"
	"ServerUpdate/utils"
	"testing"
)




func TestCreateRemoteDir(t *testing.T) {
	sshConfig := model.SshConfig{}
	_ = utils.ReadConfigByObj("config/config.toml", &sshConfig)
	bean, _ := utils.GetSshBean(sshConfig.UserName, sshConfig.Password, sshConfig.HostName, sshConfig.Port)
	defer bean.Client.Close()
	//strings := model.TomlConfig{}
	//获取 sftp 实例
	sftpClient, _ := bean.GetSftpConnect()
	defer sftpClient.Close()
	//localPath string,remotePath string
	CreateRemoteDir(sftpClient)
}



