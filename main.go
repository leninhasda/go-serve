package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	workDir, _ := os.Getwd()

	port := *flag.String("port", "5000", "Port for the application")
	dir := *flag.String("dir", workDir, "Directory to serve")
	flag.Parse()

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	srvr := http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		log.Printf("FileServer started at - http://localhost:%s\n", port)
		log.Fatal(srvr.ListenAndServe())
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srvr.Shutdown(ctx)
	log.Println("Server shutteddown gracefully")
}
