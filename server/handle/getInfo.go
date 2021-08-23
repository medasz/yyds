package handle

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hids/server/es/es"
	"log"
)

func GetInfo(c *gin.Context) {
	searchResult, err := es.ESClient.Search("loginlog").Sort("time", false).Size(1).Do(context.Background())
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"message": "success", "time": "all"})
		return
	}
	if searchResult.Hits.TotalHits.Value != 0 {
		dataJson, err := searchResult.Hits.Hits[0].Source.MarshalJSON()
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"message": "success", "time": "all"})
			return
		}
		tmp := make(map[string]interface{})
		err = json.Unmarshal(dataJson, &tmp)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{"message": "success", "time": "all"})
			return
		}
		if err!=nil{
			log.Println(err)
			c.JSON(200, gin.H{"message": "success", "time": "all"})
			return
		}
		c.JSON(200, gin.H{"message": "success", "time": tmp["time"].(string)})
		return
	}
	c.JSON(200, gin.H{"message": "success", "time": "all"})
}
