package session

import "net/http"

type StandardClient struct {
}

func NewStandardClient() (*StandardClient, error) {
	return &StandardClient{}, nil
}

func (c *StandardClient) Get(cookies []*http.Cookie) *Session {
	return nil
}

func (c *StandardClient) Set(cookies []*http.Cookie, meta *Session) {
}

func (c *StandardClient) Delete(cookies []*http.Cookie) {
}
