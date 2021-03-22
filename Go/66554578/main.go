package main

import (
	"io"
	"net/http"
)

var url = "https://fast.com/it/"

func main() {
var bytesRead int64
client := &http.Client{}
req, err := http.NewRequest("GET", url, nil)
if err != nil {
	panic(err)
}
resp, err := client.Do(req)
if err != nil {
	panic(err)
}
defer resp.Body.Close()
nBytes, err := io.Copy(io.Discard, resp.Body)
if err != nil {
	panic(err)
}
bytesRead += nBytes
}
