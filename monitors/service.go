package monitors

import (
	"os/exec"
	"strings"
)

// TODO:
// this is for linux
// for windows, should be a check and then use something like the sc query command
// then write some logic in main to compare

func ServiceCheck(serviceName string) string {
	// Set the name of the service to check
	// Execute the `systemctl` command to get the service status
	cmd := exec.Command("systemctl", "status", serviceName)
	output, _ := cmd.CombinedOutput() // we don't error check here it fails if command returns not active
    //fmt.Println(string(output))

	// Parse the output to determine the service status
	status := "unknown"
	if strings.Contains(string(output), "Active: active") {
		status = "running"
	} else if strings.Contains(string(output), "Active: failed") {
		status = "failed"
	} else if strings.Contains(string(output), "Active: inactive") {
		status = "inactive"
	}

	// Print the service status
	//fmt.Printf("%s service status: %s\n", serviceName, status)
    return status
}
