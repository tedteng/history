# history cli 

## Overview
history cli is a simple package provider a history list , recall previous history command and execute for any golang cli.

## Basic Usage

```go
package main

import (
    "https://github.com/tedteng/history"
)
func main() {
    pathHistory := "/tmp/history"
    h := history.Settings(pathHistory)
    
    // list history records and executed after select item
    if list {
        h.List()
    } else {
    // execute Previous history item
        h.Previous()
    }
}
```
The promptui feature use from  [promptui](https://raw.githubusercontent.com/manifoldco/promptui/)




