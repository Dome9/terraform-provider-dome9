---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_ruleset"
sidebar_current: "docs-resource-dome9-ruleset"
description: |-
  Create ruleset in Dome9
---

# dome9_ruleset

This resource is used to create and manage rulesets in Dome9. Rulesets are sets of compliance rules.

## Example Usage

Basic usage:

```hcl
resource "dome9_ruleset" "ruleset" {
  name        = "some_ruleset"
  description = "this is the descrption of my ruleset"
  cloud_vendor = "aws"
  language = "en"
  hide_in_compliance = false
  is_template = false
  rules {
    name = "some_rule2"
    logic = "EC2 should x"
    severity = "High"
    description = "rule description here"
    compliance_tag = "ct"
  
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the ruleset in Dome9.
* `description` - (Optional) A description of the ruleset (what it represents); defaults to empty string.
* `cloud_vendor` - (Required) Cloud vendor that the ruleset is associated with, can be one of the following: `aws`, `azure` or `google`.
* `language` - (Optional) Language of the rules; defaults to 'en' (English).


### Rules 

The `rules` supports the following arguments:
    
* `name` - (Required) Rule name
* `logic` - (Optional) Rule GSL logic. This is the text of the rule, using Dome9 GSL syntax
* `severity` - (Optional) Rule severity
* `description` - (Optional) Rule description
* `compliance_tag` - (Optional) A reference to a compliance standard


## Attributes Reference

* `id` - Ruleset Id

## Import

Ruleset can be imported; use `<RULE SET ID>` as the import ID. 

For example:

```shell
terraform import dome9_rule_set.test 00000
```
