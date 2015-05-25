# goImgur [![Build Status](https://travis-ci.org/dlion/goImgur.svg?branch=master)](https://travis-ci.org/dlion/goImgur)

> A simple library to upload an image on [imgur](https://imgur.com)

## Install

```
go get github.com/dlion/goImgur
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/dlion/goImgur"
)

func main() {

    str, err := goImgur.Upload("pathToImage.png", "clientID")
    if err != nil {
        log.Panic(err)
    }

    fmt.Println(*str)
}
```

## Prototype

```go
Upload(string, string) (*string, error)
```

## Author

* Domenico Luciani
* http://dlion.it
* domenicoleoneluciani@gmail.com

## License

MIT © [Domenico Luciani](http://github.com/DLion)
