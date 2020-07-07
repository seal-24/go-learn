//http package hhh
package httpsvr

import (
	"fmt"
	"io"
	"net/http"
)

//
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello")
}

// fsa


// NewRequest returns a new incoming server Request, suitable
// for passing to an http.Handler for testing.
//
// The target is the RFC 7230 "request-target": it may be either a
// path or an absolute URL. If target is an absolute URL, the host name
// from the URL is used. Otherwise, "example.com" is used.
//
// The TLS field is set to a non-nil dummy value if target has scheme
// "https".
//
// The Request.Proto is always HTTP/1.1.
//
// An empty method means "GET".
//
// The provided body may be nil. If the body is of type *bytes.Reader,
// *strings.Reader, or *bytes.Buffer, the Request.ContentLength is
// set.
//
// NewRequest panics on error for ease of use in testing, where a
// panic is acceptable.
//
// To generate a client HTTP request instead of a server request, see
// the NewRequest function in the net/http package.
func Example() {
	fmt.Println("Hello OverView")

	// Output:
	// Hello OverView
}

//godoc -http=:6060   生成文档
//func main() {
//	Example()
//	http.HandleFunc("/hello", HelloHandler)
//	http.ListenAndServe(":12345", nil)
//
//}
