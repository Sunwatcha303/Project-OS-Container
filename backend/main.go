package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Sunwatcha303/Project-OS-Container/config"
	"github.com/Sunwatcha303/Project-OS-Container/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // on production use gin.ReleaseMode
	config.InitDatabaseConnection()
	r := handler.Routes{}
	handlerRoute := r.InitRouter()
	AppSrv := &http.Server{Addr: "0.0.0.0:8080", Handler: handlerRoute}
	fmt.Println("[Main] Listening at port :8080")
	go func() {
		var err error = nil
		err = AppSrv.ListenAndServe()
		if err != nil {
			log.Fatalf("[Main] Unable to start server: %+v", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
}
