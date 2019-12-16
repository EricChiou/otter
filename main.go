package main

import (
	"log"

	"otter/routes"
	"otter/config"
	"otter/db/mysql"
	"otter/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := config.LoadConfig("./config.ini")
	if err != nil {
		panic(err)
	}

	err = mysql.Init(
		config.Config.MySQLAddr,
		config.Config.MySQLPort,
		config.Config.MySQLUserName,
		config.Config.MySQLPassword,
		config.Config.MySQLDBNAME,
	)
	if err != nil {
		panic(err)
	}
	defer mysql.Close()

	// set headers
	router.SetHeader("Access-Control-Allow-Origin", "*")
	router.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	router.SetHeader("Access-Control-Allow-Headers", "Content-Type")

	// init api
	routes.InitUserAPI()

	port := ":" + config.Config.ServerPort
	// start http server
	if err = router.ListenAndServe(port); err != nil {
		panic(err)
	}
	// start https server
	// if err = router.ListenAndServeTLS(port, config.Config.SSLCertFilePath, config.Config.SSLKeyFilePath); err != nil {
	// 	panic(err)
	// }

	defer func() {
		if err := recover(); err != nil {
			log.Println("start server error:", err)
		}
	}()
}
