package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type HashServer struct {
	Url            string
	ResponseRegexp string
}

func main() {

	fmt.Println(`
	##################################### 
	#             MD5 Cracker           #
	#-----------------------------------#
	#       github.com/dursunkatar      #
	#####################################
	`)

	servers := []HashServer{
		HashServer{
			Url: "http://www.nitrxgen.net/md5db/%s",
		},
		HashServer{
			Url: "https://md5decrypt.net/Api/api.php?hash=%s&hash_type=md5&email=deanna_abshire@proxymail.eu&code=1152464b80a61728",
		},
		HashServer{
			Url:            "https://hashtoolkit.com/reverse-hash/?hash=%s",
			ResponseRegexp: "<span title=\"decrypted md5 hash\">(.*?)</span>",
		},
	}

	hash := os.Args[1]

	for _, cserver := range servers {
		if ok, result := cserver.crack(hash); ok {
			fmt.Println(result)
			os.Exit(0)
		}
	}

	fmt.Println("Not Found!")
}

func (m HashServer) regex(result string) string {
	re := regexp.MustCompile(m.ResponseRegexp)
	match := re.FindStringSubmatch(result)
	if len(match) == 0 {
		return ""
	}
	return match[1]
}

func (m HashServer) crack(hash string) (bool, string) {
	res, err := http.Get(fmt.Sprintf(m.Url, hash))
	if err != nil {
		return false, ""
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	body := strings.Trim(string(bytes), " ")

	if err != nil || body == "" {
		return false, ""
	}

	if m.ResponseRegexp != "" {
		result := m.regex(body)
		if result != "" {
			return true, m.regex(body)
		}
		return false, ""
	}
	return true, body
}
