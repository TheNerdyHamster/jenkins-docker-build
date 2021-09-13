package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "9.9.9.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey there, my internal ip is %s\n", GetOutboundIP().String())
}

func main() {
	port := fmt.Sprintf("%s%d", ":", 8080)
	http.HandleFunc("/", handler)
	log.Printf("HTTP Server starting on port %v", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
