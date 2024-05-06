package main

import (
  "github.com/charmbracelet/lipgloss"
  "fmt"
  "testing"
)

const want_get_response = `{
  "userId": 1,
  "id": 1,
  "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
}`

const want_post_response = `{
  "body": "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  "title": "foo",
  "userId": 1,
  "id": 101
}`

var (
  style = lipgloss.NewStyle().
  Foreground(lipgloss.Color("#0f0"))

  c = New()
)

func TestGet(t *testing.T) {
  // this is using jsonplaceholder for the example
  // for more information enter in https://jsonplaceholder.typicode.com
  resp, err := c.Get("https://jsonplaceholder.typicode.com/posts/1")

  if err != nil {
    t.Error(err)
  }

  if resp != want_get_response {
    t.Errorf("got:%q,wanted:%q",resp,want_get_response)
  } else {
    fmt.Println(style.Render("Get test passed successfully ✓"))
  }

}

func TestPost(t *testing.T) {
  json := Json{
	  "title":  "foo",
	  "body":   "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
	  "userId": 1,
  }

  // For this example we need to define a header for the correct functioning of jsonplaceholder but this can change depending on what the specific parameters establish
  c.SetHeaders(Headers{
    "Content-type": "application/json; charset=UTF-8",
  }) 

  resp, err := c.Post("https://jsonplaceholder.typicode.com/posts", c.Json(json))

  if err != nil {
    t.Error(err)
  }

  if resp != want_post_response {
    t.Errorf("got:%q, wanted:%q",resp, want_post_response)
  } else {
    fmt.Println(style.Render("\nPost test passed ✓"))
  }
}
