package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

const (
	gitHubStarUrl = "https://api.github.com/users/%s/starred"
)

func doStarChart(_ *log.Logger, r render.Render) {

	res, err := http.Get(fmt.Sprintf(gitHubStarUrl, "kyokomi"))
	if err != nil {
		r.Error(400)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		r.Error(400)
	}

	var stars []Starred
	if err := json.Unmarshal(data, &stars); err != nil {
		r.Error(400)
	}

	r.HTML(200, "index", nil)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl"},
	}))

	// Router
	m.Get("/", doStarChart)

	m.Run()
}
