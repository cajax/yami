package yami

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"os/exec"
	"time"
)

var (
	// ErrBinNotFound is returned when the mediainfo binary was not found
	ErrBinNotFound = errors.New("mediainfo bin not found")
	// ErrTimeout is returned when the mediainfo process did not succeed within the given time
	ErrTimeout = errors.New("process timeout exceeded")

	binPath = "mediainfo"
)


func SetMediainfoBinPath(newBinPath string) {
	binPath = newBinPath
}


func GetMediaInfo(filePath string, timeout time.Duration) (mediaInfo *MediaInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return GetMediaInfoContext(ctx, filePath)
}

func GetMediaInfoContext(ctx context.Context, filePath string) (mediaInfo *MediaInfo, err error) {
	cmd := exec.Command(
		binPath,
		"--Output=XML",
		filePath,
	)

	var outputBuf bytes.Buffer
	cmd.Stdout = &outputBuf

	err = cmd.Start()
	if err == exec.ErrNotFound {
		return nil, ErrBinNotFound
	} else if err != nil {
		return nil, err
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-ctx.Done():
		err = cmd.Process.Kill()
		if err == nil {
			return nil, ErrTimeout
		}
		return nil, err
	case err = <-done:
		if err != nil {
			return nil, err
		}
	}

	mediaInfo = &MediaInfo{}
	err = xml.Unmarshal(outputBuf.Bytes(), mediaInfo)
	if err != nil {
		return mediaInfo, err
	}

	return mediaInfo, nil
}
