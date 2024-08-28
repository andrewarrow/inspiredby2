package pika

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
curl 'https://pika.art/my-library' --compressed -X POST
*/
func List() string {
	url := fmt.Sprintf("https://api.pika.art/my-library")
	m := map[string]any{}
	b, _ := json.Marshal(m)

	//[{"after":"abcb8da5-9264-49a8-89dc-6520886198aa","perPage":12}]
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:129.0) Gecko/20100101 Firefox/129.0")
	req.Header.Add("Referer", "https://pika.art/my-library")
	req.Header.Add("Origin", "https://pika.art/my-library")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("PIKA"))
	req.Header.Add("Content-Type", "text/plain;charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	fmt.Println(len(body))
	return ""

}
