## Go mod 


```bash
# Create a directory , go to that directory and create a module
mkdir demo
cd demo
go mod init demo
```

### Two kinds of projects

1. Binary/Executable projects
2. Libraries/Packages

## Run 

```bash
go run main.go
# or
go run .
```

## Compilation and build process

- Go Code --> Go Compiler --> Assembler(Convers the compiled code into .o or similar files) -> Linker -> Binary

```bash
#go run binary location
go run --work main.go
```

## Go build

```bash
go build main.go 

# build with a specific name
go build -o demo main.go

# build a binary which is for release mode, stripdown/release binary
go build -o demo-min -ldflags="-s -w" main.go
# It removes symbol information and dwarf information
``` 
