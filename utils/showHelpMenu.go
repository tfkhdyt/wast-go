package utils

import (
	"fmt"

	"github.com/tfkhdyt/wast-go/constants"
)

func ShowHelpMenu() string {
	return fmt.Sprintf(`WhatsApp Status Trimmer (WAST) - Go [%v]
Simple program to trim/split a video into 30s videos for WhatsApp Status
	
Usage: 
	wast-go <start timestamp> <end timestamp> <input file>
	wast-go <options>

Example:
	wast-go 00:00:23 00:01:40 weightless.mp4

	The output would be 3 videos:
		1. 00:00:23 - 00:00:53
		2. 00:00:53 - 00:01:23
		3. 00:01:23 - 00:01:40

Options:
	-h, --help            Show help menu
	-v, --version         Show app version`, constants.Version)
}
