package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println(err)
	}
	// byte_slice := make([]byte, 99999)
	// resp.Body.Read(byte_slice) //using Read function of ReaderInterface.
	//the body implements ReaderInterface . in which you pass a byte slice in Read function and
	//it pushes the response body in that byte slice. since slice is passed by reference , we get the
	//response body in the byte slice
	// fmt.Println(string(byte_slice))
	// io.Copy(os.Stdout, resp.Body) // here we are also printing the response body using io.Copy
	//which requires WriterInterface as first arg and Reader interface as second
	// using writerinterface we defined on which channel (terminal,file,etc) you want to print/write the data to
	//using ReaderInterface we defined from where to read the data
	lg := logWriter{} //here we use custom struct which implements Write function so satisfies Writer interface
	io.Copy(lg, resp.Body)
}

func (lg logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
