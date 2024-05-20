// maps3
// Make me compile!
//
package main

import "testing"

func TestGetPhone(t *testing.T) {
	phoneBook := map[string]string{
		"Ana":  "+01 101 102",
		"John": "+01 333 666",
	}

	phone := phoneBook["Ana"] // something seems wrong here
	expectedPhone := "+01 101 102"
	if phone != expectedPhone {
		t.Errorf("phone should be %s but got %s", expectedPhone, phone)
	}
}

func TestInsertPhone(t *testing.T) {
	phoneBook := map[string]string{
		"Ana":  "+01 101 102",
		"John": "+01 333 666",
	}

	phone, _ := phoneBook["Laura"] // don't change this line
	var expectedPhone string 
	if phone != expectedPhone {
		t.Errorf("phone should be %s but got %s", expectedPhone, phone)
	}
}

func TestDeletePhone(t *testing.T) {
	phoneBook := map[string]string{
		"Ana":  "+01 101 102",
		"John": "+01 333 666",
	} // don't change the original map

	delete(phoneBook, "John")
	totalPhones := len(phoneBook)
	expectedTotalPhones := 1
	if totalPhones != expectedTotalPhones {
		t.Errorf("phoneBook should have only %d contact, but got %d", expectedTotalPhones, totalPhones)
	}
}
