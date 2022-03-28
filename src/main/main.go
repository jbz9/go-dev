package main

import (
	"fmt"
	"go-dev/common/yaml"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	yaml.LoadYaml()
	fmt.Println("hello world")
}
