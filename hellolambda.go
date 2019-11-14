package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

//Handler : Lambda handler invoked by lambda.Start
func Handler() (string, error) {
	j, err := json.Marshal(map[string]interface{}{"message": fmt.Sprintf(
		"Hello %v! Your function executed successfully!",
		os.Getenv("HELLO_LAMBDA_MSG"))})
	if err != nil {
		return "{}", err
	}
	return string(j), nil
}

func main() {
	lambda.Start(Handler)
}
