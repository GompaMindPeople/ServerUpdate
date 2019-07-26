/*
@Time : 2019/7/22 14:49
@Author : 一条小咸鱼
@File :
@Software: GoLand
*/
package service

import (
	"ServerUpdate/utils"
	"fmt"
	"github.com/pkg/sftp"
	"strings"
)

func UploadLocalByRemote(s *sftp.Client, localPath string, remotePath string) {
	var (
		local []string
		err   error
	)

	if utils.IsDir(localPath) {
		//如果存在的话,创建远程文件夹
		CreateRemoteDir(s, localPath, remotePath)
		//如果是文件夹的话,则获取该文件夹下的所有文件
		local, err = utils.GetAllFiles(localPath)
	} else {
		sendFile(s, localPath, localPath, remotePath)
		return
	}
	if err != nil {
		fmt.Println(err)
	}
	//发送文件夹下的所有文件路径
	for _, file := range local {
		sendFile(s, file, localPath, remotePath)
	}
}

//构造远程的目录
func CreateRemoteDir(s *sftp.Client, localPath string, remotePath string) {
	var pathList []string
	dir, e := utils.GetAllDir(pathList, localPath)
	if e != nil {
		return
	}
	for _, v := range dir {
		path := whitPath(v, localPath, remotePath)
		e := s.MkdirAll(path)
		if e != nil {
			fmt.Println("构造远程目录失败--->", e, ",失败的目录-->", path)
		}
	}

}

//处理 地址 并发送 数据到 远程服务器
func sendFile(s *sftp.Client, file, localPath, remotePath string) {
	remote := whitPath(file, localPath, remotePath)
	fmt.Println("UploadLocalByRemote-->", remote)
	utils.UploadFileRemote(s, file, remote)
}

// 将 本地的路径转换成 linux 路径
func whitPath(file, localPath, remotePath string) string {
	file1 := strings.Replace(file, localPath, "", -1)
	file1 = strings.Replace(file1, "\\", "/", -1)
	return remotePath + file1
}

//func TraverseForFolder(fileName string){
//
//	info, err := os.Stat(fileName)
//	if err != nil{
//		fmt.Println("打开路径错误:",err,",错误的路径:",fileName)
//		return
//	}
//	if info.IsDir() {
//		infos, err := ioutil.ReadDir(fileName)
//		if err != nil{
//			fmt.Println("读取文件夹的时候错误:",err)
//			return
//		}
//		for  temp := range infos  {
//
//		}
//		TraverseForFolder()
//	}
//
//}
