package app

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"awesomeProject/internal/models"
	"awesomeProject/internal/services"
)

type httpReq struct {
	Method string
	URL    string
}

func ServerStart(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	// close listener
	log.Printf("Listening on %s\n", addr)

	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}
		go routeRequest(conn)
	}
}

func routeRequest(conn net.Conn) {
	// incoming request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	r := getHttpRequest(buffer)
	paths := strings.Split(r.URL, "/")

	fmt.Printf("%+v", paths)
	log.Printf("method=%s url=%s\n", r.Method, r.URL)

	if !strings.EqualFold(paths[0], "account") {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		log.Printf("unknoun url: %s", r.URL)
		return
	}

	req, err := getApiCallData(r)
	go handleReq(req, conn)
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", string(buffer[:]), time)
	conn.Write([]byte(responseStr))

	// close conn
	conn.Close()
}

func getHttpRequest(buf []byte) httpReq {
	var req string
	for _, b := range buf {
		req = req + string(b)
		if b == 13 {
			break
		}
	}
	p := strings.Split(req, " ")
	method := p[0]
	url := strings.Trim(p[1], "/")
	return httpReq{
		Method: method,
		URL:    url,
	}
}

func getApiCallData(r httpReq) (apiCallData, error) {
	paths := strings.Split(r.URL, "/")
	if len(paths) == 1 {
		return apiCallData{operation: models.Create}, nil
	}

	var op models.OpType
	accountID, err := strconv.ParseInt(paths[1], 10, 64)
	if err != nil {
		return apiCallData{}, fmt.Errorf("cant parse ID %w", err)
	}

	switch paths[2] {
	case "deposite":
		op = models.Deposite
	case "withdraw":
		op = models.Withdraw
	case "balance":
		op = models.GetBalance
	}
	return apiCallData{accountID: accountID, operation: op}, err
}

func handleReq(req apiCallData, conn net.Conn) {
	switch req.operation {
	case models.Deposite:
		services.Deposite(req)
	}
}
