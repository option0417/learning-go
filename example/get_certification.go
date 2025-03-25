package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

const (
	//TargetHost = "coline.mitake.com.tw:443"
	TargetHost = "google.com:443"
)

func main() {
	// Establish a TLS connection. Replace "google.com:443" with your target host.
	conn, err := tls.Dial("tcp", TargetHost, nil)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Retrieve connection state, which contains certificate info.
	state := conn.ConnectionState()

	// Iterate through the peer certificates.
	for i, cert := range state.PeerCertificates {
		fmt.Printf("Certificate %d:\n", i)
		fmt.Printf("  Subject: %s\n", cert.Subject)
		fmt.Printf("  Issuer: %s\n", cert.Issuer)
		fmt.Printf("  Serial Number: %s\n", cert.SerialNumber)
		fmt.Printf("  Valid From: %s\n", cert.NotBefore)
		fmt.Printf("  Valid Until: %s\n", cert.NotAfter)
		fmt.Printf("  DNS Names: %v\n", cert.DNSNames)
		fmt.Printf("  Email Addresses: %v\n", cert.EmailAddresses)
		fmt.Printf("  IP Addresses: %v\n", cert.IPAddresses)
		fmt.Printf("  Signature Algorithm: %s\n", cert.SignatureAlgorithm)
		fmt.Printf("  Public Key Algorithm: %s\n", cert.PublicKeyAlgorithm)
		fmt.Println()
	}

	fmt.Printf("Size :%d\n", len(state.SignedCertificateTimestamps))
	for i, val := range state.SignedCertificateTimestamps {
		fmt.Printf("SignedCertificateTimestamps %d: %v\n", i, val)
	}

}
