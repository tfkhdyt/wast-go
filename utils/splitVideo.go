package utils

import (
	"fmt"
	"path"
	"strings"
	"sync"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func SplitVideo(videoFile string, startTime int, stopTime int, count int, splitCount int, wg *sync.WaitGroup) {
	videoFileNameWithoutExt := strings.TrimSuffix(videoFile, path.Ext(videoFile))
	outputFileName := fmt.Sprintf("%s-wa-%d%s", videoFileNameWithoutExt, count+1, path.Ext(videoFile))
	ss := startTime + 30*count
	var t int

	if count == splitCount-1 {
		t = stopTime - (startTime + 30*count)
	} else {
		t = 30
	}

	err := ffmpeg.Input(videoFile, ffmpeg.KwArgs{"ss": ss}).
		Output(outputFileName, ffmpeg.KwArgs{"t": t}).OverWriteOutput().Run()
	if err != nil {
		panic(err.Error())
	}

	wg.Done()
}
