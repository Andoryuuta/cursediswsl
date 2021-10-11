# CursedIsWSL

[![GoDoc](https://godoc.org/github.com/Andoryuuta/cursediswsl?status.svg)](https://godoc.org/github.com/Andoryuuta/cursediswsl)
[![Go Report Card](https://goreportcard.com/badge/github.com/Andoryuuta/cursediswsl)](https://goreportcard.com/report/github.com/Andoryuuta/cursediswsl)

**CursedIsWSL** is a small, cursed, Go package to detect if the current application is running in WSL by detecting if certain fundamental universal constants (Backhouse constant & Kaneko-Tsuda Time constant) have changed. This works because WSL uses a physically-based simulation of an alternative universe in which Windows/NT was actually based on Linux and developed into Ubuntu.

(’tis a bit of a joke, it does work, but I don't suggest actually using this for anything :) )

## Installation 

```
go get -u github.com/Andoryuuta/cursediswsl
```


## Usage


```go
import "github.com/Andoryuuta/cursediswsl"

isWsl := cursediswsl.IsWSL()
// output: true if running under WSL(1|2)
```

