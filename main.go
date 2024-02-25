package main

import (
    "bufio"
    "os"
    "fmt"
    "net"
    "strings"
    "time" //Imported But Not Used Yet
)

const (
    PORT_STATUS_FILTERED           = "filtered"        // Port is filtered (i/o timeout)
    PORT_STATUS_CLOSED             = "closed"          // Port is closed (connection refused)
    PORT_STATUS_OPEN               = "open"            // Port is open
    PORT_STATUS_OPEN_OR_FILTERED   = "open|filtered"   // Port is open|filtered (UDP: no response)
    PORT_STATUS_CLOSED_OR_FILTERED = "closed|filtered" // Future
)

var (
    TcpTimeout = 1 * time.Second
)

func StandardConnect(host string, port uint16) (string, error) {
    // TCP connection attempt, if connection is successful, returns timeout
    // If the connection cannot be made, returns null with nil val
    if conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), TcpTimeout); err != nil {
      // connection error handling by checking error messages
      // reattempts are made in cases where port status is closed or filtered
      // else returns empty val
        if strings.Contains(err.Error(), "connection refused") {
            return PORT_STATUS_CLOSED, nil
        } else if strings.Contains(err.Error(), "i/o timeout") {
            return PORT_STATUS_FILTERED, nil
        } else if strings.Contains(err.Error(), "too many open files") {
            time.Sleep(100 * time.Millisecond)
            return StandardConnect(host, port)
        } else {
            return "", err
        }
    } else {
        conn.Close()
        return PORT_STATUS_OPEN, nil
    }
}
// Calculates progress bar status
func printProgressBar(iteration, total int, prefix, suffix string, length int, fill string) {
 percent := float64(iteration) / float64(total)
 filledLength := int(length * iteration / total)
 end := ">"

 if iteration == total {
  end = "="
 }
 bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))
 fmt.Printf("\r%s [%s] %f%% %s", prefix, bar, percent, suffix)
 if iteration == total {
  fmt.Println()
 }
}

// saved ascii art to print
func printAsciiArt() {
  fmt.Println(
    `
              _                                   
             (_)                                  
         __ _ _  __ _  __ _ _ __ ___   __ _ _ __  
        / _  | |/ _  |/ _  | '_   _ \ / _  | '_ \ 
       | (_| | | (_| | (_| | | | | | | (_| | |_) |
        \__, |_|\__, |\__,_|_| |_| |_|\__,_| .__/ 
         __/ |   __/ |                     | |    
        |___/   |___/                      |_|    
          ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
          ⠀⠀⠀⠀⠀⣀⣠⣤⣶⣞⡛⠿⠭⣷⡦⢬⣟⣿⣿⣷⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀
          ⠀⠀⠀⢠⡾⣯⡙⠳⣍⠳⢍⡙⠲⠤⣍⠓⠦⣝⣮⣉⠻⣿⡄⠀⠀⠀⠀⠀⠀⠀
          ⠀⠀⠀⡿⢿⣷⣯⣷⣮⣿⣶⣽⠷⠶⠬⠿⠷⠟⠻⠟⠳⠿⢷⡀⠀⠀⠀⠀⠀⠀
          ⠀⠀⣸⣁⣀⣈⣛⣷⠀⢹⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢧⠀⠀⠀⠀⠀⠀
          ⠀⣸⣯⣁⣤⣤⣀⠈⢧⠘⣆⢀⣠⠤⣄⠀⠀⠀⠀⠀⠀⠀⠀⠘⡇⠀⠀⠀⠀⠀
          ⠀⢙⡟⡛⣿⣿⣿⢷⡾⠀⢿⣿⣏⣳⣾⡆⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⠀⠀⠀
          ⢀⡞⠸⠀⠉⠉⠁⠀⠀⣠⣼⣿⣿⠀⣽⡇⠀⠀⠀⠀⠀⠀⠀⡼⠁⠀⠀⠀⠀⠀
          ⣼⡀⣀⡐⢦⢀⣀⠀⣴⣿⣿⡏⢿⣶⣟⣀⣀⣀⣀⣀⣤⣤⠞⠁⠀⠀⠀⠀⠀⠀
          ⠀⣿⣿⣿⣿⣾⣿⣿⣿⣾⡻⠷⣝⣿⡌⠉⠉⠁⠀⠀⣸⠁⠀⠀⠀⠀⠀⠀⠀⠀
          ⠀⠈⢻⣿⣿⣿⣿⡟⣿⣟⠻⣿⡿⢿⡇⠀⠀⠀⠀⠀⢹⠀⠀⠀⠀⠀⠀⠀⠀⠀
          ⠀⢠⣿⢿⣼⣿⣿⠿⣏⣹⡃⢹⣯⡿⠀⠀⠀⠀⠀⠀⠈⢧⠀⠀⠀⠀⠀⠀⠀⠀
          ⠀⣽⣿⣿⢿⠹⣿⣇⠿⣾⣷⣼⠟⠁⠀⠀⠀⢀⣠⣴⣶⣾⣷⣶⣄⡀⠀⠀⠀⠀
          ⠀⢿⣾⡟⢺⣧⣏⣿⡷⢻⠅⠀⠀⠀⢀⣠⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⡀⠀
          ⠀⠀⠙⠛⠓⠛⠋⣡⣿⣬⣤⣤⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠟⠛⠛
          ⠀⠀⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠟⠋⠉⠁⠀⠀⠀⠀⠀⠀
          ⠀⠀⠀⠀⠀⠀⠸⡿⠿⠿⠿⠿⠿⠿⠟⠛⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
    `)
}
         
func main() { 
  printAsciiArt()
  reader := bufio.NewReader(os.Stdin)
  fmt.Printf("Hostname to scan (example.com): ")
  host, _ := reader.ReadString('\n')  
  host = strings.Trim(host, "\n")
    for i := 0; i < 30; i++ {
      time.Sleep(500 * time.Millisecond) // mimics work
      printProgressBar(i+1, 30, "Progress", "Complete", 25, "=")
     }
    ports := []uint16{22, 80, 443}
    for _, port := range ports {
        if status, err := StandardConnect(host, port); err == nil {
            fmt.Printf("%s:%d -> %s\n", host, port, status)
        }
    }
}