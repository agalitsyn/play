package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	var (
		logsN = flag.Int("logs-num", 1000, "")
	)
	flag.Parse()

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	for index := 0; index < *logsN; index++ {
		randReq := logrus.Fields{
			"http_scheme": "http",
			"http_proto":  "HTTP/1.0",
			"http_method": methods[rand.Intn(len(methods))],
			"remote_addr": hosts[rand.Intn(len(hosts))],
			"user_agent":  userAgents[rand.Intn(len(userAgents))],
			"uri":         fmt.Sprintf("%s://%s/%s", schemes[rand.Intn(len(schemes))], hosts[rand.Intn(len(hosts))], urls[rand.Intn(len(urls))]),
			"@timestamp":  time.Now(),
		}
		logrus.WithFields(randReq).Infof("rand req %d", index)
	}
}

var urls = []string{
	"foo",
	"bar",
	"foobar",
	"foo-bar",
	"baz",
	"foo/bar/baz",
	"foo-baz",
	"bar/baz",
	"foobar/baz",
}

var hosts = []string{
	"10.42.100.214:57692",
	"10.42.248.227:55144",
	"10.42.129.107:41446",
	"10.42.100.214:57744",
	"10.42.248.227:57266",
	"10.42.129.107:41510",
	"10.42.0.1:34354",
	"10.42.100.214:57798",
	"10.42.248.227:59702",
	"10.42.129.107:41572",
	"10.42.179.221:35418",
	"10.42.100.214:57856",
	"10.42.248.227:33530",
	"10.42.0.1:36894",
	"10.42.129.107:41642",
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

var schemes = []string{"http", "https"}

var methods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodHead}
