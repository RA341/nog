# NOG

No build Go

A build system in go inspired by https://github.com/tsoding/nob.h

Why reinvent the wheel when you can reinvent the wrench that builds the wheel?

## WTF?

NOG is a build system written in `build.go` and `nog.go` define your build commands and whatever else you need.

Use the go's excellent stdlib to do shit you can only dream in Make.

## 🚀 Quick Start

### Step 1: Create your `build.go`

```go
package main

import "fmt"

func main() {
	GoRebuildUrself() // The magic happens here ✨
	fmt.Println("How's your noggin?")
    
	// Add your actual build logic here
}
```

### Step 2: Download the nog.go

Grab [nog.go](https://raw.githubusercontent.com/RA341/nog/refs/heads/main/nog.go) and put it next to your `build.go`.

### Step 3: Bootstrap

**Linux/macOS:**

```bash
go build -o nog nog.go build.go
```

**Windows:**

```bash
go build -o nog.exe nog.go build.go 
```

### Step 4: Go nog

Watch in amazement as your build system rebuilds itself everytime you update your `build.go` or `nog.go`

```bash
./nog
```

## Features

- **Self-Rebuilding**
- **Zero Dependencies**: Do you REALLY need anything else ???
- **Cross-Platform**: Works on Linux, macOS, Windows
- **Minimal Setup**: Three files and you're good to NOG

## FAQ

**Q: Is this stupid?**  
A: Yes.

**Q: Is this production-ready?**  
A: Yes.

**Q: Can I use this for serious projects?**  
A: Yes.

## Contributing

Found a bug? Want to make NOG even more unnecessarily complex? PRs welcome!

Just remember: We're not trying to replace Make, Ninja, Bazel, or any other perfectly good build system. We're here to
have fun and maybe learn something along the way.

## 📄 License

This project is licensed under the "Do Whatever You Want But Don't Blame Us" license. Use at your own risk and
existential peril.
