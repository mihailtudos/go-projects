package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyResponse struct {
	Message string `json:"Answer"`
}

type MyEvent struct {
	Names string `json:"what is your name?"`
	Age   int    `json:"How old are you?"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", event.Names, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
