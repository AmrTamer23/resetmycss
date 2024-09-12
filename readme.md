# 🎨 resetmycss

A tool to reset the CSS for you project using Go and Node.js.

## Usage

You can run this tool using `npx`:

```bash
npx resetmycss
```

## Contributing

Contributions are welcome! Please feel free to submit a PR.

To run the tool locally, you need to have Go installed on your system.

### 1. Build the Go binary

```bash
GOOS=linux GOARCH=amd64 go build -o resetmycss-linux main.go
GOOS=darwin GOARCH=arm64 go build -o resetmycss-macos main.go
GOOS=windows GOARCH=amd64 go build -o resetmycss-windows.exe main.go
```

### 2. Run the tool

```bash
node ./index.js
```

## References

- [Best CSS Reset](https://github.com/Lazzzer00/Best-CSS-Reset-2024)
