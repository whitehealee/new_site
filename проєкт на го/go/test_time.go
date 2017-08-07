package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now().Local()
	fmt.Println("The Current time is ", current_time.Format("2006-01-02 15:04:05.999"))
}