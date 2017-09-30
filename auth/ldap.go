package auth

import (
	"gopkg.in/ldap.v2"
	"crypto/tls"
	"fmt"
	"log"
)

func ldapsConnection() *ldap.Conn {
	// TODO: verify TLS
	// TLS, for testing purposes disable certificate verification, check https://golang.org/pkg/crypto/tls/#Config for further information.
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", "ldaps-vip.cc.ic.ac.uk", 636), tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	return l
}

func search(l *ldap.Conn, accountName string) []*ldap.Entry {
	searchRequest := ldap.NewSearchRequest(
		"dc=ic,dc=ac,dc=uk", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(sAMAccountName=" + accountName + ")", // The filter to apply
		[]string{"dn", "cn"},                    // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	return sr.Entries
}

// Example User Authentication shows how a typical application can verify a login attempt
func userAuth(l *ldap.Conn, serviceUsername string, servicePassword string, username string, password string) bool {
	// First bind with our service user
	err := l.Bind(serviceUsername, servicePassword)
	if err != nil {
		log.Fatal(err)
	}

	searchResult := search(l, username)

	if len(searchResult) != 1 {
		fmt.Println("User does not exist or too many entries returned")
		return false
	}

	// Bind as the user to verify their password
	err = l.Bind(username + "@IC.AC.UK", password)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Rebind as the service user for any further queries
	err = l.Bind(serviceUsername, servicePassword)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
