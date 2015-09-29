# vigenere
alphabet independent vigenere chipher

## gui

***GUI available*** [zhuharev/vigenere-gui](https://github.com/zhuharev/vigenere-gui)

## usage

```
...
import "github.com/zhuharev/vigenere"
...
key := "yep"
vc := vigenere.New(vigenere.EN_WITH_SPACE, key)

encrypted := vc.Encrypt("hello, Don Huan")
// output: EI ISOASBXLIYR
decrypted := vc.Decrypt("EI ISOASBXLIYR")
// output: HELLO DON HUAN
```

## similar projects and references

[http://rosettacode.org/wiki/](http://rosettacode.org/wiki/Vigen%C3%A8re_cipher#Go)
[turnerd18/vigenere-cipher](https://github.com/turnerd18/vigenere-cipher)
