/*
@Time : 2019/7/23 13:42
@Author : 一条小咸鱼
@File :
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	var (
		activityDay int64 = 30
		cycleDay    int64 = 15
		daley       int64 = 3
		//2019-07-20 0:0:0
		//首次活动的开启时间
		firstActivity int64 = 1561286885
		DayToSec      int64 = 24 * 3600
	)

	fmt.Printf("开服时间:%s,\n", time.Unix(firstActivity-(daley*DayToSec), 0))
	fmt.Printf("首次活动的开启时间:%s,\n", time.Unix(firstActivity, 0))
	nowStamp := time.Now().Unix()
	fmt.Printf("当前时间:%s,当前时间戳:%d\n", time.Now(), time.Now().Unix())
	//acivitySec := activityDay * DayToSec
	//cycleSec := cycleDay * DayToSec
	sumSec := (activityDay + cycleDay) * DayToSec
	diffSec := nowStamp - firstActivity
	//residue := ( acivitySec - diffSec) /DayToSec+1

	//周期中的第几天.
	cycDay := (diffSec%sumSec)/DayToSec + 1

	fmt.Println("当前活动周期的天数-->", cycDay)
	if cycDay <= activityDay {
		fmt.Printf("活动第%d天\n", cycDay)
	} else {
		fmt.Println("-----活动结束---------")
		fmt.Printf("等待间隔第%d天\n", cycDay-activityDay)
	}
	unixStart := time.Unix(nowStamp-(diffSec%sumSec), 0)
	fmt.Println("当前活动开启的时间", unixStart)

	unix := time.Unix(nowStamp+sumSec-(diffSec%sumSec), 0)
	fmt.Println("下一个活动开启的时间", unix)

}
