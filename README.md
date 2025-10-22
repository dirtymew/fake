## 📝 Fake - Go Fake Data Generator

This is a fork of the original [`icrowley/fake`](https://github.com/icrowley/fake) with native `go:embed` support and without global variables.

### Key Improvements
- ✨ Native Go embed (no code generation needed)
- 🚀 No globals, no mutexes etc...

### Features
- Multi-language support


### Installation
```shell script
go get github.com/dirtymew/fake
```

## Quick Start
```go

func main(){
    f := fake.New()
    fmt.Println(f.FullName())                  // "John Doe" 
    fmt.Println(f.EmailAddress())              // "john.doe@example.com"
}
```

### Requirements
Go 1.24+

---