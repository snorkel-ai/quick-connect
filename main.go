package main

import (
  "net"
  "os"
  "time"
)

func main() {
   _, err := net.DialTimeout("tcp", os.Args[1] + ":" + os.Args[2]), time.Duration(5 * time.Second))
   if err != nil {
	   os.Stderr.WriteString("failed to connect: " + err.Error()))
	   os.Exit(1)
	}
}
