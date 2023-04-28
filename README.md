> OpenMix 出品：[https://openmix.org](https://openmix.org/mix-go)

## Mix XUtil

一套让 Golang 保持甜美的工具。

A set of tools that keep Golang sweet.

## Installation

```
go get github.com/mix-go/xutil
```

## List of functions

执行http请求

```
http.Request(method string, u string, body string, opts ...Options) ([]byte, error)
```

```
http.RequestString(method string, u string, body string, opts ...Options) (string, error)
```

搜索数组中是否存在指定的值

```
slices.InArray[T comparable](item T, slice []T) bool
```

## License

Apache License Version 2.0, http://www.apache.org/licenses/
