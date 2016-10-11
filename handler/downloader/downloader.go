package downloader

import (
	"net/http"
	"net/url"

	log "code.google.com/p/log4go"
)

type Downloader struct {
	Method    string
	UrlString string
	Proxy     string
	PostForm  map[string]string
}

func NewDownloader(method, url, proxy string, header map[string]string) *Downloader {
	return &Downloader{
		Method: method,
		Url:    url,
		Proxy:  proxy,
		Header: header,
	}
}

func (dl *Downloader) Do() (resp http.Response, err error) {
	r, err := NewRequest(dl.Method, dl.Url, dl.PostForm)
	if err != nil {
		return nil, err
	}

	for key, value := range dl.Handlers {
		r.Header.Set(key, value)
	}

	proxy, err := url.Parse(dl.proxy)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	resp, err = client.Do(r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
