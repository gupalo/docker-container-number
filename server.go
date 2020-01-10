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

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port")
	flag.Parse()

	http.HandleFunc("/healthcheck", healthcheck)
	//http.HandleFunc("/all", allServers)
	http.HandleFunc("/", serverNo)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "ok")
}

//func allServers(w http.ResponseWriter, r *http.Request) {
//	var out bytes.Buffer
//	cmd := exec.Command("/usr/bin/docker", "ps", "--format", "{{.ID}} {{.Label \"com.docker.compose.container-number\"}} {{.Names}}")
//	cmd.Stdout = &out
//	err := cmd.Run()
//	if err != nil {
//		fmt.Fprintf(w, "%s", "0")
//		return
//	}
//
//	fmt.Fprintf(w, "%s", out.String())
//}

func serverNo(w http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path[1:]
	var out bytes.Buffer

	//cmd := exec.Command("/usr/bin/docker", "inspect", r.URL.Path[1:])
	cmd := exec.Command("/usr/bin/curl", "-s", "--unix-socket", "/var/run/docker.sock", "http:/v1.40/containers/"+path+"/json")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "%s", "0")
		log.Printf("ERR: %s, %s", path, out.String())
		return
	}

	//value := gjson.Get(out.String(), "0.Config.Labels.com\\.docker\\.compose\\.container-number")
	value := gjson.Get(out.String(), "Config.Labels.com\\.docker\\.compose\\.container-number")

	log.Printf("OK: %s, %s", path, value.String())

	fmt.Fprintf(w, "%s", value.String())
}
