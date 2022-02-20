# Go AWS CDK - Lambda(and DynamoDB) example

A little example that shows how to deploy a Cloudformation stack using the Go CDK. The stack consists
of a Lambda and a DynamoDB table. The function writes a message to the table.

Largely based on [LET’S GO ON AWS][go-on-aws] tutorials. Check it out, it's nice!

## Requirements
    
* [Go][go] 1.16+
* [Task][task] v3.11.0

## Setup

1. Go to the project root directory
   
    ```
    cd hello-lambda-cdk-go
    ```

2. Inside `Taskfile.yml`, change the AWS region env vars to the region that you'll be using, for example
   
   ```
    AWS_DEFAULT_REGION: us-east-2
    CDK_DEFAULT_REGION: us-east-2
   ```

3. Get go dependencies.
   
   ```
   task update
   ```

4. Deploy the stack 
   
   ```
   task deploy
   ```

The stack can be deleted with

```
task destroy
```

Also, you can run `task -ls` to obtain a list of all the available tasks.

## TODO

- [ ] Lambda & infrastructure tests.



[go-on-aws]: https://www.go-on-aws.com/
[go]: https://go.dev/
[task]: https://taskfile.dev/#/
[aws-cdk-go]: https://github.com/aws/aws-cdk-go