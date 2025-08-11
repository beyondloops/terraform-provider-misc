# Terraform Provider Misc

A minimal Terraform provider with a single data source for echoing arbitrary data while preserving sensitive value markings.

## Data Source

### `misc_echo`

The only data source in this provider. It takes arbitrary input data and echoes it back while preserving Terraform's sensitive value markings.

**Example Usage:**

```hcl
terraform {
  required_providers {
    misc = {
      source  = "beyondloops/misc"
      version = "~> 1.0"
    }
  }
}

data "misc_echo" "example" {
  input = {
    public_key  = "visible_value"
    private_key = sensitive("secret_value")
    nested = {
      api_url   = "https://api.example.com"
      api_token = sensitive("secret_token")
    }
  }
}

output "processed_data" {
  value = data.misc_echo.example.output
}