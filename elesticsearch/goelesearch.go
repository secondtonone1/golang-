package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"

	"github.com/olivere/elastic"
)

type LogData struct {
	Topic string
	Log   string
}

var host = "http://127.0.0.1:9200"
var esOnce sync.Once
var esClient *elastic.Client = nil
var logger *log.Logger = nil

func GetEsClient() *elastic.Client {
	esOnce.Do(func() {
		logger = log.New(os.Stdout, "LOGCAT", log.LstdFlags|log.Lshortfile)
		var err error
		esClient, err = elastic.NewClient(elastic.SetURL(host),
			elastic.SetErrorLog(logger))
		if err != nil {
			// Handle error
			logger.Println("create elestic client error ", err.Error())
			return
		}
		fmt.Println(esClient)
		fmt.Println(reflect.TypeOf(esClient))
		// Use the IndexExists service to check if a specified index exists.

		info, code, err := esClient.Ping(host).Do(context.Background())
		if err != nil {
			logger.Println("elestic search ping error, ", err.Error())
			esClient = nil
			return
		}
		fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

		esversion, err := esClient.ElasticsearchVersion(host)
		if err != nil {
			fmt.Println("elestic search version get failed, ", err.Error())
			esClient = nil
			return
		}
		fmt.Printf("Elasticsearch version %s\n", esversion)

	})

	return esClient
}

func Exists(logdata *LogData) bool {
	esClient := GetEsClient()
	if esClient == nil {
		logger.Println("get es client faild")
		return false
	}

	exists, err := esClient.IndexExists(logdata.Topic).Do(context.Background())
	if err != nil {
		// Handle error
		logger.Println("elestic search IndexExists error ", err.Error())
		return false
	}
	if !exists {
		// Create a new index.
		logger.Printf("elestic search index %s not exists ", logdata.Topic)
		return false
	}
	logger.Printf("elestic search index %s exists", logdata.Topic)
	return true
}

func Create(logdata *LogData, id string, typestr string) bool {

	createIndex, err := esClient.Index().Index(logdata.Topic).Type(typestr).Id(id).BodyJson(logdata).Do(context.Background())

	if err != nil {
		// Handle error
		logger.Println("create index failed, ", err.Error())
		return false
	}
	logger.Println("create success")
	logger.Println(createIndex)
	return true
}

func Get(logdata *LogData, id string, typestr string) bool {

	if !Exists(logdata) {
		logger.Println("index not exists, index: ", logdata.Topic)
		return false
	}
	esResponse, err := esClient.Get().Index(logdata.Topic).Type(typestr).Id(id).Do(context.Background())
	if err != nil {
		logger.Println("index get failed, error ", err.Error())
		return false
	}
	logger.Printf("get success, id is %s, type is %s, index is %s", esResponse.Id, esResponse.Type, esResponse.Index)
	//json.Unmarshal(*esResponse.Source, &doc)
	logdatatmp := &LogData{}
	json.Unmarshal(esResponse.Source, logdatatmp)
	logger.Println("logdata is ", logdatatmp)
	return true
}

func Update(logdata *LogData, id string, typestr string, updateField map[string]interface{}) bool {
	if !Exists(logdata) {
		logger.Println("index not exists, index: ", logdata.Topic)
		return false
	}
	esResponse, err := esClient.Update().Index(logdata.Topic).Type(typestr).Id(id).Doc(updateField).Do(context.Background())
	if err != nil {
		logger.Println("index update failed, error ", err.Error())
		return false
	}
	logger.Printf("update success, id is %s, type is %s, index is %s", esResponse.Id, esResponse.Type, esResponse.Index)
	return true
}

func Delete(logdata *LogData, id string, typestr string) bool {
	if !Exists(logdata) {
		logger.Println("index not exists, index: ", logdata.Topic)
		return true
	}

	esResponse, err := esClient.Delete().Index(logdata.Topic).Type(typestr).Id(id).Do(context.Background())
	if err != nil {
		logger.Println("index delete failed, error ", err.Error())
		return false
	}
	logger.Printf("delete success, id is %s, type is %s, index is %s", esResponse.Id, esResponse.Type, esResponse.Index)
	return true
}

/*
func Quary() {
	var res *elastic.SearchResult
	var err error
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("topic", "logdir1"))

	res, err = esClient.Search().Index("logdir1").Type("catlog").Query(query).Do(context.Background())
	if err != nil {
		logger.Println(err.Error())
		return
	}
	logger.Println(res)
}
*/

func main() {
	esClient := GetEsClient()
	if esClient == nil {
		logger.Println("get es client faild")
		return
	}
	defer esClient.Stop()
	logdata := &LogData{Topic: "logdir3", Log: "logdir3log"}
	Exists(logdata)

	Create(logdata, "2", "catlog")
	Get(logdata, "2", "catlog")
	changeMap := make(map[string]interface{})
	changeMap["Log"] = "Change to be helloworld"
	Update(logdata, "2", "catlog", changeMap)
	Get(logdata, "2", "catlog")

	logdata1 := &LogData{Topic: "logdir1", Log: "logdir1log"}
	Create(logdata1, "1", "catlog")
	/*
		Delete(logdata, "2", "catlog")
		Delete(logdata, "1", "catlog")
		Delete(&LogData{Topic: "logdir1", Log: "logdir3log"}, "1", "catlog")
	*/
}
