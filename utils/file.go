/*
@Time : 2019/7/22 15:16 
@Author : 一条小咸鱼
@File : 
@Software: GoLand
*/
package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)






/**
	打开一个文件夹,并返回其中所有遍历过的文件路径
 */
func GetLocalFileList(fileName string)(files []string){
	info, err := os.Stat(fileName)

	if err != nil{
		fmt.Println("打开路径的时候发生错误:->",err,"打开路径:->",fileName)
	}
	if info.IsDir(){
		 _, err = GetFilesAndDirs(files,fileName)
	}else{
		fmt.Println("打开的并不是一个文件夹.-->",fileName)
		return nil
	}

	if err != nil{
		fmt.Println("递归文件夹时,发生错误:->",err,"递归的路径:->",fileName)
		return nil
	}

	return files
}



func IsDir(pathName string)bool{
	info, _ := os.Stat(pathName)
	return  info.IsDir()
}



//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	//dirs, err := GetAllDir(dirPth)
	//if err != nil{
	//	fmt.Println("遍历本地的文件夹错误:->",dirPth)
	//	return
	//}
	//// 读取子目录下文件
	//for _, table := range dirs {
	//	temp, _ := GetAllFiles(table)
	//	for _, temp1 := range temp {
	//		files = append(files, temp1)
	//	}
	//}
	err = filepath.Walk(dirPth, func(path string, f os.FileInfo, err error) error {
		if f == nil{
			return err
		}
		if f.IsDir() {

			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return files, nil
}

func GetAllDir(dirs []string , dirPth string)( []string,  error){
	//temp :=  make([]string,0)
	if len(dirs) <1{
		dirs = append(dirs, dirPth)
	}
	PthSep := string(os.PathSeparator)
	////suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	 dir, err := ioutil.ReadDir(dirPth)
	 if err != nil{
	 	fmt.Println("遍历目录时发生错误!-->",err)
	 	return nil,err
	 }
	//
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			//dirList = append(dirList, dirPth+PthSep+fi.Name())
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			dirs, _ = GetAllDir(dirs, dirPth+PthSep+fi.Name())
		}
	}
	return dirs,nil
}


//获取指定目录下的所有文件和目录
func GetFilesAndDirs(files []string,dirPth string) (dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return  nil, err
	}
	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			 GetFilesAndDirs(files,dirPth + PthSep + fi.Name())
		} else {
				files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return  dirs, nil
}



