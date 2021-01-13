package camutils

import (
	"fmt"
	"gocv.io/x/gocv"
	"log"
	"pi_go_2/config"
	"sync"
)

type VideoCap struct {
	videoId int
	img gocv.Mat
	mutex sync.Mutex
	Prefix string
}

var VideoHandlers []VideoCap

func init()  {
	InitAndStartCap()
}

func InitAndStartCap()  {
	VideoHandlers = make([]VideoCap,0)
	for i:=0;i<len(config.Config.CapConfigs);i++{
		videoHandler := VideoCap{
			videoId:config.Config.CapConfigs[i].VideoId,
			img:gocv.NewMat(),
			mutex:sync.Mutex{},
			Prefix:config.Config.CapConfigs[i].Prefix,
		}
		go videoHandler.StartCap()
		// ToDo 加入VideoHandler 前， 存储的 img 不能为空图像
		VideoHandlers = append(VideoHandlers,videoHandler)
	}

}
func (cap *VideoCap) GetJpegImageBytes() (buf []byte, err error) {
	cap.mutex.Lock()
	if cap.img.Empty() {
		fmt.Println(cap.videoId)
		fmt.Println("sorry the img of cap is not contained image")
	}
	gocv.IMWrite("404.jpg", cap.img)

	imageBytes,err := gocv.IMEncode(".jpg",cap.img)
	if err!=nil{
		log.Print("cap  "+string(cap.videoId)+" error")
		log.Fatal(err.Error())
	}
	cap.mutex.Unlock()
	return imageBytes,err
}

func (cap *VideoCap) StartCap()  {
	cam_handler,err:=gocv.OpenVideoCapture(cap.videoId)
	fmt.Println("videoId:\t", cap.videoId)
	log.Print(cam_handler.Get(gocv.VideoCaptureFrameHeight))
	log.Print(cam_handler.Get(gocv.VideoCaptureFrameWidth))
	cam_handler.Set(gocv.VideoCaptureFrameHeight,1080)
	cam_handler.Set(gocv.VideoCaptureFrameWidth,1920)
	log.Print(cam_handler.Get(gocv.VideoCaptureFrameHeight))
	log.Print(cam_handler.Get(gocv.VideoCaptureFrameWidth))

	if err!=nil{
		panic(err.Error())
	}
	for{
		cap.mutex.Lock()
		cam_handler.Read(&cap.img)
		cap.mutex.Unlock()
	}
	defer cam_handler.Close()
}
