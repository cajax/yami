package main

import (
	"encoding/json"
	"github.com/cajax/yami"
	"log"
	"time"
)

func main() {
	mediainfo, err := yami.GetMediaInfo("video.mp4", 10*time.Second,"--Language=raw")

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
