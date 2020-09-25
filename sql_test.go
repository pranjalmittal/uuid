// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	stringTest := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	badTypeTest := 6
	invalidTest := "f47ac10b-58cc-0372-8567-0e02b2c3d4"

	byteTest := make([]byte, 16)
	byteTestUUID := Must(Parse(stringTest))
	copy(byteTest, byteTestUUID[:])

	// sunny day tests

	var uuid UUID
	err := (&uuid).Scan(stringTest)
	if err != nil {
		t.Fatal(err)
	}

	err = (&uuid).Scan([]byte(stringTest))
	if err != nil {
		t.Fatal(err)
	}

	err = (&uuid).Scan(byteTest)
	if err != nil {
		t.Fatal(err)
	}

	// bad type tests

	err = (&uuid).Scan(badTypeTest)
	if err == nil {
		t.Error("int correctly parsed and shouldn't have")
	}
	if !strings.Contains(err.Error(), "unable to scan type") {
		t.Error("attempting to parse an int returned an incorrect error message")
	}

	// invalid/incomplete uuids

	err = (&uuid).Scan(invalidTest)
	if err == nil {
		t.Error("invalid uuid was parsed without error")
	}
	if !strings.Contains(err.Error(), "invalid UUID") {
		t.Error("attempting to parse an invalid UUID returned an incorrect error message")
	}

	err = (&uuid).Scan(byteTest[:len(byteTest)-2])
	if err == nil {
		t.Error("invalid byte uuid was parsed without error")
	}
	if !strings.Contains(err.Error(), "invalid UUID") {
		t.Error("attempting to parse an invalid byte UUID returned an incorrect error message")
	}

	// empty tests

	uuid = UUID{}
	var emptySlice []byte
	err = (&uuid).Scan(emptySlice)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning empty byte slice")
		}
	}

	uuid = UUID{}
	var emptyString string
	err = (&uuid).Scan(emptyString)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning empty byte slice")
		}
	}

	uuid = UUID{}
	err = (&uuid).Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning nil")
		}
	}
}

func TestValue(t *testing.T) {
	stringTest := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	uuid := Must(Parse(stringTest))
	val, _ := uuid.Value()
	if val != stringTest {
		t.Error("Value() did not return expected string")
	}
}

func TestNullUUIDValueNil(t *testing.T) {
	u := NullUUID{}

	val, _ := u.Value()
	if val != nil {
		t.Error("Value() did not return expected value")
	}
}

func TestNullUUIDScanValid(t *testing.T) {
	u := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
	s1 := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"

	u1 := NullUUID{}
	err := u1.Scan(s1)
	if err != nil {
		t.Error("unable to scan null uuid")
	}
	if u1.Valid != true {
		t.Error("null uuid not valid")
	}
	if u1.UUID != u {
		t.Error("uuid are not same")
	}
}

func TestNullUUIDScanNil(t *testing.T) {
	u := NullUUID{UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}, true}

	err := u.Scan(nil)
	if err != nil {
		t.Error("unable to scan null uuid")
	}
	if u.Valid != false {
		t.Error("null uuid not valid")
	}
	if u.UUID != Nil {
		t.Error("")
	}
}
