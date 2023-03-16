// disk.go
// gets disk info

package monitors

import (
    "fmt"
    "github.com/shirou/gopsutil/v3/disk"
)



func GetDiskUsage(diskPath string) float64 {

    diskStat, err := disk.Usage(diskPath)
    if err != nil {
        fmt.Println(err)
    }
    return diskStat.UsedPercent
}

