package photo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pi_go_2/camutils"
	"pi_go_2/config"
	"time"
)

//将所有打开的摄像头在当前时刻的摄像发送到服务器
func PostAllImage(state string) {
	//发送多个image（两个）
	for i := 0; i < len(camutils.VideoHandlers); i++ { // 开启的摄像头数目
		// 封装图像
		videohandler := camutils.VideoHandlers[i]
		information := map[string]string{
			"machineid": config.Config.MachineId,
			"sequence":  videohandler.Prefix,
			"state":     state,
		}
		formData := createFormdata(information, videohandler)
		contentType, bodyBuffer, _ := createFormDataFromBytes(formData)
		response, err := http.Post(config.Config.DetectUrl, contentType, bodyBuffer)
		if err != nil {
			panic(err.Error())
		}
		respBody, err := ioutil.ReadAll(response.Body)
		fmt.Println(string(respBody))
		response.Body.Close()
	}
}

// 构造远程服务器检测的数据图像结构
func createFormdata(information map[string]string, videohandler camutils.VideoCap) PostBytes {
	filemap := make([]ImageBytes, 0)
	imageBytes, err := videohandler.GetJpegImageBytes()
	if err != nil {
		log.Print(videohandler.Prefix + "\terror")
		log.Fatal(err.Error())
	}
	imageByteFile := ImageBytes{
		FileName:    getTimestr() + videohandler.Prefix + ".jpg",
		FieldName:   "image",
		Content:     imageBytes,
		ContentType: "image/jpeg",
	}
	filemap = append(filemap, imageByteFile)
	// 把文件流跟附加属性字段都放入form 中,类似 curl -F 中多个属性
	formData := PostBytes{
		FileMap:  filemap,
		FieldMap: information,
	}
	return formData
}

func getTimestr() string {
	timeStr := time.Now().Format("2006-01-02,15:04:05") //当前时间的字符串，2006-01-02 15:04:05(06年1月2号下午3点4分5秒，06 12345 不重复)，固定写法
	return timeStr
}
