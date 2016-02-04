package cfbypass

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const HTML = `<html><title>You are being redirected...</title>
<noscript>Javascript is required. Please enable javascript before you are allowed to see this page.</noscript>
<script>var s={},u,c,U,r,i,l=0,a,e=eval,w=String.fromCharCode,sucuri_cloudproxy_js='',S='bz1TdHJpbmcuZnJvbUNoYXJDb2RlKDB4MzcpICsgImN2Ii5jaGFyQXQoMCkgKyAiMyIgKyAiZXNlYyIuc3Vic3RyKDAsMSkgKyAgJycgK1N0cmluZy5mcm9tQ2hhckNvZGUoNTUpICsgJ2UnICsgICJkIiArICIiICsnYicgKyAgJ1hsOjInLnN1YnN0cigzLCAxKSArJzknICsgICAnJyArIApTdHJpbmcuZnJvbUNoYXJDb2RlKDB4MzIpICsgICcnICsnJysnMycgKyAgJ0NsO2YnLnN1YnN0cigzLCAxKSArJ0hhJy5zbGljZSgxLDIpKyJlc3VjdXIiLmNoYXJBdCgwKStTdHJpbmcuZnJvbUNoYXJDb2RlKDQ5KSArICAnJyArJycrImQiICsgIjAiICsgICcnICsnJysnMycgKyAgJ2QnICsgICIiICsiYyIuc2xpY2UoMCwxKSArICAnJyArIjUiICsgICcnICsiNiIuc2xpY2UoMCwxKSArICc/Yicuc2xpY2UoMSwyKSsiYSIgKyAiMiIuc2xpY2UoMCwxKSArICI5IiArICAnJyArIjYiICsgICcnICsgCic4ZzJmJy5zdWJzdHIoMywgMSkgKyAnJyArIAonNicgKyAgICcnICsgCiJkIi5zbGljZSgwLDEpICsgICcnICsgCiJhc3VjdXIiLmNoYXJBdCgwKSsnJztkb2N1bWVudC5jb29raWU9J3MnKyd1cycuY2hhckF0KDApKydjJy5jaGFyQXQoMCkrJ3UnKydyc3VjdScuY2hhckF0KDApICArJ3N1aScuY2hhckF0KDIpKydfJy5jaGFyQXQoMCkrJ3N1Y3VjJy5jaGFyQXQoNCkrICdzdWN1cmlsJy5jaGFyQXQoNikrJ29zdWMnLmNoYXJBdCgwKSsgJ3UnKydkc3VjdXInLmNoYXJBdCgwKSsgJ3BzdWN1Jy5jaGFyQXQoMCkgICsncicrJ28nKycnKyd4JysnJysnc3knLmNoYXJBdCgxKSsnX3MnLmNoYXJBdCgwKSsndScrJ3UnKydpJysnc3VjdXJpZCcuY2hhckF0KDYpKydfJysnMCcrJ2JzdWN1cicuY2hhckF0KDApKyAnMCcrJzgnLmNoYXJBdCgwKSsnc3VjdXJpYScuY2hhckF0KDYpKyc5JysnJysnOXN1Y3UnLmNoYXJBdCgwKSAgKydzZCcuY2hhckF0KDEpKydhc3VjdScuY2hhckF0KDApICArIj0iICsgbzsgbG9jYXRpb24ucmVsb2FkKCk7';L=S.length;U=0;r='';var A='ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/';for(u=0;u<64;u++){s[A.charAt(u)]=u;}for(i=0;i<L;i++){c=s[S.charAt(i)];U=(U<<6)+c;l+=6;while(l>=8){((a=(U>>>(l-=8))&0xff)||(i<(L-2)))&&(r+=w(a));}}e(r);</script><script type="text/javascript">
//<![CDATA[
try{if (!window.CloudFlare) {var CloudFlare=[{verbose:0,p:1453883424,byc:0,owlid:"cf",bag2:1,mirage2:0,oracle:0,paths:{cloudflare:"/cdn-cgi/nexp/dok3v=1613a3a185/"},atok:"3f4377b8b8a2d59bcc2ae029cf3e52fb",petok:"fcc14026cd742af2c0876b365d33b93f9f3ad2dc-1454487982-1800",zone:"mangaumaru.com",rocket:"0",apps:{"ga_key":{"ua":"UA-43720140-2","ga_bs":"2"}},sha2test:0}];!function(a,b){a=document.createElement("script"),b=document.getElementsByTagName("script")[0],a.async=!0,a.src="//ajax.cloudflare.com/cdn-cgi/nexp/dok3v=38857570ac/cloudflare.min.js",b.parentNode.insertBefore(a,b)}()}}catch(e){};
//]]>
</script>
<script type="text/javascript">
/* <![CDATA[ */
var _gaq = _gaq || [];
_gaq.push(['_setAccount', 'UA-43720140-2']);
_gaq.push(['_trackPageview']);

(function() {
var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
})();

(function(b){(function(a){"__CF"in b&&"DJS"in b.__CF?b.__CF.DJS.push(a):"addEventListener"in b?b.addEventListener("load",a,!1):b.attachEvent("onload",a)})(function(){"FB"in b&&"Event"in FB&&"subscribe"in FB.Event&&(FB.Event.subscribe("edge.create",function(a){_gaq.push(["_trackSocial","facebook","like",a])}),FB.Event.subscribe("edge.remove",function(a){_gaq.push(["_trackSocial","facebook","unlike",a])}),FB.Event.subscribe("message.send",function(a){_gaq.push(["_trackSocial","facebook","send",a])}));"twttr"in b&&"events"in twttr&&"bind"in twttr.events&&twttr.events.bind("tweet",function(a){if(a){var b;if(a.target&&a.target.nodeName=="IFRAME")a:{if(a=a.target.src){a=a.split("#")[0].match(/[^?=&]+=([^&]*)?/g);b=0;for(var c;c=a[b];++b)if(c.indexOf("url")===0){b=unescape(c.split("=")[1]);break a}}b=void 0}_gaq.push(["_trackSocial","twitter","tweet",b])}})})})(window);
/* ]]> */
</script>
<script type="text/javascript">
//<![CDATA[
try{(function(a){var b="http://",c="www.mangaumaru.com",d="/cdn-cgi/cl/",e="img.gif",f=new a;f.src=[b,c,d,e].join("")})(Image)}catch(e){}
//]]>
</script></html>
`

func main() {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer([]byte(HTML)))
	if err != nil {
		return
	}
	splits := regexp.MustCompile(".{1,15}=").Split(strings.Replace(DecodeScript(doc), ";location.reload();", "", 1), -1)
	fmt.Println(GetCookieValue(strings.Split(splits[1], "+")))
}

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

func DecodeScript(doc *goquery.Document) string {
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
	return regexp.MustCompile("\\s+").ReplaceAllString(strings.Replace(string(buf.Buffer), "\\", "", -1), "")
}

func GetCookieValue(splits []string) (val string) {
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

func GetCookieKey(splits []string) (key string) {
	return
}
