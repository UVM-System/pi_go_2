package main

import (
	"fmt"
	"net/http"
	"os"
	"pi_go_2/photo"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 等待 5s 等所有摄像头都打开并开始拍照
	time.Sleep(10 * time.Second)
	router := gin.Default()
	router.POST("/photo", func(context *gin.Context) {
		state := context.PostForm("state")
		fmt.Println("state: ", state)
		photo.PostAllImage(state)
		context.JSON(http.StatusOK, gin.H{
			"status":  "SUCCESS",
			"message": "pi_go_2 finish",
		})
	})
	router.Run(":8000")
}

func test() {
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
