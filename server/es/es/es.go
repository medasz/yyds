package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

var ESClient *elastic.Client

func SaveInfo(index string, data []map[string]string) {
	for _, val := range data {
		_, err := ESClient.Index().Index(index).BodyJson(val).Do(context.Background())
		if err != nil {
			log.Println(err)
		}
	}
}

func ConnectES() {
	var err error
	ESClient, err = elastic.NewClient(elastic.SetBasicAuth("elastic","niYYmab445%^"))
	if err != nil {
		panic(err)
	}
}
