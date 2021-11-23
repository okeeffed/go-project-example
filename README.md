# Golang Project Template

This is the template to be used over the course of learning Golang.

It is an overkill layout for most projects, but it is being used to get used to the standards defined by the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Getting this space

```s
go mod init github.com/okeeffed/<package-name>
```

Inside of `cmd` there is the `main` package for each entry-point.

Within `pkg` is the public-facing packages that was imported for the example command.

The `Makefile` followed some tips found for having a re-useable build command.

## Installing packages

The example follows commands to install [Cobra](https://github.com/spf13/cobra).

```s
$ go get -u github.com/spf13/cobra
$ cobra init --pkg-name <package-name>
$ cobra add hello # adds hello command

$ go mod vendor
```

## Creating a package

This particular tutorial follows along with ["Everything you need to know about packages"](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc).

## Reading and resources

1. [How To Install and Set Up a Local Programming Environment for Go](https://www.digitalocean.com/community/tutorial_series/how-to-install-and-set-up-a-local-programming-environment-for-go)
2. [Understanding the GOPATH](https://www.digitalocean.com/community/tutorials/understanding-the-gopath)
3. [Go Modules](https://medium.com/mindorks/create-projects-independent-of-gopath-using-go-modules-802260cdfb51)
4. [Makefiles for developers](https://tutorialedge.net/golang/makefiles-for-go-developers/)
