package main

import (
    "fmt"
     "go_monitor/monitors"
     "time"
)

func main() {

    defaultDisks := []string{"/", "/home"}

    fmt.Println("Hello, World!")
    fmt.Printf("Heartbeat (unix)time %v\n", time.Now().Unix())

    monitors.Test()
    monitors.GetHostDetails()
    monitors.GetTemp()
    monitors.GetLoad()
    fmt.Println(defaultDisks)
    for _, disk := range defaultDisks {
        fmt.Printf("%v has used: ", disk)
        fmt.Println(monitors.GetDiskUsage(disk))
    }
}

