package google

import (
	"context"
	"fmt"
	"io/ioutil"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func Speech() {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	filePath := "path_to_your_audio_file.wav"
	audioData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:                 speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz:          16000,
			LanguageCode:             "en-US",
			EnableSpeakerDiarization: true,
			DiarizationSpeakerCount:  2,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: audioData},
		},
	}

	resp, err := client.Recognize(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("Transcription: %v\n", alt.Transcript)
			fmt.Printf("Confidence: %v\n", alt.Confidence)
		}
	}
}
