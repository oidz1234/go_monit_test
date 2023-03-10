// host_details.go
// gets the host details that we need
// uptime, ip for now

package monitors

import (
    "fmt"
    "github.com/shirou/gopsutil/v3/host"

)


func ParaTest() {
    host_id := host.HostID()
    fmt.Println(host_id)
}
