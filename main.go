package main

import (
	"fmt"
	"math"
	"os"
	"sync"

	"github.com/tfkhdyt/wast-go/constants"
	"github.com/tfkhdyt/wast-go/utils"
)

func main() {
	defer utils.HandlePanic()

	args := os.Args[1:]

	if len(args) == 1 {
		versionFlags := []string{"-v", "--version"}
		helpFlags := []string{"-h", "--help"}

		for _, flag := range versionFlags {
			if flag == args[0] {
				fmt.Println(constants.Version)
				return
			}
		}

		for _, flag := range helpFlags {
			if flag == args[0] {
				fmt.Println(utils.ShowHelpMenu())
				return
			}
		}

		panic(fmt.Sprintf("Option is not valid: %s", args[0]))
	} else if len(args) != 3 {
		panic("Some arguments are missing!")
	}

	startTime := args[0]
	stopTime := args[1]

	startTimeInSeconds := utils.TimeStampToSeconds(startTime)
	stopTimeInSeconds := utils.TimeStampToSeconds(stopTime)

	if startTimeInSeconds > stopTimeInSeconds {
		panic("Start time shouldn't after stop time")
	} else if startTimeInSeconds == stopTimeInSeconds {
		panic("Start time shouln't be the same as stop time")
	}

	fileName := args[2]
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
