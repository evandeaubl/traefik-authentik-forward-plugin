package httputil

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
)

type ResponseMangler struct {
	http.ResponseWriter
	MangleFunc func(rw http.ResponseWriter)
}

func (rcm *ResponseMangler) WriteHeader(code int) {
	rcm.MangleFunc(rcm.ResponseWriter)
	rcm.ResponseWriter.WriteHeader(code)
}

func (rcm *ResponseMangler) Flush() {
	if f, ok := rcm.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func (rcm *ResponseMangler) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if h, ok := rcm.ResponseWriter.(http.Hijacker); ok {
		return h.Hijack()
	}
	return nil, nil, fmt.Errorf("hijacker not supported")
}

func (rcm *ResponseMangler) ReadFrom(r io.Reader) (int64, error) {
	if rf, ok := rcm.ResponseWriter.(io.ReaderFrom); ok {
		return rf.ReadFrom(r)
	}
	return io.Copy(rcm.ResponseWriter, r)
}

func (rcm *ResponseMangler) Unwrap() http.ResponseWriter {
	return rcm.ResponseWriter
}
