variable "region" {
    type = string
    default = "us-east-1"
}

variable "policy" {
    type = string
    default = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

variable "environment" {
    type = string
    default = "prod" 
}

variable "policy_basic" {
    type = string
    default = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole" 
}