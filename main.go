package main

import (
    "fmt"
     "go_monitor/monitors"
     "time"
     "encoding/json"
     "net/http"
     "bytes"
)



type Custom struct {
    Disks []string
    Services []string
}

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

    //token := "b4cb8ca211897248be354ce5df0a75607ce8113c"
    token := "2777e820ed347af37dc7d066d9169a30c127baa3"
    authHeader := "token " + token
    endpoint := "http://localhost:8000/api/update/"

    client := &http.Client{}


    // TODO:
    // these should be check on boot from the "config" api 
    // this api does not yet exist
    defaultDisks := []string{"/", "/home"}
    defaultServices := []string{"sshd", "httpd"} // liunx defaults again can be configured


    


    interval := 15// seconds to sleep between sending can be user configured evnetually


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

        req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBytes))
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", authHeader)

        resp, err := client.Do(req)
        if err != nil {
            panic(err)
            // TODO: if panic here, just sleep and try again (if api serv is down)
        }
        defer resp.Body.Close()


        /*
        // Read the response body
        var responseMap map[string]interface{}
        err = json.NewDecoder(resp.Body).Decode(&responseMap)
        if err != nil {
            panic(err)
        }
        */

        var custom Custom
        err = json.NewDecoder(resp.Body).Decode(&custom)
        if err != nil {
            panic(err)
        }
        if custom.Disks != nil {
            defaultDisks = custom.Disks
        }
        if custom.Services != nil {
            defaultServices = custom.Services
        }

        fmt.Println("ran succ")





        time.Sleep(time.Duration(interval) * time.Second)
    }


}

