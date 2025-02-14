package main

import (
	"flag"
	"fmt"
	question1 "seven-hunter-assignment/question-1"
	question2 "seven-hunter-assignment/question-2"
	"seven-hunter-assignment/question-3/service/grpc"
	"seven-hunter-assignment/question-3/service/http"
	"sync"
)

func main() {
	fmt.Println("Question 1: จงหาเส้นทางที่มีค่ามากที่สุด")
	question1.AnswerQuestion1()
	fmt.Println("==============================")
	fmt.Println("")
	fmt.Println("")

	//inputMode := "KEYBOARD"
	inputMode := "TEST_CASE"
	fmt.Println("Question 2: จับฉันให้ได้สิ ซ้าย-ขวา-เท่ากับ")
	fmt.Println("==============================")
	if inputMode == "TEST_CASE" {
		question2.AnswerQuestion2("LLRR=")
		question2.AnswerQuestion2("==RLL")
		question2.AnswerQuestion2("=LLRR")
		question2.AnswerQuestion2("RRL=R")
	} else {
		var encoded string
		fmt.Print("Input encoded string: ")
		_, err := fmt.Scanln(&encoded)
		if err != nil {
			return
		}
		question2.AnswerQuestion2(encoded)
	}
	fmt.Println("")
	fmt.Println("")

	skipServer := flag.Bool("skip-server", false, "Run the application with skip server")
	flag.Parse()
	if *skipServer {
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := grpc.StartGrpcServer(); err != nil {
			fmt.Println("gRPC server error:", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := http.StartHttpServer(); err != nil {
			fmt.Println("HTTP server error:", err)
		}
	}()

	wg.Wait()
}
