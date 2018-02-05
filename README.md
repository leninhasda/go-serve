# go-serve
Quickly spin up a file server

## Install and usage:
[Download the binary from here](https://github.com/leninhasda/go-serve/releases/download/v0.1/go-serve)

Add executable permission and start using:
```
chmod +x go-serve
./go-serve
```

You can move the binary to `/usr/local/bin` and access globally:
```
sudo mv go-serve /usr/local/bin/
go-serve
```

By default it will serve current directory port `5000`. But you can modify that:
```
PORT=8080 go-serve // will serve current directory on port 8080
// or
DIR=~/ PORT=8080 go-serve // will serve home directory on port 8080
```
