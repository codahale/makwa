package makwa_test

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"testing"

	"github.com/codahale/makwa"
)

func TestCheckPassword(t *testing.T) {
	d := &makwa.Digest{}
	if err := d.UnmarshalText([]byte("+RK3n5jz7gs_s211_xycDwiqW2ZkvPeqHZJfjkg_yc6g5u8JOTqxcQoI")); err != nil {
		t.Fatal(err)
	}

	if err := makwa.CheckPassword(modulus, sha256.New, d, password); err != nil {
		t.Error(err)
	}

	if err := makwa.CheckPassword(modulus, sha256.New, d, []byte("wink")); err != makwa.ErrBadPassword {
		t.Errorf("Error was %v, but expected ErrBadPassword", err)
	}
}

func TestHash(t *testing.T) {
	d, err := makwa.Hash(password, salt, modulus, sha256.New, 4096, false, 12)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(d.ModulusID, modulusID) {
		t.Errorf("ModulusID was %x but expected %x", d.ModulusID, modulusID)
	}

	if !bytes.Equal(d.Hash, hash) {
		t.Errorf("Hash was %x but expected %x", d.Hash, hash)
	}

	if !bytes.Equal(d.Salt, salt) {
		t.Errorf("Salt was %x but expected %x", d.Salt, salt)
	}

	if v, want := d.WorkFactor, uint(4096); v != want {
		t.Errorf("WorkFactor was %v, but expected %v", v, want)
	}

	if d.PreHash {
		t.Errorf("PreHash was %v, but expected false", d.PreHash)
	}

	if v, want := d.PostHashLen, uint(12); v != want {
		t.Errorf("PostHashLen was %v, but expected %v", v, want)
	}
}

var (
	modulusID = []byte{0xf9, 0x12, 0xb7, 0x9f, 0x98, 0xf3, 0xee, 0xb}
	hash      = []byte{0xC9, 0xCE, 0xA0, 0xE6, 0xEF, 0x09, 0x39, 0x3A, 0xB1, 0x71, 0x0A, 0x08}
	password  = []byte("Gego beshwaji'aaken awe makwa; onzaam naniizaanizi.")
	salt      = []byte{0xC7, 0x27, 0x03, 0xC2, 0x2A, 0x96, 0xD9, 0x99, 0x2F, 0x3D, 0xEA, 0x87, 0x64, 0x97, 0xE3, 0x92}
	modulus   *big.Int
)

func init() {
	n, ok := new(big.Int).SetString(`C22C40BBD056BB213AAD7C830519101AB926AE18E3E9FC9699C806E0AE5C259414A01AC1D52E873EC08046A68E344C8D74A508952842EF0F03F71A6EDC077FAA14899A79F83C3AE136F774FA6EB88F1D1AEA5EA02FC0CCAF96E2CE86F3490F4993B4B566C0079641472DEFC14BECCF48984A7946F1441EA144EA4C802A457550BA3DF0F14C090A75FE9E6A77CF0BE98B71D56251A86943E719D27865A489566C1DC57FCDEFACA6AB043F8E13F6C0BE7B39C92DA86E1D87477A189E73CE8E311D3D51361F8B00249FB3D8435607B14A1E70170F9AF36784110A3F2E67428FC18FB013B30FE6782AECB4428D7C8E354A0FBD061B01917C727ABEE0FE3FD3CEF761`, 16)
	if !ok {
		panic("couldn't parse modulus")
	}
	modulus = n
}
