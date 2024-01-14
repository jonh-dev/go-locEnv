<h1 align="center"> Go-LocEnv </h1>

<p align="center">A library for managing environment variables in Go applications.</p>

<p align="center">
 <a href="#technologies">Technologies</a> •
 <a href="#running">Running</a> •
 <a href="#author">Author</a>
</p>

<p align="center">
  <img src="https://github.com/jonh-dev/go-locEnv/assets/101439670/46dc189f-312a-4e47-8307-23d43f308c37" />
</p>

##

### Technologies

The following tools were used in building the project:

- Go
- IDE Visual Studio Code

##

### Running

**1.** First, you need to install Go on your system. You can do this by following the instructions at the following link: https://golang.org/dl/

**2.** Choose an IDE of your choice, in this case we will use Visual Studio Code. To download it follow the link: https://code.visualstudio.com/download

**3.** Open your terminal and use `go get` to download and install the `go-envloader` library. Replace `github.com/jonh-dev/go-envloader` with the path to your `go-envloader` library:

```bash
$ go get github.com/jonh-dev/go-locEnv
```

**4.** Now you can import the go-envloader library into your Go project. Here's an example of how you can do this:

```Go
import (
    "github.com/jonh-dev/go-locEnv/config"
)
```

**5.** Now you can use the go-envloader library in your code. Here’s an example of how you can do this:

```Go
func main() {
    envLoader := config.NewEnvLoader()

    err := envLoader.LoadEnv()
    if err != nil {
        log.Fatal(err)
    }

    env := envLoader.GetEnv()
    log.Info("Current environment: ", env)
}
```

**6.** Finally, run the project using the command go run main.go to run the application.

##

### Author

![avatar](https://user-images.githubusercontent.com/101439670/181940218-4f68ffb9-0d35-40df-b8e9-86629333d244.png)

Made by Jonh Dev 🙏

[![LinkedIn Badge](https://img.shields.io/badge/-LINKEDIN-blue?style=flat-square&logo=Linkedin&logoColor=white&link="https://www.linkedin.com/in/jo%C3%A3o-carlos-schwab-zanardi-752591213/)](https://www.linkedin.com/in/jo%C3%A3o-carlos-schwab-zanardi-752591213/)

