package google

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func Speech(filePath string) string {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	audioData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:          speechpb.RecognitionConfig_FLAC,
			SampleRateHertz:   16000,
			LanguageCode:      "en-US",
			AudioChannelCount: 2,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: audioData},
		},
	}

	resp, err := client.Recognize(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	buffer := []string{}
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			buffer = append(buffer, fmt.Sprintf("%v", alt.Transcript))
		}
	}

	return strings.Join(buffer, " ")
}
