provider "aws" {
  region                      = "us-east-1"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  skip_get_ec2_platforms      = true
  skip_region_validation      = true
  access_key                  = "mock_access_key"
  secret_key                  = "mock_secret_key"
}

resource "aws_ssm_parameter" "ssm_parameter_advanced" {
  name = "my-advanced-ssm-parameter"
  type = "String"
  value = "Advanced Parameter"
  tier = "Advanced"
}

resource "aws_ssm_parameter" "ssm_parameter_advancedWithUsage" {
  name = "my-advanced-ssm-parameter"
  type = "String"
  value = "Advanced Parameter"
  tier = "Advanced"
}
