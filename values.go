package loki

import (
	"strconv"
)

type Values [][]string

func getCurrentTime() string {
	return strconv.Itoa(int(CurrentTimeInNs()))
}

func NewValues() Values {
	return [][]string{}
}

func (vs *Values) AddValue(timestamp string, message string) {
	*vs = append(*vs, []string{timestamp, message})
}

func (vs *Values) AddValueWithCurrentTime(message string) {
	vs.AddValue(getCurrentTime(), message)
}
