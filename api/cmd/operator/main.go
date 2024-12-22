package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/VincentHQL/scrctl/api/operator"
	apiv1 "github.com/VincentHQL/scrctl/api/operator/apiv1"
)

const (
	DefaultHttpPort      = 1080
	DefaultTLSCertDir    = "/etc/scrctl/cert"
	DefaultListenAddress = "127.0.0.1"
)

func startHttpServer(address string, port int) error {
	log.Println(fmt.Sprint("Operator is listening at http://localhost:", port))

	// handler is nil, so DefaultServeMux is used.
	return http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), nil)
}

func startHttpsServer(address string, port int, certPath string, keyPath string) error {
	log.Println(fmt.Sprint("Operator is listening at https://localhost:", port))
	return http.ListenAndServeTLS(fmt.Sprintf("%s:%d", address, port),
		certPath,
		keyPath,
		// handler is nil, so DefaultServeMux is used.
		//
		// Using DefaultServerMux in both servers (http and https) is not a problem
		// as http.ServeMux instances are thread safe.
		nil)
}

func fromEnvOrDefault(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func start(starters []func() error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(starters))
	for _, starter := range starters {
		go func(f func() error) {
			defer wg.Done()
			if err := f(); err != nil {
				log.Fatal(err)
			}
		}(starter)
	}
	wg.Wait()
}

func main() {
	httpPort := flag.Int("http_port", DefaultHttpPort, "Port to serve HTTP requests on.")
	httpsPort := flag.Int("https_port", -1, "Port to serve HTTPS requests on.")
	tlsCertDir := flag.String("tls_cert_dir", DefaultTLSCertDir, "Directory where the TLS certificates are located.")
	address := flag.String("listen_addr", DefaultListenAddress, "IP address to listen for requests.")

	flag.Parse()
	certPath := *tlsCertDir + "/cert.pem"
	keyPath := *tlsCertDir + "/key.pem"

	pool := operator.NewDevicePool()
	polledSet := operator.NewPolledSet()
	config := apiv1.InfraConfig{
		Type: "config",
		IceServers: []apiv1.IceServer{
			{URLs: []string{"stun:stun.l.google.com:19302"}},
		},
	}

	r := operator.CreateHttpHandlers(pool, polledSet, config)
	http.Handle("/", r)

	starters := []func() error{
		func() error { return startHttpServer(*address, *httpPort) },
	}

	if *httpsPort > 0 {
		starters = append(starters, func() error {
			return startHttpsServer(*address, *httpsPort, certPath, keyPath)
		})
	}

	start(starters)
}
