package main

import (
	"github.com/systoms/wegin/backend/routers"
)

func main() {
	backend := routers.SetupRouter()
	go func() {
		err := backend.Run()
		if err != nil {
			return
		} // 监听并在 0.0.0.0:8080 上启动服务
	}()
	select {}
}
