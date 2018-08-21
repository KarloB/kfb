package firebase

import (
	"time"

	"github.com/karlob/fbk/internal/model/conf"

	"github.com/zabawaba99/fireauth"
	"github.com/zabawaba99/firego"
)

// Client firebase client
type Client struct {
	firebase *firego.Firebase
}

// WatcherFunc func to be usd in watcher
type WatcherFunc interface {
	Do(eventType string, data interface{}) error
}

// New init client
func New(conf *conf.Configuration) (*Client, error) {
	f := firego.New(conf.DB, nil)
	gen := fireauth.New(conf.Key)
	data := fireauth.Data{"uid": "1"}
	token, err := gen.CreateToken(data, nil)
	if err != nil {
		return nil, err
	}
	f.Auth(token)

	firego.TimeoutDuration = time.Duration(conf.Timeout) * time.Second

	fb := &Client{
		firebase: f,
	}

	return fb, nil
}

// Watcher init new watcher
func (t *Client) Watcher(work WatcherFunc, stop chan (struct{})) error {
	nots := make(chan firego.Event)
	err := t.firebase.Watch(nots)
	if err != nil {
		return err
	}
	defer t.firebase.StopWatching()

	for {
		select {
		case e := <-nots:
			err = work.Do(e.Type, e.Data)
			if err != nil {
				return err
			}
		case <-stop:
			return nil
		}
	}
}
