package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rmkane/go-wiki-server/internal/handler"
	"github.com/rmkane/go-wiki-server/internal/middleware"
)

const port = 8080

func main() {
	// Serve static files from /static/
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/view/", middleware.MakeHandler(handler.ViewHandler))
	http.HandleFunc("/edit/", middleware.MakeHandler(handler.EditHandler))
	http.HandleFunc("/save/", middleware.MakeHandler(handler.SaveHandler))

	server := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	go func() {
		fmt.Printf("Server started on port %d\n", port)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\nShutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	fmt.Println("Server exited properly")
}
