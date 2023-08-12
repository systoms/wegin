package main

import (
	_ "embed"
	"fmt"
	"github.com/systoms/wegin/backend"
	"github.com/systoms/wegin/config"
	"github.com/systoms/wegin/database"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	g errgroup.Group
)

//go:embed config.yaml
var configFileData []byte

func main() {

	//初始化配置
	conf, err := config.LoadConfig(configFileData)
	if err != nil {
		fmt.Println("配置文件加载失败:", err)
		return
	}

	//初始化数据库
	db, err := database.InitDatabase(conf.Database)
	if err != nil {
		fmt.Println("无法初始化数据库:", err)
		return
	}

	//初始化后台
	g.Go(func() error {
		return backend.Run(conf, db)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
