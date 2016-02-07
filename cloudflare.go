package cfbypass

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type BitBuffer struct {
	Buffer []byte
	tw     uint32 // 24-bit
	lv     byte
}

func (b *BitBuffer) Write(bin byte) {
	if bin >= 64 {
		panic(fmt.Sprintf("%d overflows 64", bin))
	}
	b.tw = b.tw<<6 | uint32(bin)
	if b.lv++; b.lv == 4 {
		b.Finalize()
	}
}

func (b *BitBuffer) Finalize() {
	start := int(b.lv * 6)
	end := start - 8
	for ; start > 0; func() { start -= 8; end -= 8 }() {
		if end < 0 {
			end = 0
		}
		b.Buffer = append(b.Buffer, byte((b.tw>>uint(end))&((1<<uint(start))-1)))
	}
	b.lv = 0
	b.tw = 0
}

func DecodeScript(doc *goquery.Document) []string {
	script := doc.Find("script").First().Text()
	crypto := regexp.MustCompile(".='.*';.=.\\.length;").FindString(script)
	crypto = crypto[8 : len(crypto)-13]
	bbox := regexp.MustCompile("var .='.*'").FindString(script)
	bbox = bbox[7 : len(bbox)-1]
	dic := make(map[rune]byte)
	for i, b := range bbox {
		dic[b] = byte(i) & 63
	}
	var c byte
	buf := new(BitBuffer)
	for _, s := range crypto {
		c = dic[s]
		buf.Write(c)
	}
	buf.Finalize()
	return regexp.MustCompile("[a-zA-Z]{1,15}=").Split(
		strings.Replace(
			regexp.MustCompile("\\s+").ReplaceAllString(
				strings.Replace(string(buf.Buffer), "\\", "", -1), "",
			),
			";location.reload();", "", 1),
		-1)[1:]
}

func GetCookieValue(s string) (val string) {
	splits := strings.Split(s, "+")
	for _, split := range splits {
		if split == "" || split == "\"\"" || split == "''" ||
			split == "=" || regexp.MustCompile("[.\"']+").FindString(split) == "" {
			continue
		}
		switch {
		case regexp.MustCompile("(\"|').+(\"|')\\.(charAt|substr|slice)\\([,0-9]+\\)").MatchString(split):
			str := regexp.MustCompile("(\"|').+(\"|')").FindString(split)
			str = str[1 : len(str)-1]
			f := regexp.MustCompile("(charAt|substr|slice)").FindString(split)
			arg := regexp.MustCompile("\\([,0-9]+\\)").FindString(split)
			arg = arg[1 : len(arg)-1]
			switch f {
			case "charAt":
				i, _ := strconv.Atoi(arg)
				val += str[i : i+1]
			case "substr":
				args := strings.Split(arg, ",")
				start, _ := strconv.Atoi(args[0])
				length, _ := strconv.Atoi(args[1])
				val += str[start : start+length]
			case "slice":
				args := strings.Split(arg, ",")
				start, _ := strconv.Atoi(args[0])
				end, _ := strconv.Atoi(args[1])
				val += str[start:end]
			}
		case len(split) >= 19 && split[:19] == "String.fromCharCode":
			char := split[20 : len(split)-1]
			if char[:2] == "0x" {
				v, err := hex.DecodeString(char[2:4])
				if err != nil {
					fmt.Printf("Error: %v (%s)", err, split)
				} else {
					val += string(v[0])
				}
			} else {
				v, err := strconv.Atoi(char)
				if err != nil {
					fmt.Printf("Error: %v (%s)", err, split)
				} else {
					val += string(v)
				}
			}
		case len(split) == 3 && split[0] == split[2] &&
			(split[0:1] == "\"" || split[0:1] == "'"):
			val += split[1:2]
		}
	}
	return
}

func GetCookieKey(s string) (key string) {
	splits := strings.Split(s, "+")
	for _, split := range splits {
		if split == "" || split == "\"\"" || split == "''" ||
			split == "=" || regexp.MustCompile("[.\"']+").FindString(split) == "" {
			continue
		}
		switch {
		case regexp.MustCompile("(\"|').+(\"|')\\.(charAt|substr|slice)\\([,0-9]+\\)").MatchString(split):
			str := regexp.MustCompile("(\"|').+(\"|')").FindString(split)
			str = str[1 : len(str)-1]
			f := regexp.MustCompile("(charAt|substr|slice)").FindString(split)
			arg := regexp.MustCompile("\\([,0-9]+\\)").FindString(split)
			arg = arg[1 : len(arg)-1]
			switch f {
			case "charAt":
				i, _ := strconv.Atoi(arg)
				key += str[i : i+1]
			case "substr":
				args := strings.Split(arg, ",")
				start, _ := strconv.Atoi(args[0])
				length, _ := strconv.Atoi(args[1])
				key += str[start : start+length]
			case "slice":
				args := strings.Split(arg, ",")
				start, _ := strconv.Atoi(args[0])
				end, _ := strconv.Atoi(args[1])
				key += str[start:end]
			}
		case len(split) >= 19 && split[:19] == "String.fromCharCode":
			char := split[20 : len(split)-1]
			if char[:2] == "0x" {
				v, err := hex.DecodeString(char[2:4])
				if err != nil {
					fmt.Printf("Error: %v (%s)", err, split)
				} else {
					key += string(v[0])
				}
			} else {
				v, err := strconv.Atoi(char)
				if err != nil {
					fmt.Printf("Error: %v (%s)", err, split)
				} else {
					key += string(v)
				}
			}
		case len(split) == 3 && split[0] == split[2] &&
			(split[0:1] == "\"" || split[0:1] == "'") && split[1:2] != "=":
			key += split[1:2]
		}
	}
	return
}
