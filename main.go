package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

// https://play.golang.org/p/BDt3qEQ_2H
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func handler(w http.ResponseWriter, r *http.Request) {
	ip, err := externalIP()
	if err != nil {
		fmt.Println("Failed to get externalIP:", err)
		ip = "Unknown"
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Failed to get os.Hostname:", err)
		hostname = "Unknown"
	}

	fmt.Fprintf(w, `
Hello, Container World v1.2!
	
Server Hostname: %s
Server IP:       %s
Remote Addr:     %s
X-Forwarded-For: %s
`,
		hostname,
		ip,
		r.Header["X-Forwarded-For"],
		r.RemoteAddr)
}

func main() {
	listenAddr := "0.0.0.0:8080"

	log.Println("Starting HTTP server on", listenAddr)
	http.HandleFunc("/", handler)
	http.ListenAndServe(listenAddr, nil)
}
