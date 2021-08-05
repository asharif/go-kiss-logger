# go-kiss-logger

A super awesome logger that aims to keep it simple, stupid!

# How to use it?

In your go.mod file add:

```
require (
        github.com/asharif/go-kiss-logger v1.0.8
)
```

In your code add:

```go
import (
        log "github.com/asharif/go-kiss-logger"
)
```
...

```go
//examples
logger := log.GetInstance()
logger.Info("This is an INFO level log")
logger.Warn("This is an WARN level log")
logger.Error("This is an ERROR level log")
logger.Fatal("This is an FATAL level log. It will exit with the provided code", 1)
```
Output will look like:
```
2020-12-02T09:05:48.789Z | INFO | ...sers/asharif/development/go-kiss-logger/main.go:10 | This is an INFO level log
2020-12-02T09:05:48.789Z | WARN | ...sers/asharif/development/go-kiss-logger/main.go:11 | This is an WARN level log
2020-12-02T09:05:48.789Z | ERROR | ...sers/asharif/development/go-kiss-logger/main.go:12 | This is an ERROR level log
main.main()
        /Users/asharif/development/go-kiss-logger/main.go:12 +0x8f
2020-12-02T09:05:48.789Z | FATAL | ...sers/asharif/development/go-kiss-logger/main.go:13 | This is an FATAL level log. It will exit with the provided code
main.main()
        /Users/asharif/development/go-kiss-logger/main.go:13 +0xbb
exit status 1
```
Note that INFO and WARN will write to the stdout, ERROR and FATAL will write to the stderr.

# How to bump version numbers

See the .github/workflows/main.yaml. Specifically the following section:

```
default_bump: major #values major, minor, patch. default is patch
```
