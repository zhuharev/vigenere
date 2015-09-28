package vigenere

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	alphabet = "abcdefgh "
	key      = "hhhhhhhh"

	alphabet_ru = "абвгдеёжзиклмнопрст "
	key_ru      = "мама"
)

func TestCleanString(t *testing.T) {
	Convey("test clean string", t, func() {
		alphabet := "abcdef "
		//key := "cafe"
		in := "dead beef"
		//v := New(alphabet, key)

		cleaned := cleanString(alphabet, in)

		So(cleaned, ShouldEqual, "DEAD BEEF")

		cleaned = cleanString(alphabet, in+" some")

		So(cleaned, ShouldEqual, "DEAD BEEF E")

	})
}

func TestEncrypt(t *testing.T) {
	Convey("test encryprion string", t, func() {
		in := "dead beef"
		v := New(alphabet, key)

		encrypted := v.Encrypt(in)

		So(encrypted, ShouldEqual, "BCHBG CCD")

		v = New(alphabet_ru, key_ru)
		encrypted = v.Encrypt("привет как дела")
		So(encrypted, ShouldEqual, "ЖРБВСТЛКМКЛДСЛМ")

	})
}

func TestDecrypt(t *testing.T) {
	Convey("test encryprion string", t, func() {

		//in := "dead beef"
		v := New(alphabet, key)

		decrypted := v.Decrypt("BCHBG CCD")

		So(decrypted, ShouldEqual, "DEAD BEEF")

		v = New(alphabet_ru, key_ru)
		decrypted = v.Decrypt("ЖРБВСТЛКМКЛДСЛМ")
		So(decrypted, ShouldEqual, "ПРИВЕТ КАК ДЕЛА")

	})
}
