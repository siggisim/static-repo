package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

var (
	port      = flag.Int("port", 8080, "The port to listen on")
	staticDirectory = flag.String("static_directory", "./static", "the directory containing static files to host")
)

func init() {
	// Parse all flags.
	flag.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>\n")
	fmt.Fprint(w, "<p>Hello world!</p>")
	fmt.Fprint(w, "Check out the static content over at <a href=\"/static\">/static</a>\n")
	fmt.Fprint(w, "</body></html>\n")
}

func main() {
	// Figure out where our runfiles (static content bundled with the binary) live.
	rfp, err := bazel.RunfilesPath()
	if err != nil {
		log.Fatalf("Error determining runfiles path: %v", err)
	}

	// Handle "/static/*" requests by serving those static files out of the bundled runfiles.
	pkgStaticDir := filepath.Join(rfp, *staticDirectory)
	fs := http.FileServer(http.Dir(pkgStaticDir))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// Handle "/" with our custom handler above.
	http.HandleFunc("/", handler)

	hostAndPort := fmt.Sprintf("0.0.0.0:%d", *port)
	log.Printf("Listening on http://%s\n", hostAndPort)
	log.Fatal(http.ListenAndServe(hostAndPort, nil))
}
