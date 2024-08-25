package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Summarize(text string) {
	/*
		curl -X POST "https://api.groq.com/openai/v1/chat/completions" \
		     -H "Authorization: Bearer $GROQ_API_KEY" \
		     -H "Content-Type: application/json" \
		     -d '{"messages": [{"role": "user", "content": "Explain the importance of fast language models"}], "model": "llama3-8b-8192"}'
	*/
	url := fmt.Sprintf("https://api.groq.com/openai/v1/chat/completions")

	message := map[string]any{}
	message["role"] = "user"
	message["content"] = "Summarize in 30 words the text: " + text
	messages := []any{message}
	m := map[string]any{}
	m["messages"] = messages
	m["model"] = "llama3-8b-8192"
	b, _ := json.Marshal(m)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	fmt.Println(string(body))
	json.Unmarshal(body, &m)
}
