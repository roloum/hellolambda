# Using AWS Lambda with API Gateway

Before I start deploying AWS Lambda functions using the [Serverless](https://serverless.com) framework, I wanted to understand what's happening in the background and set up an API gateway for a basic hellolambda function I created.

## Create lambda function

	aws2 lambda create-function --function-name $functionName --runtime go1.x --zip-file fileb://$zipFileName.zip --handler $functionName --role arn:aws:iam::$accountId:role/$role --region $region

## Create REST API

	aws2 apigateway create-rest-api --name $apiName

## Pull API resources to obtain parent resource's ID and resource's last path segment

	aws2 apigateway get-resources --rest-api-id $restApiId

## Create API resource

	aws2 apigateway create-resource --rest-api-id $restApiId --path-part $resourceLastPathSegment --parent-id $parentResourceId

## Create HTTP method

	aws2 apigateway put-method --rest-api-id $restApiId --resource-id $resourceId --http-method POST --authorization-type NONE

## Integrate HTTP method with Lambda function

	aws2 apigateway put-integration --rest-api-id $restApiId --resource-id $resourceId --http-method POST --type AWS --integration-http-method POST --uri arn:aws:apigateway:$region:lambda:path/2015-03-31/functions/arn:aws:lambda:$region:$accountId:function:$functionName/invocations

## Set up HTTP method response

	aws2 apigateway put-method-response --rest-api-id $restApiId --resource-id $resourceId --http-method POST --status-code 200 --response-models application/json=Empty

## Set up Integration response

	aws2 apigateway put-integration-response --rest-api-id $restApiId --resource-id $resourceId --http-method POST --status-code 200 --response-templates application/json=""

## Deploy API

	aws2 apigateway create-deployment --rest-api-id $restApiId --stage-name prod

## Add permissions to invoke lambda function

	aws2 lambda add-permission --function-name $functionName --statement-id apigateway-prod-2 --action lambda:InvokeFunction --principal apigateway.amazonaws.com --source-arn "arn:aws:execute-api:$region:$accountId:$restApiId/prod/POST/$apiNameManager"
