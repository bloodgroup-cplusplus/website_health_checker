package main 

import (
	"fmt"
	"net"
	"time"
)


// this is called from the main funciton
func Check(destination string, port string) {
	address := destination + ":" + port 
	timeout : time.Duration(5 *time.Second)
	// within 5 seconds if we get a response back from the website is up and running else there is an error 
	conn,err := net.DialTimeout("tcp", address,timeout)

	var status string 

	if err !=nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error :%v",destination,err)
	}

	else {
		status = fmt.Sprintf("[UP] %v is reachable, \n from :%v \n To :%v",destination,conn.LocalAddr(),conn.RemoteAddr())
	}
	return status 
}
