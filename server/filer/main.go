package main

import (
	"context"
	"fmt"
	. "github.com/WAY29/icecream-go/icecream"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/dhowden/tag"
	//"log"
	//"os"
	//"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Format string
type FileType string
type Track struct {
	num int
	tot int
}

type Disc struct {
	num int
	tot int
}

type Music struct {
	format   Format
	fileType FileType

	title       string
	album       string
	artist      string
	albumArtist string
	composer string
	genre string
	year string

	track Track
	disc Disc

	picture tag.Picture
	comment string
	raw map[string]interface{}
}

type Dir struct {
	name string
	sub []*Dir
	file []Music
}

type Filer interface {
	FetchPath(path string)
}

func main() {
	Ic("Working??")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		Ic("Fatal error")
	}

	client := s3.NewFromConfig(cfg)

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("suriya-music"),
	})
	if err != nil {
		Ic(err)
	}
	fmt.Println("####")

	for _, object := range output.Contents {
		Ic(*object.Key)
		Ic(object)

		//i := s3.GetObjectInput{
		//	Bucket: aws.String("suriya-music"),
		//	Key: object.Key,
		//}
		//psClient := s3.NewPresignClient(client)
		//
		//res, _ := psClient.PresignGetObject(context.TODO(), &i)
		//
		//Ic(res)

	}






	//   /Users/suriyaganesh/Documents/VL/muse/site/music/horse.mp3
}
