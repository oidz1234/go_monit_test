package main

import (
    "fmt"
     "go_monitor/monitors"
     "time"
     "encoding/json"
)



type mesure struct {
    Heartbeat int64
    Hostid string
    Hostname string
    Uptime uint64
    Os string
    Platform string
    Ip string
    Temp  []monitors.TemperatureReading
    Load  map[string]float64
    Disks map[string]float64
    Memory float64
    Upload uint64
    Download uint64
    Services map[string]string

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
        m.Heartbeat = heartbeat
        //fmt.Printf("Heartbeat (unix)time %v\n", time.Now().Unix())

        m.Hostid, m.Hostname, m.Uptime, m.Os, m.Platform, m.Ip = monitors.GetHostDetails()
        // for some reason canne do this directly
        //temp := monitors.GetTemp()
        m.Temp = monitors.GetTemp()

        m.Load = monitors.GetLoad(loadmap)

        for _, disk := range defaultDisks {
            diskmap[disk] = monitors.GetDiskUsage(disk)
        }
        m.Disks = diskmap
        m.Memory = monitors.GetMem()
        m.Upload, m.Download = monitors.GetNetStats()
        // TODO:
        // should do the caculation (new up - old up)
        // to get amount sent in given timeframe
        // but that does not want to work
        // so will do it api side for now
        for _, service := range defaultServices {
            //fmt.Printf("Service %v %v\n", service, monitors.ServiceCheck(service))
            servicemap[service] = monitors.ServiceCheck(service)

        }
        m.Services = servicemap
        // TODO:
        // Only if not windows (as no IOwait)
        // TODO: 
        // Find equivlent on windows
        // TODO:
        // IMplement correctly lol
        //monitors.GetIOWait()



        jsonBytes, err := json.Marshal(m)

        if err != nil {
            panic(err)
        }
        fmt.Println(string(jsonBytes))


        time.Sleep(time.Duration(interval) * time.Second)
    }


}

