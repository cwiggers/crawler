package processer

import (
	"regex"

	"github.com/PuerkitoBio/goquery"
)

type Processer struct {
	Url           string   `json:"url"`
	Body          []byte   `json:"body"`
	Query         string   `json:"query"`
	Title         []string `json:"title"`
	Text          []string `json:"text"`
	CreateTime    []string `json:"create_time"`
	ReprintNumber []string `json:"reprint_number"`
}

func (pr *Processer) Process() {

}
