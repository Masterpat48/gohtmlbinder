## DOCUMENTATION
---------------
**gohtmlbinder** is a lightweight Go package that makes it easy to bind HTML templates to HTTP routes using [Gorilla Mux](https://github.com/gorilla/mux).  
It provides a simple and reusable way to serve web pages without having to manually handle `html/template`, routes, and servers in every project.
---------------
# How to use
- Install the module with:
```bash
  go get github.com/Masterpat48/gohtmlbinder
```

- Initialize the binder with:
```go
  b := binder.New("file.html")
```
  
- Make a route with static data:
```go
  b.NewRoute("/route", "file.html")
```

- Make a route with dynamic data:
```go
  b.NewRouteData("/route", "file.hmtl", func(r *http.Request) any{dynamic data}
```

- Print all the registred routes:
```go
	b.PrintRoutes()
```

- Start the server:
```go
  b.serve(":port")
```
--------------
# Example
```go
package main

import (
	"github.com/Masterpat48/gohtmlbinder"
)

func main() {
	//initialize the binder W the index file (need to have the index file created)
	b := binder.New("index.html")
	//create the main route
	b.NewRoute("/", "index.html")
	//print in console all the registred routes
	b.PrintRoutes()
	//start the server on port 1000
	b.Serve(":1000")
}
```
