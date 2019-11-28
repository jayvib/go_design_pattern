package barrier_file_downloader

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestDownloader_Download(t *testing.T) {
	item := &Item{
		url: "https://www.gutenberg.org/ebooks/16328.txt.utf-8",
		processor: func(r *http.Response) error {
			if r.Request.Response.Status != "200" {
				return errors.New("something wrong with the response")
			}
			// check the content-type
			mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
			if err != nil {
				return err
			}
			fmt.Println(mediaType)
			if !strings.Contains(mediaType, ".txt") {
				return errors.New("Invalid file")
			}
			f, err := os.Create("test.txt")
			if err != nil {
				return err
			}
			_, err = io.Copy(f, r.Body)
			if err != nil {
				return err
			}
			return nil
		},
	}
	downloader := NewDownloader(item)
	itemChan := downloader.Download()
	for item := range itemChan {
		if item.err != nil {
			t.Errorf("Error while downloading the %s", item.url)
		}
	}
}
