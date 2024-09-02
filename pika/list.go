package pika

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/router"
)

/*
curl 'https://pika.art/my-library' --compressed -X POST
*/
type PikaInfo struct {
	Status      string
	Id          string
	Video       string
	VideoPoster string
	PromptText  string
	Duration    int64
}

func ListAllAndUpdate(c *router.Context) {
	lastId := ""
	count := 1
	for {
		items, ok := List(lastId)
		if ok == false {
			continue
		}
		time.Sleep(time.Second * 1)
		fmt.Println(count)
		count++
		for _, item := range items {
			c.Params = map[string]any{}
			c.Params["duration"] = item.Duration
			c.Params["video_url"] = item.Video
			c.Params["video_poster"] = item.VideoPoster
			c.Update("link_section", "where id_pika=", item.Id)
			c.Update("pika", "where id_pika=", item.Id)
			c.Update("pika_render", "where id_pika=", item.Id)
			c.Params["id_pika"] = item.Id
			c.Insert("pika_inventory")
			lastId = item.Id
		}
		if len(items) == 0 {
			return
		}
	}
}

func List(after string) ([]PikaInfo, bool) {
	items := []PikaInfo{}
	url := fmt.Sprintf("https://pika.art/my-library")
	m := map[string]any{}
	m["after"] = after
	m["perPage"] = 12
	listOf := []any{m}
	b, _ := json.Marshal(listOf)

	//[{"after":"abcb8da5-9264-49a8-89dc-6520886198aa","perPage":12}]
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return items, false
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
		return items, false
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return items, false
	}

	lines := strings.Split(string(body), "\n")
	if len(lines) < 2 {
		return items, false
	}
	js := lines[1][2:]

	json.Unmarshal([]byte(js), &m)
	data, _ := m["data"].(map[string]any)
	results, _ := data["results"].([]any)
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
			duration, _ := video["duration"].(float64)

			//fmt.Println(id)
			//fmt.Println(promptText)
			//fmt.Println(status)
			//fmt.Println(resultUrl)
			//fmt.Println(videoPoster)
			if strings.Contains(resultUrl, "(") == false {
				//continue
			}
			if strings.Contains(resultUrl, ")") == false {
				//continue
			}
			if strings.Contains(resultUrl, "_sfx") == false {
				//	continue
			}
			pi := PikaInfo{}
			pi.Video = resultUrl
			pi.PromptText = promptText
			pi.Status = status
			pi.Id = id
			pi.VideoPoster = videoPoster
			pi.Duration = int64(duration)
			items = append(items, pi)
		}
	}

	return items, true
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
