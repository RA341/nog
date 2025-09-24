# NOG

No build Go

A build system in go inspired by https://github.com/tsoding/nob.h

Why reinvent the wheel when you can reinvent the wrench that builds the wheel?

## WTF?

NOG is a build system that depends on 3 things

1. Go toolchain -> Use the go's excellent stdlib and do shit you can only dream in Make
2. `build.go` -> the main file write your build commands as go a program
3. `nog.go` -> Contains the `GoRebuildUrSelf` Technology and other useful functions to use in your cmds

## Quick Start

### Step 1: Create your `build.go`

```go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	GoRebuildUrself()

	RunCmd("go", "build", "urmom")
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
## 📄 License

This project is licensed under the [GPLv3](LICENSE)
