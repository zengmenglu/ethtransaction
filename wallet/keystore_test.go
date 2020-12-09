package wallet

import "testing"

func TestCreatKs(t *testing.T){
		CreateKs("secret")
}

func TestImport(t *testing.T){
	ks:= CreateKs("secret")
	importKs(ks,"secret")
}