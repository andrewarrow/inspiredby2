package eleven

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiURL = "https://api.elevenlabs.io/v1/synthesize"
	apiKey = "YOUR_API_KEY"
)

type TTSRequest struct {
	Text   string `json:"text"`
	Voice  string `json:"voice,omitempty"`
	Format string `json:"format,omitempty"`
}

func TextToSpeech() {
	//model_id: eleven_turbo_v2_5
	// https://api.elevenlabs.io/v1/text-to-speech/XjLkpWUlnhS8i7gGz3lZ
	// https://api.elevenlabs.io/v1/text-to-speech/repzAAjoKlgcT2oOAIWt
	requestBody := TTSRequest{
		Text:   "Researchers found that heart rate variability (HRV) relates to cognitive dexterity, enabling flexible thinking and focus, while higher HRV is correlated with a lower risk of cardiovascular conditions and increased autonomic nervous system flexibility.",
		Voice:  "en_us_male",
		Format: "mp3",
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling request data:", err)
		return
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	err = ioutil.WriteFile("output.mp3", respBody, 0644)
	if err != nil {
		fmt.Println("Error saving audio file:", err)
		return
	}

	fmt.Println("TTS audio saved as output.mp3")
}
