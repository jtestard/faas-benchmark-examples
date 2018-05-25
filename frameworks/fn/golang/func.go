package main

import (
	"context"
	"encoding/json"
	"io"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	json.NewEncoder(out).Encode(struct {
		Msg string `json:"message"`
	}{Msg: "Hello from Fn!"})
}
