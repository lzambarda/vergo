# Semver

Simple library to parse, format, compare and bump versions based on the [semantic versioning guidelines](https://semver.org).

Compatible with YAML through https://github.com/go-yaml/yaml

## Usage

```go
a := vergo.New(1,2,3)
b := vergo.ParseSemver("v1.2.4")

fmt.Println(a) // 1.2.3
fmt.Println(b) // v1.2.4 preserves the versioning

fmt.Println(a.Before(b)) // true
fmt.Println(a.After(b)) // false

a.Bump(vergo.Major)

fmt.Println(a) // 2.0.0

fmt.Println(a.Before(b)) // false
fmt.Println(a.After(b)) // true
```

Run it on https://go.dev/play/p/E6EyO2LfjtC
