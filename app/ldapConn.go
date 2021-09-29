package app

import (
	"crypto/tls"

	"gopkg.in/ldap.v2"
)


//ldapServer := "test.lab:636"

func NewLdapConn () (*ldap.Conn, error) {

	ldapServer := "test.lab:636"
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	conn, err := ldap.DialTLS("tcp", ldapServer, tlsConfig) //conn

	if err != nil {
		return nil,err
	}

	err = conn.Bind(`test\admin`, "Te$t0vP@$$") //conn log& pass

	if err != nil {
		return nil, err
	}
	return conn, nil
}