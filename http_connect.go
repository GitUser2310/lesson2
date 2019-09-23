package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	tmp1 := template.New("main")
	tmp1, _ = tmp1.Parse(`
				<div style="display: inline-block; background-color: orange; border: 1px solid #aaa; border-radius: 3px; padding: 30px; margin: 20px;">
				{{if ne . "str"}}
				  Hello, friends!!!
				{{end}}
				<pre>{{.}}</pre>
				</div>`)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		c := http.Client{}

		resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		// return body in pattern
		tmp1.Execute(w, string(body))
	})

	http.ListenAndServe(":8081", nil)
}
