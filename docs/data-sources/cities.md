# mcbroken_cities Data Source

Data source to get all available cities/national average of broken ice cream machines

## Example Usage

```hcl
data "mcbroken_cities" "all" {}

// Get national average of broken ice cream machines
output "global_broken_average" {
    value = data.mcbroken_cities.all.broken
}

// Get list of all cities and their outage percentage
output "all_available_cities" {
    value = data.mcbroken_cities.all.cities
}
```

## Argument Reference

None

## Attribute Reference

* `cities` - List of maps that contains cities and their broken ice cream machine average
* `broken` - National average of broken ice cream machines
