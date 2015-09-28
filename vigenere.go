package vigenere

import (
	"strings"
)

type VigenereChipher struct {
	Alphabet string
	Key      string
}

func New(alphabet string, key string) *VigenereChipher {
	vc := new(VigenereChipher)
	vc.Alphabet = strings.ToUpper(alphabet)
	vc.Key = strings.ToUpper(key)
	return vc
}

func (vc *VigenereChipher) AlphabetLen() int {
	return len([]rune(vc.Alphabet))
}

func (vc *VigenereChipher) AlphabetMap() map[rune]int {
	m := map[rune]int{}
	for i, v := range []rune(vc.Alphabet) {
		m[v] = i
	}
	return m
}

func (vc *VigenereChipher) Encrypt(in string) string {
	in = cleanString(vc.Alphabet, in)
	key := vc.Key
	var res []rune
	rkey := []rune(key)
	m := vc.AlphabetMap()
	for i, v := range []rune(in) {
		indx := (m[v] + m[rkey[i%len(rkey)]]) % vc.AlphabetLen()
		res = append(res, []rune(vc.Alphabet)[indx])
	}
	return string(res)
}

func (vc *VigenereChipher) Decrypt(in string) string {
	in = cleanString(vc.Alphabet, in)
	key := vc.Key
	var res []rune
	rkey := []rune(key)
	m := vc.AlphabetMap()
	for i, v := range []rune(in) {
		indx := (m[v] - m[rkey[i%len(rkey)]] + vc.AlphabetLen()) % vc.AlphabetLen()
		res = append(res, []rune(vc.Alphabet)[indx])
	}
	return string(res)
}

func cleanString(alphabet, in string) string {
	sc := func(str, chr string) string {
		return strings.Map(func(r rune) rune {
			if strings.IndexRune(chr, r) < 0 {
				return -1
			}
			return r
		}, str)
	}

	alphabet = strings.ToUpper(alphabet)
	out := strings.ToUpper(in)
	out = sc(out, alphabet)

	return out
}
