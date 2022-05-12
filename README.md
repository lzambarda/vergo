# Semver

Simple library to parse, format, compare and bump versions based on the [semantic versioning guidelines](https://semver.org).

Compatible with YAML through https://github.com/go-yaml/yaml

## Usage

```go
a := vergo.New(1, 2, 3, "")
b, _ := vergo.ParseSemver("v1.2.4")

fmt.Println(a) // 1.2.3
fmt.Println(b) // v1.2.4 preserves the versioning

fmt.Println(a.Before(b)) // true
fmt.Println(a.After(b))  // false

a.Bump(vergo.BumpMajor)

fmt.Println(a) // 2.0.0

fmt.Println(a.Before(b)) // false
fmt.Println(a.After(b))  // true

c := vergo.New(1, 0, 0, "rc1")         // or -rc1
_ = c.Bump(vergo.BumpReleaseCandidate) // will error if it is not rc<NUMBER>
fmt.Println(c)                         // 2.0.0-rc2 nil
```

Run it on https://go.dev/play/p/E6EyO2LfjtC

## Semver labels

Vergo has limited support for semver with labels (e.g. `v1.0.0-rc1`).
It can parse any sort of label but can only apply a bump to labels in the form of `rc<NUMBER>`.
