```shell
    go get github.com/aws/aws-lambda-go/lambda 
    GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
    zip myFunction.zip bootstrap
    
     AWS_PROFILE=PROFILE aws iam create-role --role-name lambda-ex --assume-role-policy-document '{"Version": "2012-10-17","Statement": [{ "Effect": "Allow", "Principal": {"Service": "lambda.amazonaws.com"}, "Action": "sts:AssumeRole"}]}'
     
     AWS_PROFILE=PROFILE aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
     
     AWS_PROFILE=PROFILE aws iam attach-role-policy  --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

     
     AWS_PROFILE=PROFILE aws lambda create-function --function-name myFunction \
        --runtime provided.al2 --handler bootstrap \
        --architectures arm64 \
        --role arn:aws:iam::ROLE_ID:role/lambda-ex \
        --zip-file fileb://myFunction.zip

    AWS_PROFILE=PROFILE aws lambda invoke --function-name myFunction --cli-binary-format raw-in-base64-out --payload '{"what is your name?": "Mihail", "How old are you?": 28}' output.txt
    
    AWS_PROFILE=PROFILE aws lambda update-function-code --function-name myFunction --zip-file fileb://myFunction.zip
```