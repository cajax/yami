package main

import (
	"encoding/json"
	"github.com/cajax/yami"
	"log"
	"time"
)

func main() {
	videoDetails("video.mp4")
	imageDetails("image.jpg")
}

func videoDetails(filename string){
	mediainfo, err := yami.GetMediaInfo(filename, 10*time.Second,"--Language=raw")

	if err != nil {
		log.Panic(err)
	}

	b, err := json.MarshalIndent(mediainfo, "", "  ")
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(b))

	video := mediainfo.GetFirstVideoTrack()
	audio := mediainfo.GetFirstAudioTrack()

	log.Printf("Video Codec:\t%s", video.CodecID)
	log.Printf("Framerate:\t\t%f", video.FrameRate)


	log.Printf("Audio Format:\t%s", audio.Format)
	log.Printf("Channels:\t\t%d", audio.Channels)
}

func imageDetails(filename string){
	mediainfo, err := yami.GetMediaInfo(filename, 10*time.Second,"--Language=raw")

	if err != nil {
		log.Panic(err)
	}

	b, err := json.MarshalIndent(mediainfo, "", "  ")
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(b))

	image := mediainfo.GetFirstImageTrack()

	log.Printf("Format:\t%s", image.Format)
	log.Printf("Width:\t%d", image.Width)
	log.Printf("Height:\t%d", image.Height)
}