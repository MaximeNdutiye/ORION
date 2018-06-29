# Describing the resources that will be created
# Provider

provider "aws" {
    region = "${var.aws_region}"
}

# Cloudwatch event rule
# Tells the lambda when to run
resource "aws_cloudwatch_event_rule" "check-file-event" {
}

# Cloudwatch event target
resource "aws_cloudwatch_event_target" "check-file-event-lambda-target" {

}

# IAM Role for Lambda function
resource "aws_iam_role" "orion_lambda_role" {
}

# AWS Lambda function
resource "aws_lambda_function" "orion" {
    filename = "orion.zip"
    function_name = "orion"
}
