package main

import (
	"backend/config"
	"backend/handlers"
	"backend/middleware"
	"fmt"
	"net/http"
	"os"
)

func main() {

	//1. init config file
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprint(os.Stderr, "error making log file")
		return
	}
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	con := config.InitConfig(logFile)
	//2. init server
	mux := http.NewServeMux()
	//3. register handler
	handlers.Register(mux, con)
	//4. register middleware
	server := middleware.Register(mux)
	//5. run server
	http.ListenAndServe(listenAddr, server)
}
