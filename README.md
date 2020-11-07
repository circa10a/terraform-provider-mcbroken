# terraform-provider-mcbroken <img src="https://i.imgur.com/fAS7XqO.png" height="5%" width="5%" align="left"/>

![Build Status](https://github.com/circa10a/terraform-provider-mcbroken/workflows/release/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/circa10a/terraform-provider-mcbroken)](https://goreportcard.com/report/github.com/circa10a/terraform-provider-mcbroken)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/circa10a/terraform-provider-mcbroken?style=plastic)
[![Buy Me A Coffee](https://img.shields.io/badge/BuyMeACoffee-Donate-ff813f.svg?logo=CoffeeScript&style=plastic)](https://www.buymeacoffee.com/caleblemoine)

Base the count of your infrastucture resources on the national average of broken mcdonald's ice machines or by a city of your choosing. Powered by [Mcbroken](https://mcbroken.com/).

- [terraform-provider-mcbroken](#terraform-provider-mcbroken)
  * [Usage](#usage)
  * [Development](#development)
    + [Linting](#linting)
    + [Mac](#mac)
    + [Linux](#linux)
    + [Windows](#windows)

## Usage

- View the provider on the [Hashicorp Registry](https://registry.terraform.io/providers/circa10a/mcbroken/latest/docs)

```hcl
terraform {
  required_providers {
    mcbroken = {
      source  = "circa10a/mcbroken"
    }
  }
}

provider "mcbroken" {}

// Data source to get all available cities/national average of broken ice cream machines
data "mcbroken_cities" "all" {}
// Data source to get current outage percentage of a specific city
data "mcbroken_city" "Dallas" {
    city = "Dallas"
}
// If specified city isn't found, returns -1
data "mcbroken_city" "not_found" {
    city = "not_found"
}
// Get national average of broken ice cream machines
output "global_broken_average" {
    value = data.mcbroken_cities.all.broken
}
// Get list of all cities and their outage percentage
output "all_available_cities" {
    value = data.mcbroken_cities.all.cities
}
// Get outage percentage of a specific city
output "user_specified_city" {
    value = data.mcbroken_city.Dallas.broken
}
// When user specified city isn't found, return -1
output "user_specified_city_not_found" {
    value = data.mcbroken_city.not_found.broken
}

# Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

# Outputs:

# all_available_cities = [
#   {
#     "broken" = 13.04
#     "city" = "New York"
#   },
#   {
#     "broken" = 13.04
#     "city" = "San diego"
#   },
#   {
#     "broken" = 12.5
#     "city" = "Philadelphia"
#   },
#   {
#     "broken" = 11.11
#     "city" = "Boston"
#   },
#   {
#     "broken" = 10.81
#     "city" = "Washington"
#   },
#   {
#     "broken" = 10.53
#     "city" = "Los Angeles"
#   },
#   {
#     "broken" = 9.88
#     "city" = "Chicago"
#   },
#   {
#     "broken" = 8.51
#     "city" = "Phoenix"
#   },
#   {
#     "broken" = 8.11
#     "city" = "Dallas"
#   },
#   {
#     "broken" = 8.05
#     "city" = "Houston"
#   },
#   {
#     "broken" = 7.41
#     "city" = "San Jose"
#   },
#   {
#     "broken" = 6.67
#     "city" = "San Francisco"
#   },
#   {
#     "broken" = 3.77
#     "city" = "San antonio"
#   },
#   {
#     "broken" = 0
#     "city" = "Seattle"
#   },
# ]
# global_broken_average = 7.45
# user_specified_city = 8.11
# user_specified_city_not_found = -1
```

## Development

### Linting

```bash
make lint
```

### Mac

```bash
make build-mac && \
cd examples && \
terraform init && \
terraform apply
```

### Linux

```bash
make build-linux && \
cd examples && \
terraform init && \
terraform apply
```

### Windows

![alt text](https://media.giphy.com/media/4cuyucPeVWbNS/giphy.gif)
