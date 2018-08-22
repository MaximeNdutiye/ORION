provider "aws" {
  region = "${var.aws_region}"
}

resource "aws_lambda_function" "orion" {
  function_name    = "orion-serverless"
  handler          = "orion"
  runtime          = "go1.x"
  filename         = "../../build/lambda.zip"
  source_code_hash = "${base64sha256(file("../../build/lambda.zip"))}"
  role             = "${aws_iam_role.orion_lambda_exec_role.arn}"
  timeout          = 300
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

resource "aws_iam_policy" "lambda-log-policy" {
  name        = "lambda-log-policy"
  description = "Allow lambda to work with log groups and log streams"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:DescribeLogStreams"
    ],
      "Resource": [
        "arn:aws:logs:*:*:*"
    ]
  }
 ]
}
EOF
}

resource "aws_iam_policy" "lambda-s3-policy" {
  name        = "lambda-s3-policy"
  description = "My test policy"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:GetObject",
        "s3:PutObject"
      ],
      "Resource": "arn:aws:s3:::*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda-s3-policy-attachement" {
  role       = "${aws_iam_role.orion_lambda_exec_role.name}"
  policy_arn = "${aws_iam_policy.lambda-s3-policy.arn}"
}

resource "aws_iam_role_policy_attachment" "lambda-logs-policy-attachement" {
  role       = "${aws_iam_role.orion_lambda_exec_role.name}"
  policy_arn = "${aws_iam_policy.lambda-log-policy.arn}"
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

# Bucket for storing Orion config
resource "aws_s3_bucket" "orion_config_bucket" {
  bucket = "orion-config-bucket"
}

resource "aws_s3_bucket_object" "orion_config" {
  key        = "config.json"
  bucket     = "${aws_s3_bucket.orion_config_bucket.id}"
  source     = "config/config.json"
}


resource "aws_s3_bucket_object" "orion_test_image" {
  key        = "image.jpg"
  bucket     = "${aws_s3_bucket.orion_image_bucket.id}"
  source     = "test/image.jpg"
}

# Bucket for storing images
resource "aws_s3_bucket" "orion_image_bucket" {
  bucket = "orion-image-bucket"
}