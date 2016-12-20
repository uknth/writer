# Writer
Rotating File Writer in Golang. Implements io.Writer.
   
[![Build Status](https://travis-ci.org/uknth/writer.svg?branch=master)](https://travis-ci.org/uknth/writer)
[![GoDoc](https://godoc.org/github.com/uknth/writer?status.svg)](https://godoc.org/github.com/uknth/writer)
[![](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
   
### Usage

```
	wr := writer.NewWriter("app.log", 3600)
	
	// Set log Output
	// This will log in file rotating it after every 1 hour
	log.SetOutput(wr)   
```

---
> [uknth.me](http://uknth.me) &nbsp;&middot;&nbsp;
> GitHub [@uknth](https://github.com/uknth) &nbsp;&middot;&nbsp;
> Twitter [@uknth](https://twitter.com/uknth)