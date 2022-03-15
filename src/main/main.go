package main

import (
	"fmt"
	"go-dev/common/yaml"
)

func main() {
	yaml.LoadYaml()
	fmt.Println("hello world")
}
