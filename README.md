
# TopLib GO
## A library for making account checkers with Go

## Overview
TopLib is a library for making account checkers with Go. It is a simple,
lightweight, and easy to use library for making account checkers.

With this you can easily make a program that checks the validity your
accounts!




## Installation

```bash
go get github.com/Minagoroshi/TopLibGo
```

## Features
- Proxy Support
- Console UI utils
- Account-List management
- Proxy Checker


## Usage

```go
package main

import "github.com/Minagoroshi/TopLibGo"

func main() {
w := toplib.NewWorker()
w.AddWork(checker)
w.StartWorker()
}

func checker() {
// Your Code Here
}
```


## Disclaimer
While I am aware of the types of software this library 
can be used to make, I highly recommend that you use it 
only for your own accounts, or for your own apis. Any
use of this library for any other purpose is at your own
risk, and I am not responsible for any damages that may
arise from the use of this library.


## Contributions
### Any and all contributions are welcome!
