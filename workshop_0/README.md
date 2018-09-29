# Hello World

In this section we cover how to write your very first application in GO.

1.  Visit https://golang.org/dl/ and download the latest version of GoLang.

2.  Once you have GoLang succefully installed, open your favorite text editor([VScode](https://code.visualstudio.com/)) and write in the following:

```
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

3.  Now compile the program above into a binary:

```
go build main.go
```

4.  Run the program:

```
./main
```

You can see the program displayed "hello world".

Alternatively, you may run the program from the srouce without building:

```
go run main.go
```
