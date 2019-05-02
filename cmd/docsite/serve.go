package main

import (
	"flag"
	"net"
)

func init() {
	flagSet := flag.NewFlagSet("serve", flag.ExitOnError)
	var (
		httpAddr    = flagSet.String("http", ":5080", "HTTP listen address for previewing")
		tlsCertPath = flagSet.String("tls-cert", "", "path to TLS certificate file")
		tlsKeyPath  = flagSet.String("tls-key", "", "path to TLS key file")
	)

	handler := func(args []string) error {
		flagSet.Parse(args)

		host, port, err := net.SplitHostPort(*httpAddr)
		if err != nil {
			return err
		}

		if host == "" {
			host = "0.0.0.0"
		}

		site, _, err := siteFromFlags()
	}
}
