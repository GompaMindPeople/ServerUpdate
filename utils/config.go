/*
@Time : 2019/7/22 13:33 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package utils

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)



/**
	读取ioml文本的配置文件,转换成一个结构体
	@fileName 配置文件的路径
	@configObj 配置文件所映射的对象,需要传入指针数据
 */
func ReadConfigByObj(fileName string,configObj interface{}) ( err error) {
	var (
		fp       *os.File
		fcontent []byte
	)

	if fp, err = os.Open(fileName); err != nil {
		fmt.Println("打开配置文件错误 ", err)
		return
	}

	if fcontent, err = ioutil.ReadAll(fp); err != nil {
		fmt.Println("读取配置错误 error ", err)
		return
	}

	if err = toml.Unmarshal(fcontent, configObj); err != nil {
		fmt.Println("toml.解析错误 error ", err)
		return
	}
	return
}
