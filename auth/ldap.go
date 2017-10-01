package auth

import (
	"gopkg.in/ldap.v2"
	"crypto/tls"
	"fmt"
	"log"
)

const (
	ldapUrl = "ldaps-vip.cc.ic.ac.uk"
	ldapPort = 636
	ldapConnPoolSize = 10
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

type InitFunction func() (interface{}, error)

type ConnectionPoolWrapper struct {
	size int
	conn chan interface{}
}

var pool = &ConnectionPoolWrapper{}

func init() {
	err := pool.InitPool(ldapConnPoolSize, ldapsConnection)
	if err != nil {
		log.Fatal("LDAP connection pool couldn't be created")
	}
}

/**
 Call the init function size times. If the init function fails during any call, then
 the creation of the pool is considered a failure.
 We call the same function size times to make sure each connection shares the same
 state.
*/
func (p *ConnectionPoolWrapper) InitPool(size int, initfn InitFunction) error {
	// Create a buffered channel allowing size senders
	p.conn = make(chan interface{}, size)
	for x := 0; x < size; x++ {
		conn, err := initfn()
		if err != nil {
			return err
		}

		// If the init function succeeded, add the connection to the channel
		p.conn <- conn
	}
	p.size = size
	return nil
}

func (p *ConnectionPoolWrapper) GetConnection() interface{} {
	return <-p.conn
}

func (p *ConnectionPoolWrapper) ReleaseConnection(conn interface{}) {
	p.conn <- conn
}

type LDAPWrapper struct {
	pool ConnectionPoolWrapper
	username string
	password string
}

func (wrapper *LDAPWrapper) bind(l *ldap.Conn) {
	err := l.Bind(wrapper.username, wrapper.password)
	if err != nil {
		log.Fatal(err)
	}
}

func ldapsConnection() (interface{}, error) {
	// TLS, check https://golang.org/pkg/crypto/tls/#Config for further information.
	tlsConfig := &tls.Config{ServerName: ldapUrl}
	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapUrl, ldapPort), tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	return l, nil
}

func (wrapper *LDAPWrapper) search(accountName string) []*ldap.Entry {
	l := pool.GetConnection().(*ldap.Conn)
	wrapper.bind(l)

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

	pool.ReleaseConnection(l)

	return sr.Entries
}

func (wrapper *LDAPWrapper) searchForName(accountName string) string {
	entries := wrapper.search(accountName)

	firstName := entries[0].GetAttributeValue(ldapFirstNameAttribute)
	surname := entries[0].GetAttributeValue(ldapSurnameAttribute)

	return firstName +" "+ surname
}

func (wrapper *LDAPWrapper) isDoCSoc(accountName string) bool {
	entries := wrapper.search(accountName)

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
func (wrapper *LDAPWrapper) userAuth(serviceUsername string, servicePassword string, username string, password string) bool {
	l := pool.GetConnection().(*ldap.Conn)
	wrapper.bind(l)

	// First bind with our service user
	err := l.Bind(serviceUsername, servicePassword)
	if err != nil {
		log.Fatal(err)
	}

	if !wrapper.isDoCSoc(username) {
		log.Println("User is not a member of DoCSoc")
	}

	searchResult := wrapper.search(username)

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
	pool.ReleaseConnection(l)

	return true
}
