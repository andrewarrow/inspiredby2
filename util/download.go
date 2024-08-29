package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Download(id, url string) {
	tokens := strings.Split(url, "/")
	last := tokens[len(tokens)-1]

	fmt.Println("c", id, url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	outFile, err := os.Create("data/" + id + "_" + last)
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
