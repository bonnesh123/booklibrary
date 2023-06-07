package main

import (
	"log"
	"net/http"
	"training/app/logging"
	"training/app/util"

	"github.com/gin-gonic/gin"
)

var config *util.Config

func init() {
	log.SetFlags(1)
	logging.SetUpLogging()
	myConfig, err := util.LoadConfig(".")
	if err != nil {
		logging.AppLog.WriteLogsWError("cannot load config: ",
			map[string]interface{}{"source": config, "err": err})
	}
	config = &myConfig
	util.LoadCollection("./static/")
	util.InitialiazeRedis()
}

func main() {
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)

	logging.AppLog.WriteLogsInfo("App running in env:", map[string]interface{}{
		"runtime_setup": config.RuntimeSetup, "app_port": config.AppPort})
	//log.Println("App running in env:", config.RuntimeSetup, "and on port:", config.AppPort)
	//log.Println("App running in env:", viper.GetString("RUNTIME_SETUP"), "and on port:", viper.GetInt("APP_PORT "))

	//var router *gin.Engine
	router := gin.Default()

	router.Static("/static/", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/collection", getCollection)

	//router.Run(viper.GetString("SERVER_ADDRESS") + ":" + viper.GetString("APP_PORT"))
	router.Run(config.ServerAddress + ":" + config.AppPort)
	// log.Println("Starting the app, listening on :8080...")
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func getCollection(c *gin.Context) {
	val := util.GetBookList()
	c.HTML(http.StatusOK, "library.html", gin.H{
		"books": val.BookList,
	})
}
