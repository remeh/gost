package gost

import (
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/remeh/go-webserver"
//  "code.google.com/p/go-uuid/uuid"
)

const (
    CONTROLLER_HTTP_PORT = 9100
)

// Basic controller which uses the HTTP protocol
// to receive task to broadcast.
type HttpController struct {
    httpServer  webserver.App   // Web listener
    gost        Gost            // Gost runtime
}

func (c *HttpController) Start() {
    c.httpServer.Init()

    // A post action to receive the task to execute.
    c.httpServer.Router.Add("Http controller action", "POST", &HttpControllerAction{c.gost}, "/:topic/:action/:tid")

    // Let's listen for HTTP call in background.
    fmt.Printf("[gost] [HttpController] Starts on port %d\n", CONTROLLER_HTTP_PORT)
    go c.httpServer.Start(CONTROLLER_HTTP_PORT)
}

func (c *HttpController) Close() {
}

// The action which receive the HTTP call to broadcast
// a task.
type HttpControllerAction struct {
    gost        Gost
}

func (a *HttpControllerAction) Init() {
}

func (a *HttpControllerAction) Execute(writer http.ResponseWriter, request *http.Request, parameters map[string]string) (int, string) {
    tid := parameters["tid"]
    if len(tid) == 0 {
        return 500, "" // TODO json error response
    }

    topic := parameters["topic"]
    if len(topic) == 0 {
        return 500, "" // TODO json error response
    }

    action := parameters["action"]
    if len(topic) == 0 {
        return 500, "" // TODO json error response
    }

    // read the body
    body, err := ioutil.ReadAll(request.Body)
    if err != nil {
        return 500, "" // TODO json error response
    }

    task    := NewSimpleTask(tid, topic, action, body)

    // Broadcast the task to the worker.
    a.gost.GetBroadcaster().Broadcast(task)

    return 200, ""
}
