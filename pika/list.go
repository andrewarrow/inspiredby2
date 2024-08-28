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
	"strings"
)

/*
curl 'https://pika.art/my-library' --compressed -X POST
*/
func List() string {
	url := fmt.Sprintf("https://pika.art/my-library")
	m := map[string]any{}
	m["after"] = "abcb8da5-9264-49a8-89dc-6520886198aa"
	m["perPage"] = 12
	listOf := []any{m}
	b, _ := json.Marshal(listOf)

	//[{"after":"abcb8da5-9264-49a8-89dc-6520886198aa","perPage":12}]
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
	req.Header.Set("Next-Action", "6c92104373247fc5471997bb7e08fab5c44eec8a")
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

	lines := strings.Split(string(body), "\n")
	if len(lines) < 2 {
		return ""
	}
	js := lines[1][2:]

	json.Unmarshal([]byte(js), &m)
	data, _ := m["data"].(map[string]any)
	results := data["results"].([]any)
	for _, item := range results {
		thing, _ := item.(map[string]any)
		id, _ := thing["id"].(string)
		promptText, _ := thing["promptText"].(string)
		videos := thing["videos"].([]any)
		for _, videoThing := range videos {
			video, _ := videoThing.(map[string]any)
			status, _ := video["status"].(string)
			resultUrl, _ := video["resultUrl"].(string)
			videoPoster, _ := video["videoPoster"].(string)

			fmt.Println(id)
			fmt.Println(promptText)
			fmt.Println(status)
			fmt.Println(resultUrl)
			fmt.Println(videoPoster)
		}
	}

	return ""

}
