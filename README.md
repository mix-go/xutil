> Produced by OpenMix: [https://openmix.org](https://openmix.org/mix-go)

## Mix XUtil

A set of tools that keep Golang sweet.

## Installation

```
go get github.com/mix-go/xutil
```

## xhttp [[more]](xhttp/README.md)

| Function                                                                         | Description                      |  
|----------------------------------------------------------------------------------|----------------------------------|
| xhttp.Request(method string, u string, opts ...RequestOption) (*Response, error) | Execute an http request.         |
| xhttp.Do(req *http.Request, opts ...RequestOption) (*Response, error)            | Execute an http request.         |
| xhttp.WithBody(body Body) RequestOption                                          | Set configuration item           |
| xhttp.WithHeader(header http.Header) RequestOption                               | Set configuration item           |
| xhttp.WithContentType(contentType string) RequestOption                          | Set configuration item           |
| xhttp.WithTimeout(timeout time.Duration) RequestOption                           | Set configuration item           |
| xhttp.WithDebugFunc(f DebugFunc) RequestOption                                   | Set configuration item           |
| xhttp.WithRetry(f RetryIfFunc, opts ...retry.Option) RequestOption               | Set configuration item           |
| xhttp.BuildJSON(v interface{}) Body                                              | Generate json string             |
| xhttp.BuildQuery(m map[string]string) Body                                       | Generate urlencoded query string |

## xslices

| Function                                              | Description                                          |  
|-------------------------------------------------------|------------------------------------------------------|
| xslices.InArray[T comparable](item T, slice []T) bool | Searches if the specified value exists in the array. |

## xstrings

| Function                                                   | Description                                                    |  
|------------------------------------------------------------|----------------------------------------------------------------|
| xstrings.IsNumeric(s string) bool                          | Used to check if the variable is a number or a numeric string. |
| xstrings.SubString(s string, start int, length int) string | Return part of a string                                        |

## xconv

| Function                                                | Description                       |  
|---------------------------------------------------------|-----------------------------------|
| xconv.StructToMap(i interface{}) map[string]interface{} | Convert struct to map.            |
| xconv.StringToBytes(s string) []byte                    | Convert string to bytes (0 copy). |
| xconv.BytesToString(b []byte) string                    | Convert bytes to bytes (0 copy).  |

## xcrypt

| Function                                                            | Description    |  
|---------------------------------------------------------------------|----------------|
| xcrypt.AESEncrypt(plainText, mode, key, iv string) (string, error)  | AES encryption |
| xcrypt.AESDecrypt(cipherText, mode, key, iv string) (string, error) | AES Decryption |

## xfmt [[more]](xfmt/README.md)

A formatting library that can print data inside nested pointer addresses of structures.

The supported methods are identical to the `fmt` system library

| Function                                                | Description                     |  
|---------------------------------------------------------|---------------------------------|
| xfmt.Sprintf(format string, args ...interface{}) string |                                 |
| xfmt.Sprint(args ...interface{}) string                 |                                 |
| xfmt.Sprintln(args ...interface{}) string               |                                 |
| xfmt.Printf(format string, args ...interface{})         |                                 |
| xfmt.Print(args ...interface{})                         |                                 |
| xfmt.Println(args ...interface{})                       |                                 |
| xfmt.Disable()                                          | Equivalent to fmt when disabled |
| xfmt.Enable()                                           |                                 |

## xenv [[more]](xenv/README.md)

Environment configuration library with type conversion.

| Function                                  | Description |  
|-------------------------------------------|-------------|
| err := xenv.Load(".env")                  |             |
| err := xenv.Overload(".env")              |             |
| i := xenv.Getenv("key").String("default") |             |
| i := xenv.Getenv("key").Bool(false)       |             |
| i := xenv.Getenv("key").Int64(123)        |             |
| i := xenv.Getenv("key").Float64(123.4)    |             |

## License

Apache License Version 2.0, http://www.apache.org/licenses/
