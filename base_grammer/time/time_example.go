package main

import (
	"fmt"
	"time"
)

func main() {
	parse, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-04-14 12:00:00", time.Local)
	if err != nil {
		return
	}
	fmt.Println(parse)
}
