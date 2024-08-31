package pika

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func CreateFolder(id string) string {
	url := fmt.Sprintf("https://pika.art/my-library")
	listOf := []string{fmt.Sprintf(`{"folderName":"%s"}`, id)}
	b, _ := json.Marshal(listOf)

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:129.0) Gecko/20100101 Firefox/129.0")
	req.Header.Add("Referer", "https://pika.art/my-library")
	req.Header.Add("Origin", "https://pika.art/")
	req.Header.Add("Cookie", os.Getenv("PIKA_COOKIE"))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "text/plain;charset=UTF-8")
	req.Header.Add("Accept", "text/x-component")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Next-Action", "adf46c534ded157ba58132bafee719d875e9986d")
	req.Header.Set("Next-Router-State-Tree", "%5B%22%22%2C%7B%22children%22%3A%5B%22(dashboard)%22%2C%7B%22children%22%3A%5B%22my-library%22%2C%7B%22children%22%3A%5B%22__PAGE__%22%2C%7B%7D%5D%7D%5D%7D%5D%7D%2Cnull%2Cnull%2Ctrue%5D")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Priority", "u=4")
	req.Header.Set("TE", "trailers")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Fatalf("Failed to create gzip reader: %v", err)
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}
	fmt.Println(string(body))
	return ""

}

func MoveVideoToFolder(id, folder string) string {
	url := fmt.Sprintf("https://pika.art/my-library")
	listOf := []string{fmt.Sprintf(`{"videoId":"%s","folderName":"%s"}`, id, folder)}
	b, _ := json.Marshal(listOf)

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:129.0) Gecko/20100101 Firefox/129.0")
	req.Header.Add("Referer", "https://pika.art/my-library")
	req.Header.Add("Origin", "https://pika.art/")
	req.Header.Add("Cookie", os.Getenv("PIKA_COOKIE"))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "text/plain;charset=UTF-8")
	req.Header.Add("Accept", "text/x-component")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Next-Action", "9fd10e6a93a9327f131b754414158ac9af54eb92")
	req.Header.Set("Next-Router-State-Tree", "%5B%22%22%2C%7B%22children%22%3A%5B%22(dashboard)%22%2C%7B%22children%22%3A%5B%22my-library%22%2C%7B%22children%22%3A%5B%22__PAGE__%22%2C%7B%7D%5D%7D%5D%7D%5D%7D%2Cnull%2Cnull%2Ctrue%5D")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Priority", "u=4")
	req.Header.Set("TE", "trailers")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Fatalf("Failed to create gzip reader: %v", err)
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}
	fmt.Println(string(body))
	return ""

}
