
// temp.go
// gets temperature of stuff

package monitors

import (
    "fmt"
    "github.com/shirou/gopsutil/v3/load"
)

func GetLoad() {

 	load, err := load.Avg()
	if err != nil {
		fmt.Println(err)
	}

    //fmt.Println(load.Load1)
    fmt.Println(load)
    // Example on how to get specifc value (loop over it durrr)
}
/*
func GetPs() {

 	ps, err := load.Misc()
	if err != nil {
		fmt.Println(err)
	}

    //fmt.Println(load.Load1)
    fmt.Println(ps)
    // Example on how to get specifc value (loop over it durrr)
}
*/
