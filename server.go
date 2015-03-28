package boxcars

import (
	"fmt"
	"net/http"
)

func Listen(port int) {
	http.HandleFunc("/", OnRequest)
	exit := make(chan error)

	go func() {
		debug("Starting at %d", port)
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			debug("Fatal: %v", err)
		}
		exit <- err
	}()

	go func() {
		debug("Starting TLS at %d", port + 1)
		err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port+1), "cert.pem", "key.pem", nil)
		if err != nil {
		}
		exit <- err
	}()
	for {
		select {
		case err := <-exit:
			debug("Fatal: %v", err)
			return
		}
	}
}
