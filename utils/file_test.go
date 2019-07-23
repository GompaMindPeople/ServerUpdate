/*
@Time : 2019/7/23 11:09 
@Author : 一条小咸鱼
@File : 
@Software: GoLand
*/
package utils

import (
	"fmt"
	"testing"
)

func TestGetAllDir(t *testing.T) {

	var list  []string
	strings, e := GetAllDir(list, "E:\\server\\app")
	if e != nil{
		t.Error(e)
		return
	}

	for  _,v := range strings{
		fmt.Println(v)
	}


}



