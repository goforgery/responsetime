package responsetime

import (
    "github.com/goforgery/forgery2"
	"strconv"
	"time"
)

// Adds the X-Response-Time header displaying the response duration in milliseconds.
//
// Example:
//
//    stackr.CreateServer().Use(stackr.ResponseTime())
func Create() func(*f.Request, *f.Response, func()) {
	return func(req *f.Request, res *f.Response, next func()) {
		start := time.Now().UnixNano()
		res.On("header", func() {
			duration := time.Now().UnixNano() - start
			res.SetHeader("X-Response-Time", strconv.FormatInt(int64(time.Duration(duration)/time.Millisecond), 10)+"ms")
		})
		next()
	}
}
