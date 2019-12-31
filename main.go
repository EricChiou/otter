package main

import (
	"log"

	"otter/acl"
	"otter/config"
	cons "otter/constants"
	"otter/db/mysql"
	"otter/router"
	"otter/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load config
	if err := config.Load(cons.ConfigFilePath); err != nil {
		panic(err)
	}

	// init db
	if err := mysql.Init(); err != nil {
		panic(err)
	}
	defer mysql.Close()

	// load acl
	if err := acl.Load(); err != nil {
		panic(err)
	}

	// set headers
	router.SetHeader("Access-Control-Allow-Origin", "*")
	router.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	router.SetHeader("Access-Control-Allow-Headers", "Content-Type")

	// init api
	routes.Init()

	// start http server
	if err := router.ListenAndServe(config.Conf.ServerPort); err != nil {
		panic(err)
	}
	// start https server
	// if err = router.ListenAndServeTLS(config.Conf.ServerPort, config.Conf.SSLCertFilePath, config.Conf.SSLKeyFilePath); err != nil {
	// 	panic(err)
	// }

	defer func() {
		if err := recover(); err != nil {
			log.Println("start server error:", err)
		}
	}()
}
