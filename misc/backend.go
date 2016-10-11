package misc

import (
	log "code.google.com/p/log4go"
	"gopkg.in/mgo.v2"
)

type backend struct {
	Db mgo.Session
}

var Backend *backend

func InitBackend() error {
	sess, err := mgo.Dial(Conf.Mongo.Addr)
	if err != nil {
		log.Warn("Connect to DB failed, Err:[%s]", err)
		return err
	}
	sess.SetMode(mgo.Monotonic, true)
	Backend = &backend{
		Db: sess,
	}
	return nil
}
