package main


import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	. "github.com/WAY29/icecream-go/icecream"
)


func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}
func callback (path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Fatalf(err.Error())
	}
	Ic(info)
	fmt.Printf("path: %s", info.IsDir())
	fmt.Printf("path: %s", info.Mode())
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
	iterate(currentDirectory+"/../site/music")


    fs := http.FileServer(http.Dir("../site"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}



}