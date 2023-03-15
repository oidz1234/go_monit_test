package main

import (
    "fmt"
     "go_monitor/monitors"
     "time"
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Printf("Heartbeat (unix)time %v\n", time.Now().Unix())

    monitors.Test()
    monitors.GetHostDetails()
    monitors.GetTemp()
    monitors.GetLoad()
}

