package auth

import (
	"crypto/tls"
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

const (
	ldapUrl          = "ldaps-vip.cc.ic.ac.uk"
	ldapPort         = 636
	ldapConnPoolSize = 10
	docsocDL         = "CN=zz-icu-docsoc-members-dl,OU=Distribution,OU=Groups,OU=Imperial College (London),DC=ic,DC=ac,DC=uk"
	domain           = "@ic.ac.uk"
)

const (
	ldapUsernameAttribute  = "sAMAccountName"
	ldapFirstNameAttribute = "givenName"
	ldapSurnameAttribute   = "sn"
	ldapMemberOf           = "memberOf"
	ldapDomainComponent    = "dn"
	ldapCommonName         = "cn"
	ldapBaseDN             = "dc=ic,dc=ac,dc=uk"
)

func newLDAPConn() (*ldap.Conn, error) {
	// TLS, check https://golang.org/pkg/crypto/tls/#Config for further information.
	tlsConfig := &tls.Config{ServerName: ldapUrl}
	return ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapUrl, ldapPort), tlsConfig)
}

func getAccount(accountName string, conn *ldap.Conn) (*ldap.Entry, error) {
	searchRequest := ldap.NewSearchRequest(
		ldapBaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"("+ldapUsernameAttribute+"="+accountName+")",
		[]string{ldapDomainComponent, ldapCommonName, ldapFirstNameAttribute, ldapSurnameAttribute, ldapMemberOf},
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	switch len(sr.Entries) {
	case 0:
		return nil, fmt.Errorf("No user (%s) found", accountName)
	case 1:
		return sr.Entries[0], nil
	default:
		return nil, fmt.Errorf("Username matched 2 or more accounts!")
	}
}

func getFullName(entry *ldap.Entry) string {
	firstName := entry.GetAttributeValue(ldapFirstNameAttribute)
	surname := entry.GetAttributeValue(ldapSurnameAttribute)

	return fmt.Sprintf("%s %s", firstName, surname)
}

func isDoCSoc(entry *ldap.Entry) (bool, error) {
	return contains(entry.GetAttributeValues(ldapMemberOf), docsocDL), nil
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

func userAuth(username string, password string) (UserInfo, error) {
	ui := UserInfo{}
	conn, err := newLDAPConn()

	// Bind as the user to verify their password
	email := username + domain
	err = conn.Bind(email, password)
	if err != nil {
		return ui, err
	}

	entry, err := getAccount(username, conn)
	if err != nil {
		return ui, err
	}

	ok, err := isDoCSoc(entry)
	if err != nil {
		return ui, err
	}

	if !ok {
		return ui, fmt.Errorf("Not DoCSoc")
	}

	ui.Name = getFullName(entry)
	ui.Email = email

	return ui, err
}
