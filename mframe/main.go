package main

import (
	"fmt"
	"net/http"
	"swwgo/mframe/package/logging"
	"swwgo/mframe/package/orm"
	_ "swwgo/mframe/package/setting"
	"swwgo/mframe/package/setting/conf"
	"swwgo/mframe/router"
)

func main() {
	r := router.InitRouter()

	logging.Init()
	orm.Init()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", conf.ServiceSetting.HttpPort),
		Handler: r,
	}

	logging.Debug(s.Addr)

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
