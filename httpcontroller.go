package gooch

import (
    "bytes"
    "fmt"
    "net/http"

    "github.com/remeh/go-webserver"
    "code.google.com/p/go-uuid/"
)

const (
    CONTROLLER_HTTP_PORT = 9100
)

// Basic controller which uses the HTTP protocol
// to receive task to broadcast.
type HttpController struct {
    httpServer webserver.App // Web listener
}

func (c *HttpController) Start() {
    c.httpServer.Init()

    // A post action to receive the task to execute.
    c.httpServer.Router.Add("Http controller action", "POST", &HttpControllerAction{}, "/:id")

    // Let's listen for HTTP call in background.
    fmt.Printf("[gooch] [HttpController] Starts on port %d\n", CONTROLLER_HTTP_PORT)
    go c.httpServer.Start(CONTROLLER_HTTP_PORT)
}

// The action which receive the HTTP call to broadcast
// a task.
type HttpControllerAction struct {
}

func (a *HttpControllerAction) Init() {
}

func (a *HttpControllerAction) Execute(writer http.ResponseWriter, request *http.Request, parameters map[string]string) (int, string) {
    tid := parameters["tid"]
    if len(tid) == 0 {
        return 500, "No task id provided." // TODO json error response
    }

    // TODO broadcast the task to the workers.
    id      := uuid.New()
    task    := NewSimpleTask("http", "echo", bytes.NewBufferString("Content of the task").Bytes())

    return 200, ""
}
