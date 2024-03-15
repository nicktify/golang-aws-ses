package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/joho/godotenv"
)

func main() {
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "development"
	}

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	client := ses.NewFromConfig(cfg)

	sourceEmail := os.Getenv("SOURCE_EMAIL")

	fmt.Println("Sending email from", sourceEmail)

	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{"destination@example.com"},
		},
		Source: aws.String(sourceEmail), // Replace with your "From" address
		Message: &types.Message{
			Body: &types.Body{
				Text: &types.Content{
					Data: aws.String("Hello, World!"),
				},
			},
			Subject: &types.Content{
				Data: aws.String("Hello, World!"),
			},
		},
	}

	_, err = client.SendEmail(context.TODO(), input)
	if err != nil {
		panic(err)
	}
}
