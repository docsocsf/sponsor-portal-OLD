package auth

import (
	"gopkg.in/ldap.v2"
	"crypto/tls"
	"fmt"
	"log"
)

func ldapsConnection() *ldap.Conn {
	// TLS, for testing purposes disable certificate verification, check https://golang.org/pkg/crypto/tls/#Config for further information.
	tlsConfig := &tls.Config{ServerName: "ldaps-vip.cc.ic.ac.uk"}
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
		[]string{"dn", "cn", "givenName", "sn"},                    // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	return sr.Entries
}

func searchForName(l *ldap.Conn, accountName string) string {
	entries := search(l, accountName)

	firstName := entries[0].GetAttributeValue("givenName")
	surname := entries[0].GetAttributeValue("sn")

	return firstName +" "+ surname
}

func isDoCSoc(l *ldap.Conn, accountName string) bool {
	entries := search(l, accountName)

	return contains(entries[0].GetAttributeValues("memberOf"), "CN=zz-icu-docsoc-members-dl,OU=Distribution,OU=Groups,OU=Imperial College (London),DC=ic,DC=ac,DC=uk")
}

// From: https://stackoverflow.com/a/27272103
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// Example User Authentication shows how a typical application can verify a login attempt
func userAuth(l *ldap.Conn, serviceUsername string, servicePassword string, username string, password string) bool {
	// First bind with our service user
	err := l.Bind(serviceUsername, servicePassword)
	if err != nil {
		log.Fatal(err)
	}

	if !isDoCSoc(l, username) {
		log.Println("User is not a member of DoCSoc")
	}

	searchResult := search(l, username)
	meh := searchForName(l, "jep114")
	log.Println(meh)

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
