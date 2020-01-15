package main

import (
	"log"
	// "time"
    "encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type DeviceInfo struct {
	ID string
	IPv4 string
}


func main() {
	theBigTable := make(map[string]DeviceInfo);

	route := gin.Default()


	route.POST("/post-info", func(c *gin.Context) {
		// curl -H "Content-Type:application/json" -X POST --data '{"ID": "0000", "IPv4": "172.23.32.1"}' http://127.0.0.1:9090/post-info
		data, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("c.Request.body: %v", string(data))
		var deviceInfo DeviceInfo 
		json.Unmarshal(data, &deviceInfo)
		log.Println(deviceInfo);

		theBigTable[deviceInfo.ID] = deviceInfo;
		log.Println(theBigTable);
	})


	route.GET("/get-all-list", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"data": theBigTable,
		})
	})


	route.Run(":9090")
}
