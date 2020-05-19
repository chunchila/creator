# This Program is creating random binary data files

## usage :
```bash
Usage of creator:
  -buffer int
        write buffer size  (default 1024)
  -file string
        file name (default "large.file")
  -size int
        size in MB (default 100)
```

> Please use it with care

> Build Command:
```bash 
go build -ldflags "-w -s -X main.version=v1.0" -o create .
```
