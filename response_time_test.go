package responsetime

import (
	"github.com/goforgery/forgery2"
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestResponseTime(t *testing.T) {

	Describe("Create()", func() {

		var app *f.Application
		var req *f.Request
		var res *f.Response

		BeforeEach(func() {
			app = f.CreateApp()
			req = f.CreateRequestMock(app)
			res, _ = f.CreateResponseMock(app, false)
		})

		It("should return [3]", func() {
			app.Use(Create())
			app.Handle(req, res, 0)
			AssertEqual(len(res.Writer.Header().Get("X-Response-Time")), 3)
		})

		It("should return [3]", func() {
			app.Use(Create())
			app.Use(func(req *f.Request, res *f.Response, next func()) {
				res.Write("")
			})
			app.Handle(req, res, 0)
			AssertEqual(len(res.Writer.Header().Get("X-Response-Time")), 3)
		})

		It("should return [3]", func() {
			app.Use(Create())
			app.Use(func(req *f.Request, res *f.Response, next func()) {
				res.End("")
			})
			app.Handle(req, res, 0)
			AssertEqual(len(res.Writer.Header().Get("X-Response-Time")), 3)
		})

		It("should return [3, true]", func() {
			app.Use(Create())
			app.Use(func(req *f.Request, res *f.Response, next func()) {
				res.SetHeader("X-Set", "true")
				res.End("")
			})
			app.Handle(req, res, 0)
			AssertEqual(res.Writer.Header().Get("X-Set"), "true")
		})
	})

	Report(t)
}
