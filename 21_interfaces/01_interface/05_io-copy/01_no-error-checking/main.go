package main

import (
	"strings"
	"io"
	"os"
	"bytes"
	"net/http"
)

func main() {
	msg := "Do not dwell in the past, do not dream of the future, concentrate the mind on the present."
	rdr := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr)

	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)

	res, _ := http.Get("http://www.stanleysvensson.com")
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()
}
