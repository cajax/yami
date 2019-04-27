//Package yami is a wrapper for mediainfo cli tool.
//It provides simple access to media details
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

//SetMediainfoBinPath sets path to Mediainfo binary
func SetMediainfoBinPath(newBinPath string) {
	binPath = newBinPath
}

//GetMediaInfo executes mediainfo with specific time limit and returns parsed output.
// You may provide additional arguments like --Language=raw or SSL options
func GetMediaInfo(filePath string, timeout time.Duration, arg ...string) (mediaInfo *MediaInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return GetMediaInfoContext(ctx, filePath, arg...)
}

//GetMediaInfoContext executes mediainfo with given context and returns parsed output.
// You may provide additional arguments like --Language=raw or SSL options
func GetMediaInfoContext(ctx context.Context, filePath string, arg ...string) (mediaInfo *MediaInfo, err error) {

	Args := append([]string{"--Output=XML"}, arg...)
	Args = append(Args, filePath)

	cmd := exec.Command(
		binPath,
		Args...,
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
