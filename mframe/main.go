package main

import (
	"fmt"
	"net/http"
	"swwgo/mframe/package/setting/conf"
	"swwgo/mframe/router"
)

func main() {
	r := router.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", conf.ServiceSetting.HttpPort),
		Handler: r,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
