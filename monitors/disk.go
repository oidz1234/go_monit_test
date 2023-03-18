// disk.go
// gets disk info

package monitors

import (
    "fmt"
    "github.com/shirou/gopsutil/v3/disk"
)


func GetDefaultDisks(amount int) {
    // get list of all disks ✅
    // figure out total space of all disks
    // figure out the top X by space
    // return them
    parts, err := disk.Partitions(false)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(parts)

}



func GetDiskUsage(diskPath string) float64 {

    diskStat, err := disk.Usage(diskPath)
    if err != nil {
        fmt.Println(err)
    }
    return diskStat.UsedPercent
}

