Fortunately, there's a tool for hosting Go present talks at:
`https://go-talks.appspot.com/github.com/nmcapule/talks/wasm.slide#1`

Check it out.

## Present

```
$ # http arg not required, but for chromebook, you need it.
$ present -http=0.0.0.0:3999
```

## Run

```
$ ./build.sh
$ go run serve.go
$ # Access <ip from ifconfig>:8080
```
