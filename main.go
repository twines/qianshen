/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/8/28 9:09
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"qianshen/models"
	"qianshen/pkg/setting"
	"qianshen/routers"
)

func init() {
	setting.Setup()
	models.Setup()
}
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	//s.ListenAndServe()
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := s.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}

	log.Println("Server exiting")
}
