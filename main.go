package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	ip, _ := getIP(r)

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Parseform() error: %v", err)
		log.Fatal(fmt.Sprintf("Form Issue: User Agent: %s, IP: %s", r.UserAgent(), ip))
		return
	}
	if r.URL.Path != "/form" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		log.Fatal(fmt.Sprintf("404:User Agent: %s, IP: %s", r.UserAgent(), ip))
		return
	}
	if r.Method != "POST" && r.Method != "GET" {
		http.Error(w, "Method is Not Supported", http.StatusMethodNotAllowed)
		log.Fatal(fmt.Sprintf("METHOD NOT SUPPORTED: User Agent: %s, IP: %s", r.UserAgent(), ip))
		return
	}
	log.Println(fmt.Sprintf("HERE User Agent: %s, IP: %s", r.UserAgent(), ip))
	log.Println("POST REQUEST SUCCESSFUL")
	name := r.FormValue("name")
	address := r.FormValue("address")
	phone_number := r.FormValue("phone-number")
	dob := r.FormValue("dob")
	fmt.Fprintf(w, "Name:%s \nAddress: %s \nPhone Number: %v \nDOB:%v", name, address, phone_number, dob)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ip, _ := getIP(r)
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		log.Fatal(fmt.Sprintf("404 ERROR: User Agent: %s, IP: %s", r.UserAgent(), ip))
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is Not Supported", http.StatusMethodNotAllowed)
		log.Fatal(fmt.Sprintf("WRONG METHOD USED: User Agent: %s, IP: %s", r.UserAgent(), ip))
		return
	}
	log.Println(fmt.Sprintf("User Agent: %s, IP: %s", r.UserAgent(), ip))
	fmt.Fprintf(w, "hello!!!")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	log.Printf("Server started at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
