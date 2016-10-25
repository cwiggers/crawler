package downloader

import (
	"io/ioutil"
	"net/http"

	log "github.com/alecthomas/log4go"
)

type Downloader struct {
	Url          string `json:"url"`
	NextPage     string `json:"next_page"`
	SubPageUrl   string `json:"sub_page_url"`
	Body         []byte `json:"body"`
	SubPageModel struct {
		Title         []string `json:"title"`
		Text          []string `json:"text"`
		CreateTime    []string `json:"create_time"`
		ReprintNumber []string `json:"reprint_number"`
	} `json:"sub_page_model"`
}

func (dl *Downloader) Download() {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(dl.Url)
	if err != nil {
		log.Warn("failed to get %s, err:[%s]", dl.Url, err)
		return
	}

	if resp.StatusCode != 200 {
		log.Warn("failed to get %s, status code:[%d:%s]", dl.Url, resp.StatusCode, resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn("failed to get body from %s", dl.Url)
		return
	}
	// websocket
}
