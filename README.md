# Simple logger library

## 1. Installation
> Require: Go 1.15+ version with module enabled

To import the library into you code, use `go get` command:
```shell
go get gitlab.com/with-junbach/go-modules/log
```

## 2. Quick start

Initialize logger global instance and use it everywhere
```go
package app
import "gitlab.com/with-junbach/go-modules/log"

func main() {
	var conf = log.Config{
		Level:     "debug",
		Output:    []string{"stdout"},
		ErrOutput: []string{"stdout", "stderr"},
		Format:    "console",
		Module:    "test-module",
	}

	if err := log.InitLoggerInst(conf); err != nil {
		// handle error in case of failed initiation
		panic(err)
	}
	
	// flush the log buffer when the app shut down
	defer log.Close()
	
	// At this point you can call log functions at any place in your application
	log.Info("this is web app running at port %d", 8080)
}
```

> Note: Please refer the `xxx_test.go` files for more examples of using the library

> You can use the library `gitlab.com/with-junbach/go-modules/config` to load the config more flexible from config file, which supports to load from `.yaml` `.json` and environment variables
## 3. Log with more complex data
Beside plain text message, you can use logger library to log more structured data (eg. `map` or `json`). This supports for both format `console` and `json`
To customize the data you want to log, use the `log.Print` or `log.PrintMap` function

```go
# after initiate log instance
fields := make(map[string]interface{})
fields["request_id"] = "1234-5678-abcd"
fields["protocol"] = "http"
fields["user_agent"] = "postman"
fields["remote_ip"] = "::1"
fields["status"] = http.StatusOK
log.PrintMap(log.InfoLevel, fmt.Sprintf("request returned with status: %d", fields["status"]), fields)
log.Print(log.DebugLevel, "example to print out many data in log", "server", "localhost", "port", 1313, "app_name", "awesome-web-app")
```

## 4. Log with context
This feature is used to support print out log in handle request process. It will attach `request_id` field found in the context into log message, helps easier to track and debug. 
* To enable this feature, we should import a function to logger to "guide" it how to extract `request_id` field from an input `context.Context`.
    ```go
  # run after initiate global log instance
    const RequestIDKey = 0
    var getReqIdFunc = func(ctx context.Context) string {
		return ctx.Value(RequestIDKey).(string)
	}
  log.ImportGetRequestIDFunction(getReqIdFunc)
  log.KInfo(ctx, "got data with input id is %s", itemID)
  
  # log with extra data attached to field object_data
  log.KDebug(ctx, map[string]string{"status": "not_ready"}, "log with some object")

    ```
  > To make this work, make sure your context is implement the interface `context.Context` and there is a proper `request_id` field attach to it matching with your input `getReqIdFunc`
  
  > If you use echo, you can use the library at `gitlab.com/with-junbach/go-modules/echo` with some useful and simple middleware to help achieve that.


* It also helps to remove some unwanted fields you want to omit when log object data (eg. sensitive information likes `username` and/or `password`)
```go
type RequestData struct {
	Flag      bool
	SecretKey string
}

func (req RequestData) Customize() interface{} {
	ref := req
	if ref.SecretKey != "" {
		ref.SecretKey = "_hidden_"
	}
	return ref
}
# run after initiate global log instance
log.KDError(ctx, RequestData{Flag: false, SecretKey: "very-secret"}, "log context")
```
