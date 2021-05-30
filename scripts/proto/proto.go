// +build ignore

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("protoc", "-I", "urlshortener", "urlshortener/service.proto",
		"--descriptor_set_out=urlshortener/service.protoset", "--go_out=urlshortener/proto",
		"--go-grpc_out=urlshortener/proto", "--go_opt=paths=source_relative", "--go-grpc_opt=paths=source_relative",
		"--doc_out=urlshortener/docs", "--doc_opt=html,index.html", "urlshortener/service.proto",
	)
	cmd.Dir = "./../.."
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf(fmt.Sprint(err) + ": " + stderr.String())
	}
}
