// Go to github.com/coreos/etcd/hack/tls-setup and launch etcd cluster
// Test with curl
// $ curl --cacert certs/ca.pem --key certs/proxy1-key.pem --cert certs/proxy1.pem  https://localhost:8080/v2/keys

// Test client:
// $ go run etcdclient.go -etcd localhost:8080 -key $GOPATH/src/github.com/coreos/etcd/hack/tls-setup/certs/proxy1-key.pem -cert $GOPATH/src/github.com/coreos/etcd/hack/tls-setup/certs/proxy1.pem -cacert $GOPATH/src/github.com/coreos/etcd/hack/tls-setup/certs/ca.pem

package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"time"

	"log"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/etcd"
)

func main() {
	// Flags
	var (
		etcdAddr   = flag.String("etcd", "", "etcd service address (can be proxy).")
		certFile   = flag.String("cert", "", "Client certificate file.")
		keyFile    = flag.String("key", "", "Private key file name.")
		caCertFile = flag.String("cacert", "", "CA certificate file name.")
	)
	flag.Parse()

	if *etcdAddr == "" {
		log.Fatal("etcd address is required")
	}

	// Read certificates
	var tlsC *tls.Config = nil
	if *certFile != "" && *keyFile != "" {
		cc := &store.ClientTLSConfig{
			CertFile:   *certFile,
			KeyFile:    *keyFile,
			CACertFile: *caCertFile,
		}

		var rpool *x509.CertPool = nil
		if *caCertFile != "" {
			if pemBytes, err := ioutil.ReadFile(cc.CACertFile); err == nil {
				rpool = x509.NewCertPool()
				rpool.AppendCertsFromPEM(pemBytes)
			} else {
				log.Printf("Error reading etcd cert CA file. Err: %s", err)
			}
		}

		if tlsCert, err := tls.LoadX509KeyPair(cc.CertFile, cc.KeyFile); err == nil {
			tlsC = &tls.Config{
				RootCAs:      rpool,
				Certificates: []tls.Certificate{tlsCert},
			}
		} else {
			log.Printf("Error loading keypair for TLS client. Err: %s", err)
		}
	}

	// Construct store
	etcd.Register()
	kv, err := libkv.NewStore(
		store.ETCD,
		[]string{*etcdAddr},
		&store.Config{
			TLS:               tlsC,
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		log.Fatalf("Cannot create store. Err: %s", err)
	}

	// Basic test
	key := "foo"
	err = kv.Put(key, []byte("bar"), nil)
	if err != nil {
		log.Printf("Error trying to put value at key: %s. Err: %s", key, err)
	}

	pair, err := kv.Get(key)
	if err != nil {
		log.Printf("Error trying accessing value at key: %s. Err: %s", key, err)
	}

	err = kv.Delete(key)
	if err != nil {
		log.Printf("Error trying to delete key %s. Err: %s", key, key)
	}

	if pair != nil {
		log.Printf("Tested value: %s", pair.Value)
	}
}