package bar

type Client struct {
  User string
  Pass string
}

func bar() {
  client := Client{
    User: "helloworld"
    Pass: "2f2aafa267574197c7b7737f9fc93906452ca47525439c3f7ec1addfd27a842c"
  }

  log.Printf("Client credentials: %s/%s\n", client.User, client.Pass)
    
}
