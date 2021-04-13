# history cli 

## Overview
history cli is a simple package provider a history list , recall previous history command and execute for any golang cli.
## Basic Usage

Write history items to local before list or privouse item call

```go
package main

import (
    "os",
    "github.com/tedteng/history"
)
func main() {
    pathHistory := "/tmp/history"
    h := history.Settings(pathHistory,"")
    h.Write(os.Args[0:]) 
}
```

list history items and execute the selected item
```go
package main

import (
    "github.com/tedteng/history"
)
func main() {
    pathHistory := "/tmp/history"
    h := history.Settings("","bash")  //binary for executed  eg os.Args[0] 
    h.List()
}
```

execute previous history item.

```go
package main

import (
    "github.com/tedteng/history"
)
func main() {
    pathHistory := "/tmp/history"
    h := history.Settings("",os.Args[0])
    h.Previous()
}
```


The promptui feature use from  [promptui](https://raw.githubusercontent.com/manifoldco/promptui/)




