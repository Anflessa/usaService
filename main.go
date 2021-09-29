package main

import (
	"LDAPapi/app"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)


func main() {

	//u16 :=utf16.Encode([]rune("encodedsecret"))

	b := []byte{
		'T',
		0x00,
		'E',
		0x00,
		'S',
		0x00,
		'T',
		0x00,
		'T',
		0x00,
		'E',
		0x00,
		'S',
		0x00,
		'T',
		0x00,
		'\n',
	}


	addRequest := ldap.AddRequest{ //главная матрешка
		//DN:         "CN=golangTest1,CN=Users,DC=test,DC=lab", //путь
		DN:         "CN=lesov2,OU=Тестовые пользователи,DC=test,DC=lab", //путь
		Attributes: []ldap.Attribute{},
	}

	fmt.Println(string(b))

	//addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "cn",Vals: []string{"lesov"}})
	//addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "givenname",Vals: []string{"lesov"}})
	//addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "sn",Vals: []string{"lesov"}})
	//addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "uid",Vals: []string{"lesov"}})
	//addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "mail",Vals: []string{"testlesov@post.ru"}})
	//addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "telephoneNumber",Vals: []string{"2223"}})
	addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "unicodePwd",Vals: []string{"123 109 100 53 125 66 55 105 50 70 112 88 114 121 113 116 57 108 116 80 48 113 116 119 66 102 103 61 61"}})
	//addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "sAMAccountName",Vals: []string{"lesov"}})
	//addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "userPrincipalName",Vals: []string{"lesov@test.lab"}})
	//addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "UserAccountControl",Vals: []string{"512"}})
	addRequest.Attributes = append(addRequest.Attributes, ldap.Attribute{Type: "objectclass",Vals: []string{"top","person","organizationalPerson","user"}})



	//a:=ldap.Attribute{Type: attrType,Vals:valsArr} //создание элемент атрибута(одна из структур атрибутес)

	conn, err := app.NewLdapConn() //Установили соединение
	if err!=nil{
		fmt.Println("when connect:",err)
	}

	err = conn.Add(&addRequest) //Push request to LDAP (try create user)
	if err!=nil{
		fmt.Println("when push request:",err)
	}
	//conn.Start() //Maybe starting

}

//<?
//$ldaprdn = 'negreev.r@test.lab'; // ldap rdn или dn
//$ldappass = 'Zz123456';
//$ldaptree = "OU=Тестовые пользователи,DC=test,DC=lab";
//$ldaptree_add ="CN=test_test,OU=Тестовые пользователи,DC=test,DC=lab";
//
//
//$newPassword = "DfGhJk9";
//
//$newPass = iconv( 'UTF-8', 'UTF-16LE', '"' . $newPassword . '"');
//
///*
//for($i=0;$i<$len;$i++) {
//    echo $newPassword[$i];
//    $newPassw .= "{$newPassword[$i]}\000";
//}*/
//
//$ldaprecord['cn'] = "test_test";
//$ldaprecord['givenname'] = "test3";
//$ldaprecord['sn'] = "test2_test2";
//$ldaprecord["uid"]="test_test3";
//$ldaprecord['mail'] = "test@post.ru";
//$ldaprecord['telephoneNumber'] = "2222";
//$ldaprecord["unicodePwd"] = $newPass;
//$ldaprecord["sAMAccountName"] = "test_test4";
//$ldaprecord["userprincipalname"]="test_test4@test.lab";
//$ldaprecord["UserAccountControl"] = 512;
//$ldaprecord['objectclass'][0] = "top";
//$ldaprecord['objectclass'][1] = "person";
//$ldaprecord['objectclass'][2] = "organizationalPerson";
//$ldaprecord['objectclass'][3] = "user";
//
//
//$ldapconn = ldap_connect("ldaps://DC-TEST",636);
//
//
//ldap_set_option($ldapconn, LDAP_OPT_PROTOCOL_VERSION, 3) or die('Unable to set LDAP protocol version');
//ldap_set_option($ldapconn, LDAP_OPT_REFERRALS, 0); // We need this for doing an LDAP search.
//
//if ($ldapconn) {
//
//
//$ldapbind = ldap_bind($ldapconn, $ldaprdn, $ldappass);
//
//if ($ldapbind) {
//
//echo $ldapbind;
//$r = ldap_add($ldapconn,$ldaptree_add, $ldaprecord);
//
////добавим в группу
//$dn = "CN=тестовая1,OU=Тестовые группы,DC=test,DC=lab"; // distinguished name/DN of the group you want to add to
//$info["member"] = "CN=test_test,OU=тестовые пользователи,DC=test,DC=lab"; // DN of the user you want to add
//ldap_modify($ldapconn, $dn, $info);
//// $result = ldap_list($ldapconn,$ldaptree, "(objectClass=user)");
//// $data = ldap_get_entries($ldapconn, $result);
//// $r = ldap_add($ds, $dn, $ldaprecord);
//// SHOW ALL DATA
//echo '<h1>Dump all data</h1><pre>';
//// print_r($data);
//echo '</pre>';
//} else {
//echo " привязка LDAP не удалась...";
//}
//
//}
