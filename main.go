package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type dan struct {
	Dan  Dan    `json:"dan"`
	Stat string `json:"stat"`
}

type Dan struct {
	Says string `json:"says"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	rand.Seed(time.Now().UnixNano())

	sayings := []string{
		"You're doing a great job, Dan.",
		"Keep up the good work, Dan",
		"Don't stop now, Dan",
		"You're doing just fine, Dan",
	}

	randomSaying := rand.Intn(len(sayings))

	says := Dan{
		Says: sayings[randomSaying],
	}

	body := dan{
		Dan:  says,
		Stat: "ok",
	}

	b, _ := json.Marshal(body)

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
