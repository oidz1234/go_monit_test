package main

import (
    "fmt"
     "go_monitor/monitors"
)

func main() {
    fmt.Println("Hello, World!")

    monitors.Test()
    monitors.GetHostDetails()
    monitors.GetTemp()
}
