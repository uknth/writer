# Writer
Rotating File Writer in Golang. Implements io.Writer.

### Usage

```
	wr := writer.NewWriter("app.log", 3600)
	
	// Set log Output
	// This will log in file rotating it after every 1 hour
	log.SetOutput(wr)   
```
