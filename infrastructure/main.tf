provider "aws" {
  region = var.region
}

resource "aws_iam_role" "lambda_role" {
  name = "go_lambda_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy_attach" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = var.policy
}

resource "aws_iam_role_policy_attachment" "lambda_basic" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = var.policy_basic
}

resource "aws_lambda_function" "go_lambda" {
  depends_on = [
    aws_iam_role_policy_attachment.lambda_policy_attach,
    aws_iam_role_policy_attachment.lambda_basic
  ]

  function_name    = "go_lambda_function"
  role             = aws_iam_role.lambda_role.arn
  handler          = "bootstrap"
  runtime          = "provided.al2"
  filename         = "../backend/build/lambda.zip"
  source_code_hash = filebase64sha256("../backend/build/lambda.zip")
  timeout          = 10
  memory_size      = 128

}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.go_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.api.execution_arn}/*/*"
}
