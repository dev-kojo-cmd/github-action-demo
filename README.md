# github-action-demo

A tiny Go demo with 3 simple methods and unit tests.

## Methods

- `Add(a, b int) int`
- `Multiply(a, b int) int`
- `IsEven(n int) bool`

## Install Go (if needed)

Go is already installed on this machine (`go1.25.4`).

For a fresh Windows machine, you can install with:

```powershell
winget install -e --id GoLang.Go
```

Then verify:

```powershell
go version
```

## Run tests

From the project folder:

```powershell
go test ./...
```

To run tests with verbose output:

```powershell
go test -v ./...
```
