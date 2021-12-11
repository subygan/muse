package main

import (
	"context"
	"encoding/json"
	"fmt"
	. "github.com/WAY29/icecream-go/icecream"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
)


func callback (path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Fatalf(err.Error())
	}
	Ic(info)
	fmt.Printf("path: %s", info.IsDir())
	Ic(info.Sys())
	fmt.Printf("File Name: %s\n", info.Name())
	return nil
}

func iterate(path string) {
	err := filepath.Walk(path, callback)
	if err != nil {
		return 
	}
}

func getAllKeys(client *s3.Client) ([]*string, int32) {
	bucketName:= "suriya-music"
	params := &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	}
	// Create the Paginator for the ListObjectsV2 operation.
	p := s3.NewListObjectsV2Paginator(client, params, func(o *s3.ListObjectsV2PaginatorOptions) {
		if v := 100; v != 0 {
			o.Limit = int32(v)
		}
	})

	// Iterate through the S3 object pages, printing each object returned.
	var i int
	var count int32

	var l []*string
	log.Println("Objects:")
	for p.HasMorePages() {
		i++

		// Next Page takes a new context for each page retrieval. This is where
		// you could add timeouts or deadlines.
		page, err := p.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("failed to get page %v, %v", i, err)
		}

		count += page.KeyCount
		// Log the objects found
		for _, obj := range page.Contents {
			fmt.Println("Object:", *obj.Key)
			l = append(l, obj.Key)
		}
	}
	fmt.Println("#### ", count)

	return l, count
}



func returnAllArticles(w http.ResponseWriter, r *http.Request){


	//output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
	//	Bucket: aws.String("suriya-music"),
	//})
	//if err != nil {
	//	Ic(err)
	//}

	key1 := obj.keys[rand.Intn(int(obj.count))]
	h := s3.GetObjectInput{
		Bucket: aws.String("suriya-music"),
		Key: key1,
	}
	psClient := s3.NewPresignClient(obj.Client)

	Res, _ := psClient.PresignGetObject(context.TODO(), &h)

	Ic(Res)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Url{
		Url: "fuck you",
	})
}


type Url struct {
	Url string `json:"Url"`
}


type Obj struct {
	Client *s3.Client
	keys []*string
	count int32
}

var obj Obj

func main() {
	Ic("Working??")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		Ic("Fatal error")
	}

	client := s3.NewFromConfig(cfg)
	obj.Client = client

	l, count := getAllKeys(client)
	obj.keys = l
	obj.count = count

	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", returnAllArticles)

	log.Fatal(http.ListenAndServe(":3000", myRouter))
}