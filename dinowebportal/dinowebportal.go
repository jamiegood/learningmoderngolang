package dinowebportal

import (
	"fmt"
	"net/http"
)

// RunWebPortal ...
func RunWebPortal(addr string) error {
	http.HandleFunc("/", routeHandler)
	//err := http.ListenAndServe(addr, nil)
	return http.ListenAndServe(addr, nil)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "weclome to the portal %s", r.RemoteAddr)
}
