package pika

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Generate(video, text string) string {
	url := fmt.Sprintf("https://api.pika.art/generate")

	data := bytes.NewBufferString("")
	data.WriteString("-----------------------------266926460920144731353527800262\r\n")
	data.WriteString("Content-Disposition: form-data; name=\"styleId\"\r\n\r\n\r\n")
	data.WriteString("-----------------------------266926460920144731353527800262\r\n")
	data.WriteString("Content-Disposition: form-data; name=\"promptText\"\r\n\r\n")
	data.WriteString(text + "\r\n")
	if video == "" {
		data.WriteString("-----------------------------266926460920144731353527800262\r\n")
		data.WriteString("Content-Disposition: form-data; name=\"sfx\"\r\n\r\ntrue\r\n")
	} else {
		data.WriteString("-----------------------------266926460920144731353527800262\r\n")
		data.WriteString(fmt.Sprintf("Content-Disposition: form-data; name=\"video\"\r\n\r\n%s\r\n", video))
	}
	data.WriteString("-----------------------------266926460920144731353527800262\r\n")
	if video == "" {
		data.WriteString("Content-Disposition: form-data; name=\"options\"\r\n\r\n{\"frameRate\":24,\"parameters\":{\"guidanceScale\":25,\"motion\":4,\"negativePrompt\": \"human eyes writing words\"},\"camera\":{\"zoom\":null,\"pan\":null,\"tilt\":null,\"rotate\":null},\"extend\":false}\r\n")
	} else {
		data.WriteString("Content-Disposition: form-data; name=\"options\"\r\n\r\n{\"frameRate\":24,\"parameters\":{\"guidanceScale\":25,\"motion\":4,\"negativePrompt\": \"human eyes writing words\"},\"camera\":{\"zoom\":null,\"pan\":null,\"tilt\":null,\"rotate\":null},\"extend\":true}\r\n")
	}
	data.WriteString("-----------------------------266926460920144731353527800262\r\n")
	data.WriteString(fmt.Sprintf("Content-Disposition: form-data; name=\"userId\"\r\n\r\n%s\r\n", os.Getenv("PIKA_USER")))
	data.WriteString("-----------------------------266926460920144731353527800262--\r\n")

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, data)
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

	fmt.Println(len(body))
	//time.Sleep(time.Second * 90)
	var m map[string]any
	json.Unmarshal(body, &m)
	d1, _ := m["data"].(map[string]any)
	id, _ := d1["id"].(string)
	return id

}
