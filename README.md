# hash-tool

A simple tool which makes http requests and prints the address of the request along with the MD5 hash of the response.

## How to run

If you have go installed, you can run this tool running this command: 
```
go run ./cmd/hash-tool
```

You can also run it with next command:
```
./bin/hash-tool
```
Usually you don't commit bin files, but it's done here for an easier use. Keep in mind that this bin file is compiled for macOS systems, `GOOS=darwin, GOARCH=amd64`. You can compile it yourself with next commands:
```
make build
```
or
```
go build -o ./bin/hash-tool ./cmd/hash-tool
```

**Note** that this tool is making parallel requests for better performance, you can limit their number with `-parallel` flag

## Examples of use

```
./bin/hash-tool google.com

./bin/hash-tool http://ya.ru

./bin/hash-tool -parallel 3 google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
```
