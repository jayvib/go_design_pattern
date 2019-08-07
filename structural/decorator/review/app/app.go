package app

import (
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// This app is use to review the concept of
// decorator pattern.

func NewLogMiddleware(w io.Writer, h http.Handler) http.Handler {
	l := logrus.New()
	l.Out = w
	return &LogMiddleware{
		h: h,
		l: l,
	}
}

type LogMiddleware struct {
	h http.Handler
	l *logrus.Logger
}

func (l *LogMiddleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	l.l.Printf("Method: %s Time: %v\n", req.Method, time.Now())
	l.h.ServeHTTP(res, req)
}
