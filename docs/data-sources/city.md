# mcbroken_city Data Source

// Get outage percentage of a specific city

## Example Usage

```hcl
// Data source to get current outage percentage of a specific city
data "mcbroken_city" "Dallas" {
    city = "Dallas"
}

// Get outage percentage of a specific city
output "user_specified_city" {
    value = data.mcbroken_city.Dallas.broken
}

// If specified city isn't found, returns -1
data "mcbroken_city" "not_found" {
    city = "not_found"
}

// When user specified city isn't found, return -1
output "user_specified_city_not_found" {
    value = data.mcbroken_city.not_found.broken
}
```

## Argument Reference

* `city` - (Required) Name of city to get percentage of broken ice cream machines

## Attribute Reference

* `broken` - Percentage of broken ice cream machines in the user specified city. `-1` if no data
