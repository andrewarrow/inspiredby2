package google

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/vertexai/genai"
)

func Summarize() {
	ctx := context.Background()
	projectID := "local-dev-353516"
	location := "us-central1"
	modelName := "gemini-1.5-pro-001"

	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)

	textToSummarize := ` person helping them be aware what it feels like for them at the lower HRV and
what it feels like at higher HIV and then understanding the factors that they can control um to continue to amplify their HRV
talk to me about the relationship between HRV autonomic flexibility and the brain and brain function  so we used to think of
what it feels like at higher HIV and then understanding the factors that they can control um to continue to amplify their HRV
talk to me about the relationship between HRV autonomic flexibility and the brain and brain function  so we used to think of
the heart rate variability just as a metric of the autonomic nervous systems resilience to flexibly respond which means if I need to amp up and run
across the street I can this isn't this isn't associated with relaxation or just being calm it's acclimating or adapting flexibly to the needs of the most
and resilience in the last 10 to 15 years uh myself and several other researchers have focused on heart rate variabilities impact.`
	prompt := fmt.Sprintf("Summarize the following text:\n\n%s", textToSummarize)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatalf("Error generating content: %v", err)
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		summary := resp.Candidates[0].Content.Parts[0]
		fmt.Printf("Summary: %s\n", summary)
	} else {
		fmt.Println("No summary generated.")
	}
}
