package main

import (
	"fmt"
	"os"
	"pi_go_2/config"
	"pi_go_2/photo"
	"time"
)

func main() {
	// 等待 5s 等所有摄像头都打开并开始拍照
	time.Sleep(10 * time.Second)
	count := 0
	for {
		if photo.GetCapStartPin() == 0 && photo.GetCapEndPin() == 1 {
			count++
		}else if photo.GetCapEndPin() == 0 && photo.GetCapStartPin() == 1 {
			count--
		}
		if count > config.Config.Delay {
			fmt.Println("start...")
			count = 0
			photo.PostAllImage("start")
			// 每次发完照片，等待 0.5 秒钟
			time.Sleep(500 * time.Millisecond)
		}
		if count < - config.Config.Delay{
			fmt.Println("end...")
			count = 0
			photo.PostAllImage("end")
			// 每次发完照片，等待 0.5 秒钟
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func testCap() {
	for order := 0; ; order = 0 {
		fmt.Println("Please input the order: ")
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println("|   1. Make all cap take pictures and upload --start  |")
		fmt.Println("|   2. Make all cap take pictures and upload --end    |")
		fmt.Println("|   0. Exit                                           |")
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Scanln(&order)
		switch order {
		case 1:
			photo.PostAllImage("start")
		case 2:
			photo.PostAllImage("end")
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Input error!!! Please input again")
		}
	}
}
