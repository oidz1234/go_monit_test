package main

import (
    "fmt"
     "go_monitor/monitors"
     "time"
)

func main() {

    defaultDisks := []string{"/", "/home"}
    defaultServices := []string{"sshd", "httpd"} // liunx defaults again can be configured


    interval := 1 // seconds to sleep between sending can be user configured evnetually


    for {


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
        fmt.Println(monitors.GetMem())
        up, down := monitors.GetNetStats()
        // TODO:
        // should do the caculation (new up - old up)
        // to get amount sent in given timeframe
        // but that does not want to work
        // so will do it api side for now
        fmt.Printf("Total uploaded %v\n", up)
        fmt.Printf("Total download %v\n", down)
        for _, service := range defaultServices {
            fmt.Printf("Service %v %v\n", service, monitors.ServiceCheck(service))
        }




        time.Sleep(time.Duration(interval) * time.Second)
    }


}

