package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func TimeStampToSeconds(timeStamp string) int {
	if len(timeStamp) != 8 || !strings.Contains(timeStamp, ":") || len(strings.Split(timeStamp, ":")) != 3 {
		panic(fmt.Sprintf("Timestamp is not valid: %v", timeStamp))
	}
	timeStampArr := strings.Split(timeStamp, ":")
	timeStampArrNum := make([]int, 3)

	for idx, el := range timeStampArr {
		elInt, err := strconv.Atoi(el)
		if err != nil {
			panic(err.Error())
		}
		timeStampArrNum[idx] = elInt
	}

	seconds := timeStampArrNum[0]*3600 + timeStampArrNum[1]*60 + timeStampArrNum[2]
	return seconds
}
