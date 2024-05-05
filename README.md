<p align="center">
<h1 align="center">Grest</h1>
<p align="center">basically is <a href="https://github.com/go-resty/resty">Resty</a> but bad</p>

## Instalation

```bash
go get github.com/BrunProgramming/grest
```

## Usage

import the necessary modules

```go
import (
    "fmt"
    "os"
    "github.com/BrunProgramming/grest"
)
```

Now the code for the GET request

```go
c := grest.New()

// this is using jsonplaceholder for the example
// for more information enter in https://jsonplaceholder.typicode.com
resp, err := c.Get("https://jsonplaceholder.typicode.com/posts/1")

if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}

fmt.Println(resp)
```

in the console look like this
```bash
go run .

{
  "userId": 1,
  "id": 1,
  "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
}
```

Now the example for the POST request

```go
c := grest.New()

json := grest.Json{
	"title":  "foo",
	"body":   "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
	"userId": 1,
}

// For this example we need to define a header for the correct functioning of jsonplaceholder but this can change depending on what the specific parameters establish
c.SetHeaders(grest.Headers{
    "Content-type": "application/json; charset=UTF-8",
})

resp, err := c.Post("https://jsonplaceholder.typicode.com/posts", c.Json(json))

if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}

fmt.Println(resp)
```

Now this example looks like this in the console

```bash
go run .

{
  "body": "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  "title": "foo",
  "userId": 1,
  "id": 101
}
```
