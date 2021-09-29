package app

import (
	"crypto/tls"
	"github.com/go-ldap/ldap/v3"
)


//ldapServer := "test.lab:636"

func NewLdapConn () (*ldap.Conn, error) {

	ldapServer := "test.lab:636"
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	conn, err := ldap.DialTLS("tcp", ldapServer, tlsConfig) //conn

	if err != nil {
		return nil,err
	}

	err = conn.Bind(`admin`, "Te$t0vP@$$") //conn log& pass
	if err != nil {
		return nil, err
	}



	return conn, nil
}


