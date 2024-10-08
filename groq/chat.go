package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Summarize(text string) string {
	/*
		curl -X POST "https://api.groq.com/openai/v1/chat/completions" \
		     -H "Authorization: Bearer $GROQ_API_KEY" \
		     -H "Content-Type: application/json" \
		     -d '{"messages": [{"role": "user", "content": "Explain the importance of fast language models"}], "model": "llama3-8b-8192"}'
	*/
	url := fmt.Sprintf("https://api.groq.com/openai/v1/chat/completions")

	message := map[string]any{}
	message["role"] = "user"
	message["content"] = "Summarize in 30 words the following text but do not include anything in your reply other than the summary itself. Do not say 'here is your 30 word summary' just reply with the summary. Here is the text: " + text
	messages := []any{message}
	m := map[string]any{}
	m["messages"] = messages
	m["model"] = "llama3-8b-8192"
	b, _ := json.Marshal(m)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("GROQ_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

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

	//	fmt.Println(string(body))

	json.Unmarshal(body, &m)
	choices := m["choices"].([]any)
	if len(choices) == 0 {
		return ""
	}
	thing, _ := choices[0].(map[string]any)
	if len(thing) == 0 {
		return ""
	}
	choiceMessage := thing["message"].(map[string]any)
	s, _ := choiceMessage["content"].(string)
	return s

	/*

	  "choices": [
	    {
	      "index": 0,
	      "message": {
	        "role": "assistant",
	        "content": "Here is a 30-word summary of the text:\n\nUnderstanding heart rate variability (HRV) helps individuals recognize its impact on their emotions, cognitions, and physical responses, and recognizing factors to influence HRV to enhance resilience."
	      },

	*/
}
