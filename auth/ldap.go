package auth

import (
	"gopkg.in/ldap.v2"
	pool "gopkg.in/fatih/pool.v2"
	"crypto/tls"
	"fmt"
	"log"
)

const (
	ldapUrl = "ldaps-vip.cc.ic.ac.uk"
	ldapPort = 636
	docsocDL = "CN=zz-icu-docsoc-members-dl,OU=Distribution,OU=Groups,OU=Imperial College (London),DC=ic,DC=ac,DC=uk"
)

const (
	ldapUsernameAttribute = "sAMAccountName"
	ldapFirstNameAttribute = "givenName"
	ldapSurnameAttribute = "sn"
	ldapMemberOf = "memberOf"
	ldapDomainComponent = "dn"
	ldapCommonName = "cn"
)


func ldapsConnection() *ldap.Conn {
	// TLS, for testing purposes disable certificate verification, check https://golang.org/pkg/crypto/tls/#Config for further information.
	tlsConfig := &tls.Config{ServerName: ldapUrl}
	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapUrl, ldapPort), tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	return l
}

func search(l *ldap.Conn, accountName string) []*ldap.Entry {
	searchRequest := ldap.NewSearchRequest(
		"dc=ic,dc=ac,dc=uk", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(" + ldapUsernameAttribute +"=" + accountName + ")", // The filter to apply
		[]string{ldapDomainComponent, ldapCommonName, ldapFirstNameAttribute, ldapSurnameAttribute},                    // A list attributes to retrieve
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

	firstName := entries[0].GetAttributeValue(ldapFirstNameAttribute)
	surname := entries[0].GetAttributeValue(ldapSurnameAttribute)

	return firstName +" "+ surname
}

func isDoCSoc(l *ldap.Conn, accountName string) bool {
	entries := search(l, accountName)

	if len(entries) == 0 {
		return false
	}

	return contains(entries[0].GetAttributeValues(ldapMemberOf), docsocDL)
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
