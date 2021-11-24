package main

import (
	"context"
	"encoding/json"
	"fmt"
	. "github.com/WAY29/icecream-go/icecream"
	"github.com/aws/aws-sdk-go-v2/aws"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gorilla/mux"
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

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	Ic("Working??")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		Ic("Fatal error")
	}

	client := s3.NewFromConfig(cfg)

	//output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
	//	Bucket: aws.String("suriya-music"),
	//})
	//if err != nil {
	//	Ic(err)
	//}

	fmt.Println("####")

	key1 := "04_Uproar.mp3"
	i := s3.GetObjectInput{
		Bucket: aws.String("suriya-music"),
		Key: &key1,
	}
	psClient := s3.NewPresignClient(client)

	Res, _ := psClient.PresignGetObject(context.TODO(), &i)

	Ic(Res)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Url{
		Url: Res.URL,
	})
}


type Url struct {
	Url string `json:"Url"`
}

func main() {



	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", returnAllArticles)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":3000", myRouter))



	//currentDirectory, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//iterate(currentDirectory+"/../../site/music")
    //fs := http.FileServer(http.Dir("../../site"))
    //
	//http.Handle("/", fs)
	//
	//log.Println("Listening on :3000...")
	//err = http.ListenAndServe(":3000", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
}