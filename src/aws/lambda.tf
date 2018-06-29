provider "aws" {
  region = "us-east-1"
}

resource "aws_lambda_function" "orion" {
  function_name = "orion"

  # The bucket name as created earlier with "aws s3api create-bucket"
  s3_bucket = "orion-lambda-source"
  s3_key    = "v1.0.0/example.zip"

  # "orion" is the filename within the zip file (main.js) and "handler"
  # is the name of the property under which the handler function was
  # exported in that file.
  handler = "orion.handler"
  runtime = "go1.x"

  filename         = "../../build/orion.zip"

  role = "${aws_iam_role.orion_lambda_exec_role.arn}"
}

# IAM role which dictates what other AWS services the Lambda function
# may access.
resource "aws_iam_role" "orion_lambda_exec_role" {
  name = "orion_lambda_exec_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_permission" "orion-lambda-permission" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.orion.arn}"
  principal     = "apigateway.amazonaws.com"

  # The /*/* portion grants access from any method on any resource
  # within the API Gateway "REST API".
  source_arn = "${aws_api_gateway_deployment.orion-api-gateway-deployment.execution_arn}/*/*"
}