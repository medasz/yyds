package model

import "github.com/olivere/elastic/v7"

var ESClient *elastic.Client

func Init() {
	var err error
	ESClient, err = elastic.NewClient()
	if err != nil {
		panic(err)
	}
}
