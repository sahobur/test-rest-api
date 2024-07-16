package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"awesomeProject/internal/models"
	"awesomeProject/internal/services"
)

type httpReq struct {
	Method        string
	URL           string
	Host          string
	ContentType   string
	ContentLenght string
	Body          string
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
	fmt.Printf("req: %+v\n", r)
	//
	//fmt.Println("--------------------------------------------------------")
	//fmt.Printf("method=%s url=%s\n", r.Method, r.URL)
	//fmt.Printf("host=%s\n", r.Host)
	//fmt.Printf("Content type=%s\n", r.ContentType)
	//fmt.Printf("Content len=%s\n", r.ContentLenght)
	//fmt.Printf("Data=%s\n", r.Body)
	//fmt.Println("--------------------------------------------------------")

	if !strings.EqualFold(paths[0], "account") {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		log.Printf("unknoun url: %s", r.URL)
		return
	}

	var a models.Amount
	if len(r.Body) != 0 {
		err = json.Unmarshal([]byte(r.Body), &a)
		if err != nil {
			log.Printf("error unmarshal %s", err)
		}
	}
	fmt.Print(a)
	if r.Method == http.MethodPost {
		if paths[2] == "deposite" {
			id, err := strconv.ParseInt(paths[1], 10, 64)
			if err != nil {
				log.Printf("error id format %s", err)
				return
			}
			srv := services.NewAccountService(id)
			srv.Deposite(a.Amount)
		}
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
	var s string
	var data [16]string
	var body string
	i := 0
	startBodyIdx := 0
	var cl string
	for j, b := range buf {
		s += string(b)
		if b == 13 {
			if strings.Contains(s, "-Length") {
				sp := strings.Split(s, ":")
				cl = strings.Trim(sp[1], " \n\r")
				startBodyIdx = j + 4
				break
			}
			data[i] = s
		}
	}

	p := strings.Split(data[0], " ")
	method := p[0]
	url := strings.Trim(p[1], "/")
	//fmt.Printf("start body %d", startBodyIdx)
	for i = startBodyIdx; i < 1023; i++ {
		body += string(buf[i])
		if buf[i+1] == 0 {
			break
		}
	}
	return httpReq{
		Method:        method,
		URL:           url,
		ContentLenght: cl,
		Body:          body,
	}
}

func getApiCallData(r httpReq) (models.ApiCallData, error) {
	paths := strings.Split(r.URL, "/")
	if len(paths) == 1 {
		return models.ApiCallData{Operation: models.Create}, nil
	}

	var op models.OpType
	accountID, err := strconv.ParseInt(paths[1], 10, 64)
	if err != nil {
		return models.ApiCallData{}, fmt.Errorf("cant parse ID %w", err)
	}

	switch paths[2] {
	case "deposite":
		op = models.Deposite
	case "withdraw":
		op = models.Withdraw
	case "balance":
		op = models.GetBalance
	}
	return models.ApiCallData{AccountID: accountID, Operation: op}, err
}

func handleReq(req models.ApiCallData, conn net.Conn) {
	switch req.Operation {
	case models.Deposite:
		srv := services.NewAccountService(req.AccountID)
		srv.Deposite(10)
	}
}
