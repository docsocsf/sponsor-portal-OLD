package auth

import (
	"gopkg.in/ldap.v2"
	"crypto/tls"
	"fmt"
	"log"
	"errors"
	"github.com/davecgh/go-spew/spew"
)

const (
	ldapUrl = "ldaps-vip.cc.ic.ac.uk"
	ldapPort = 636
	ldapConnPoolSize = 10
	docsocDL = "CN=zz-icu-docsoc-members-dl,OU=Distribution,OU=Groups,OU=Imperial College (London),DC=ic,DC=ac,DC=uk"
	domain = "@ic.ac.uk"
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

func (wrapper *LDAPWrapper) bind(l *ldap.Conn) error {
	err := l.Bind(wrapper.username, wrapper.password)
	if err != nil {
		return err
	}
	return nil
}

func ldapsConnection() (interface{}, error) {
	// TLS, check https://golang.org/pkg/crypto/tls/#Config for further information.
	tlsConfig := &tls.Config{ServerName: ldapUrl}
	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapUrl, ldapPort), tlsConfig)
	if err != nil {
		return l, err
	}

	return l, nil
}

func (wrapper *LDAPWrapper) search(accountName string) ([]*ldap.Entry, error) {
	l := pool.GetConnection().(*ldap.Conn)
	wrapper.bind(l)

	searchRequest := ldap.NewSearchRequest(
		ldapBaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(" + ldapUsernameAttribute +"=" + accountName + ")",
		[]string{ldapDomainComponent, ldapCommonName, ldapFirstNameAttribute, ldapSurnameAttribute, ldapMemberOf},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return []*ldap.Entry{}, err
	}

	pool.ReleaseConnection(l)

	return sr.Entries, nil
}

func (wrapper *LDAPWrapper) searchForName(accountName string) (string, error) {
	entries, err := wrapper.search(accountName)

	if err != nil {
		return "", err
	}

	firstName := entries[0].GetAttributeValue(ldapFirstNameAttribute)
	surname := entries[0].GetAttributeValue(ldapSurnameAttribute)

	return firstName +" "+ surname, err
}

func (wrapper *LDAPWrapper) isDoCSoc(accountName string) (bool, error) {
	entries, err := wrapper.search(accountName)

	spew.Dump(entries)

	if len(entries) == 0 {
		return false, err
	}

	return contains(entries[0].GetAttributeValues(ldapMemberOf), docsocDL), nil
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

func (wrapper *LDAPWrapper) userAuth(username string, password string) (bool, error) {
	l := pool.GetConnection().(*ldap.Conn)
	wrapper.bind(l)

	log.Println("username: " + username)

	if ok, _ := wrapper.isDoCSoc(username); !ok {
		return false, errors.New("user is not a member of DoCSoc")
	}

	searchResult, err := wrapper.search(username)

	if len(searchResult) != 1 {
		return false, errors.New("user does not exist or too many entries returned")
	}

	// Bind as the user to verify their password
	err = l.Bind(username + domain, password)
	if err != nil {
		return false, err
	}

	//// Rebind as the service user for any further queries
	//err = l.Bind(serviceUsername, servicePassword)
	//if err != nil {
	//	return false, err
	//}

	pool.ReleaseConnection(l)

	return true, nil
}
