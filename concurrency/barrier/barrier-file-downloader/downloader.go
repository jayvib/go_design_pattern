package barrier_file_downloader

import (
	"mime"
	"net/http"
	"sync"
	"time"
)

// This is an attempt to implement a downloader using the barrier pattern concept.
// Barrier Design Pattern will put up a barrier so that nobody passes until we have all the results we need

var defaultTimeout = time.Duration(2 * time.Minute)

type Response struct {
	resp *http.Response
	err error
}

type Item struct {
	Response
	url string
	processor func(r *http.Response) error
}

type Downloader interface {
	Download() <-chan Item
}

type downloader struct {
	timeout time.Duration
	pool int
	items []*Item
	lock sync.Mutex
}

func (d *downloader) Download() <-chan Item {

	return nil
}

func isValidFileRequest(r *http.Response) bool {
	mime.ParseMediaType(r.Header.Get("Content-Type"))
	return false
}

func NewDownloader(items ...*Item) *downloader {
	i := make([]*Item, 0)
	if items != nil {
		i = append(i, items...)
	}
	return &downloader{
		timeout: defaultTimeout,
		pool: 1,
		items: i,
	}
}

func (d *downloader) download(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: d.timeout,
	}
	http.nEW
	client.Do()
}