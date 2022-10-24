package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("hello world")
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE"})
	t.AppendRow(table.Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"})
	fmt.Println(t.Render())
}
