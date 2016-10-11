package scheduler

import (
	"github.com/cwiggers/crawler/downloader"
	"github.com/cwiggers/crawler/misc"

	"gopkg.in/mgo.v2/bson"
)

type Scheduler struct {
	name string
}

func (s *Scheduler) Push(dl *downloader.Downloader) (err error) {
	db := misc.Backend.Db.Copy()
	c := db.DB("scheduler").C(s.name)

	if err = c.Insert(dl); err != nil {
		return
	}
}

func (s *Scheduler) Poll() (dl downloader.Downloader, err error) {
	db := misc.Backend.Db.Copy()
	c := db.DB("scheduler").C(s.name)

	cond := bson.M{"finish": 0}
	filter := bson.M{"_id": 0, "finish": 0}
	if err = c.Find(cond).Select(filter).One(&dl); err != nil {
		return
	}
}
