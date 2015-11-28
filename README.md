# responsetime

[![Build Status](https://secure.travis-ci.org/goforgery/responsetime.png?branch=master)](http://travis-ci.org/goforgery/responsetime)

Response time logger for Forgery2.

## Use

Sets the header `X-Response-Time` with the time in milliseconds that __Forgery2__ took to process the request.

```javascript
package main

import (
	"github.com/goforgery/forgery2"
	"github.com/goforgery/responsetime"
)

func main() {
	app := f.CreateApp()
	app.Use(responsetime.Create())
	app.Listen(3000)
}
```

## Test

    go test
