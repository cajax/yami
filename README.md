# Yet Another MediaInfo Go Wrapper

This wrapper allows you to run [mediainfo](https://mediaarea.net/en/MediaInfo) and parse its output in usable manner.

Yami is inspired of [go-ffprobe](https://github.com/vansante/go-ffprobe) and can be used as a drop-in replacement with minor changes in mapping.

Current version mapping matches [XML v2](https://mediaarea.net/mediainfo/mediainfo_2_0.xsd) 


### Example

```go
package main

import (
	"github.com/cajax/yami"
	"log"
	"time"
)

func main() {
	mediainfo, err := yami.GetMediaInfo("video.mp4", 10*time.Second)

	if err != nil {
		log.Panic(err)
	}

	video := mediainfo.GetFirstVideoTrack()

	log.Printf("Video Codec:\t%s", video.CodecID)
	log.Printf("Framerate:\t\t%f", video.FrameRate)
}

```
