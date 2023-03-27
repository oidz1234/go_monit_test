package main

import (
    "fmt"
     "go_monitor/monitors"
     "time"
)



type mesure struct {
    heartbeat int64
    hostid string
    hostname string
    uptime uint64
    os string
    platform string
    ip string
    temp  []monitors.TemperatureReading
    load  map[string]float64
    disks map[string]float64
    memory float64
    upload uint64
    download uint64
    services map[string]string

}


func main() {

    defaultDisks := []string{"/", "/home"}
    defaultServices := []string{"sshd", "httpd"} // liunx defaults again can be configured


    


    interval := 1 // seconds to sleep between sending can be user configured evnetually


    for {

        loadmap := make(map[string]float64)
        diskmap := make(map[string]float64)
        servicemap := make(map[string]string)

        m := mesure{}
        heartbeat := time.Now().Unix()
        m.heartbeat = heartbeat
        //fmt.Printf("Heartbeat (unix)time %v\n", time.Now().Unix())

        m.hostid, m.hostname, m.uptime, m.os, m.platform, m.ip = monitors.GetHostDetails()
        // for some reason canne do this directly
        //temp := monitors.GetTemp()
        m.temp = monitors.GetTemp()

        m.load = monitors.GetLoad(loadmap)

        for _, disk := range defaultDisks {
            diskmap[disk] = monitors.GetDiskUsage(disk)
        }
        m.disks = diskmap
        m.memory = monitors.GetMem()
        m.upload, m.download = monitors.GetNetStats()
        // TODO:
        // should do the caculation (new up - old up)
        // to get amount sent in given timeframe
        // but that does not want to work
        // so will do it api side for now
        for _, service := range defaultServices {
            //fmt.Printf("Service %v %v\n", service, monitors.ServiceCheck(service))
            servicemap[service] = monitors.ServiceCheck(service)

        }
        m.services = servicemap
        // TODO:
        // Only if not windows (as no IOwait)
        // TODO: 
        // Find equivlent on windows
        // TODO:
        // IMplement correctly lol
        //monitors.GetIOWait()

        fmt.Printf("%+v\n", m)



        time.Sleep(time.Duration(interval) * time.Second)
    }


}

