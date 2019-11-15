package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

//Handler : Lambda handler invoked by lambda.Start
func Handler() (string, error) {

	ses, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")})
	if err != nil {
		return "{}", err
	}

	svc := ssm.New(ses)
	paramName := "hellolambda.ParameterStore"
	decryption := false
	param, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           &paramName,
		WithDecryption: &decryption})
	if err != nil {
		return "{}", err
	}

	j, err := json.Marshal(map[string]interface{}{"message": fmt.Sprintf(
		"Hello %v! %v!! Your function executed successfully!",
		os.Getenv("HELLO_LAMBDA_MSG"), *param.Parameter.Value)})
	if err != nil {
		return "{}", err
	}
	return string(j), nil
}

func main() {
	lambda.Start(Handler)
}
