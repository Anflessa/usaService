package main

import (
	"LDAPapi/app"
	"fmt"
	"gopkg.in/ldap.v2"
)



func main() {
dn:="CN=golangTest1,CN=Users,DC=test,DC=lab" //путь
var addRequest ldap.AddRequest
valsArr:=[]string{
	"golangTest1",
}
attrType:="sAMAccountName"

//a:=ldap.Attribute{Type: attrType,Vals:valsArr} //создание элемент атрибута(одна из структур атрибутес)

addRequest.DN = dn
addRequest.Attribute(attrType,valsArr)//при создании записи в ад кидаем тип атрибута и его значение

	conn, err := app.NewLdapConn()
	conn.Start()
	fmt.Println(conn,err)

	err1:=conn.Add(&addRequest)
	fmt.Println(err1)
}
