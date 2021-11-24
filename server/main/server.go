package main


import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	. "github.com/WAY29/icecream-go/icecream"
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

func main() {

	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	iterate(currentDirectory+"/../../site/music")
    fs := http.FileServer(http.Dir("../../site"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}