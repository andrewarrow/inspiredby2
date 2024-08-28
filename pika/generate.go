package pika

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Generate(text string) string {
	url := fmt.Sprintf("https://api.pika.art/generate")

	m := map[string]any{}
	b, _ := json.Marshal(m)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:129.0) Gecko/20100101 Firefox/129.0")
	req.Header.Add("Referer", "https://pika.art")
	req.Header.Add("Origin", "https://pika.art")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("PIKA"))
	req.Header.Add("Content-Type", "multipart/form-data; boundary=---------------------------266926460920144731353527800262")

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

	fmt.Println(string(body))
	return ""

}
