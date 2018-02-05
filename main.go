package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func env(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}

func main() {
	port := env("PORT", "5000")
	workDir, _ := os.Getwd()
	dir := env("DIR", workDir)

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	srvr := http.Server{
		Addr: ":" + port,
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
