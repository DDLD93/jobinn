package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ddld93/jobinn/configs" //add this
	"github.com/ddld93/jobinn/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func init()  {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
  
}

func main() {
    port := os.Getenv("PORT")
    host := os.Getenv("HOST")
    router := gin.New()
	
    
    //run database
    configs.ConnectDB()
    configs.QestdbConn()
	routes.UserRoute(router) //add this


    router.Run(fmt.Sprintf("%s:%v",host,port))
    fmt.Printf("Server running on port %s",port)
}