package main

import (
  "fmt"
  "net"
  "os"
  "time"
)

func main() {
  start := time.Now()
  conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]), time.Duration(5 * time.Second))
   if err != nil {
	   os.Stderr.WriteString(fmt.Sprintf("failed to connect: %s", err.Error()))
	   os.Exit(1)
	}
  defer conn.Close()
  elapsed := time.Since(start)
  fmt.Printf("Connection to %s:%s took %s\n", os.Args[1], os.Args[2], elapsed)
}
