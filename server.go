package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

var dockerApiVersion string

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port")
	flag.StringVar(&dockerApiVersion, "api-version", "v1.40", "Docker API version")
	flag.Parse()

	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/", serverNo)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "ok")
}

func serverNo(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	cmd := exec.Command(
		"/usr/bin/curl",
		"-s",
		"--unix-socket",
		"/var/run/docker.sock",
		"http:/"+dockerApiVersion+"/containers/"+path+"/json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "%s", "0")
		log.Printf("ERR: %s=%s", path, out.String())
		return
	}

	value := gjson.Get(out.String(), "Config.Labels.com\\.docker\\.compose\\.container-number")

	log.Printf("OK: %s=%s", path, value.String())
	fmt.Fprintf(w, "%s", value.String())
}
