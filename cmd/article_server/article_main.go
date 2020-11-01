package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/vikasd145/article_project/internal/config"
	"github.com/vikasd145/article_project/pkg/article_admin"

	"github.com/vikasd145/article_project/pkg/Rotuers"
)

var (
	ormconfig     = flag.String("ormdb config", "../configs/ormdb.xml", "Path to orm config")
	articleconfig = flag.String("atricle config", "../configs/article.xml", "Path to article config")
)

func main() {
	flag.Parse()
	ormcon := &config.OrmContainer{}
	err := ormcon.LoadFromFile(*ormconfig)
	if err != nil {
		fmt.Errorf("Error in loading orm config file err:%v", err)
		return
	}
	conf := &config.Config{}
	err = conf.LoadFromFile(*articleconfig)
	if err != nil {
		fmt.Errorf("Error in loading article conf err:%v", err)
		return
	}
	conf.Ormconfigs = ormcon.Ormconfigs
	_, err = article_admin.AdminInitialize(conf)
	if err != nil {
		fmt.Errorf("Error in intialializing admin err:%v", err)
		return
	}
	log.Fatalf("Server Crashed", http.ListenAndServe(":8080", Rotuers.NewRouter("")))
}
