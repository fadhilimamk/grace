package grace

import (
  "net/http"
  "log"
)

func ExampleMainHandler() {
  http.HandleFunc("/foo/bar", foobarHandler)
  log.Fatal(Serve(":9000", nil))
}

func foobarHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("foobar"))
}
