package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Download(dir, id, url string) {
	tokens := strings.Split(url, "/")
	last := tokens[len(tokens)-1]
	tokens = strings.Split(last, ".")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	outFile, err := os.Create(dir + "/" + id + "." + tokens[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("File saved successfully!")
}
