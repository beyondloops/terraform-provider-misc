# Terraform Provider Utilities

A minimal Terraform provider with a single data source for preserving sensitive value markings in arbitrary data structures.

## Data Source

### `utilities_preserve_sensitivity`

The only data source in this provider. It takes arbitrary input data and echoes it back as output while preserving Terraform's sensitive value markings.

**Example Usage:**

```hcl
terraform {
  required_providers {
    utilities = {
      source  = "beyondloops/utilities"
      version = "~> 1.0"
    }
  }
}

data "utilities_preserve_sensitivity" "example" {
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
  value = data.utilities_preserve_sensitivity.example.input
}
```

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.23 (for development)

## Installation

### From Terraform Registry

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    utilities = {
      source  = "beyondloops/utilities"
      version = "~> 1.0"
    }
  }
}
```

Then run:
```shell
terraform init
```

### Manual Installation

1. Download the appropriate binary from the [releases page](https://github.com/yourorg/terraform-provider-utilities/releases)
2. Place it in your Terraform plugins directory
3. Run `terraform init`

## Building The Provider

1. Clone the repository:
```shell
git clone https://github.com/yourorg/terraform-provider-utilities.git
cd terraform-provider-utilities
```

2. Build the provider:
```shell
go install
```
