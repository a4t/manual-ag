# manual-ag

## Description
manual-ag for change AutoScalingGroup desired tool.

## Usage
```
package main

import (
  "github.com/a4t/manual-ag"
)

func main() {
  client := manualag.Init("synapse-production-spot-cluster", &manualag.Config{
    Region:             "ap-northeast-1",
    AwsAccessKeyId:     "XXXXXXXXXXXXXXXXXX",
    AwsSecretAccessKey: "YYYYYYYYYYYYYYYYYYYYYYYYY",
  })

  _, err := client.Change(8, "add or set")
}
```

### add or set ?
- add
  - Now instance count add number
- set
  - Specify instance count

## Install

To install, use `go get`:

```bash
$ go get -d github.com/a4t/manual-ag
```

## Contribution

1. Fork ([https://github.com/a4t/manual-ag/fork](https://github.com/a4t/manual-ag/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run `go fmt`
1. Create a new Pull Request

## Author

[a4t](https://github.com/a4t)
