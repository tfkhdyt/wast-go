package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"wast-go/utils"
)

func main() {
	args := os.Args
	startTime := args[1]
	stopTime := args[2]
	startTimeInSeconds := utils.TimeStampToSeconds(startTime)
	stopTimeInSeconds := utils.TimeStampToSeconds(stopTime)
	fileName := args[3]
	splitCount := math.Ceil((float64(stopTimeInSeconds) - float64(startTimeInSeconds)) / 30)

	fmt.Printf("Start time\t: %v (%v)\n", startTime, startTimeInSeconds)
	fmt.Printf("Stop time\t: %v (%v)\n", stopTime, stopTimeInSeconds)
	fmt.Printf("Split count\t: %v \n", splitCount)
	fmt.Printf("File name\t: %v \n", fileName)
	fmt.Println("=====================================================")

	var wg sync.WaitGroup

	wg.Add(int(splitCount))

	for i := 0; i < int(splitCount); i++ {
		go utils.SplitVideo(fileName, startTimeInSeconds, stopTimeInSeconds, i, int(splitCount), &wg)
	}

	wg.Wait()
	fmt.Println("=====================================================")
	fmt.Println("Video has been trimmed successfully!")
}
