# Go
First steps using Go language  

## Where to start?

* [Tutorial: Get started with Go][getting-started]

    * [Go Dev page: Download and install][download1]
    * Prerequisites
        * **A tool to edit your code.** I have chosen VSCode (free) because I had already installed it.
    

## Hello

1 - Open a terminal and go to your root working directory

```
% cd
% cd VSCodeProjects/GoProjects

```

2 -  Create a folder for this new code

```
% mkdir hello
% cd hello
```

3 - Enable dependency tracking for your code.

```
hello % go mod init example/hello
```

4 - In your text editor, create a file hello.go in which to write your code.

    When I created the file using VSCode, the IDE  detected a new Go file and asked to install plugins.


    > Go for VS Code v0.38.0
    The official Go extension for Visual Studio Code, providing rich language support for Go projects.

    Installing 7 tools at /Users/alvaro.menezes/go/bin in module mode.
        gotests
        gomodifytags
        impl
        goplay
        dlv
        staticcheck
        gopls

5 - Paste the following code into your hello.go file and save the file.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

6 - Run your code to see the greeting.
```go
$ go run .
Hello, World!
```

[download1]: https://go.dev/doc/install
[getting-started]: https://go.dev/doc/tutorial/getting-started