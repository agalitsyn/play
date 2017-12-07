package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"time"

	"gopkg.in/olivere/elastic.v5"
)

func main() {
	var (
		elasticScheme      = flag.String("elastic-scheme", "http", "")
		elasticHost        = flag.String("elastic-host", "localhost", "")
		elasticPort        = flag.String("elastic-port", "9200", "")
		elasticIndex       = flag.String("elastic-index", "mylog", "")
		elasticBulkActions = flag.Int("elastic-bulk-actions", 1000, "")
		elasticBulkSize    = flag.Int("elastic-bulk-size", 5*1024*1024, "")
		messages           = flag.Int("messages-num", 10000, "")
		watchInt           = flag.Duration("watch-interval", 10*time.Second, "")
		workers            = flag.Int("workers", runtime.GOMAXPROCS(-1), "")
	)
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	u := url.URL{
		Scheme: *elasticScheme,
		Host:   net.JoinHostPort(*elasticHost, *elasticPort),
	}
	client, err := elastic.NewClient(elastic.SetErrorLog(log.New(os.Stdout, "EL: ", log.LstdFlags)), elastic.SetURL(u.String()), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}

	exists, err := client.IndexExists(*elasticIndex).Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		createIndex, err := client.CreateIndex(*elasticIndex).Do(ctx)
		if err != nil {
			log.Fatal(err)
		}
		if !createIndex.Acknowledged {
			log.Fatal()
		}
	}

	svc := client.BulkProcessor().Workers(*workers).BulkActions(*elasticBulkActions).BulkSize(*elasticBulkSize)
	proc, err := svc.Stats(true).Do(ctx)
	if err != nil {
		log.Fatal(err)
	}

	printStats := func() {
		stats := proc.Stats()
		fmt.Printf("workers: %d\n", len(stats.Workers))
		fmt.Printf("successed: %d\n", stats.Succeeded)
		fmt.Printf("created: %d\n", stats.Created)
		fmt.Printf("failed: %d\n", stats.Failed)
		fmt.Printf("flushed: %d\n", stats.Flushed)
		fmt.Printf("commited: %d\n", stats.Committed)
		fmt.Println()
	}

	start := time.Now()

	go func() {
		log.Println("start watcher")
		for {
			select {
			case <-time.After(*watchInt):
				elapsed := time.Since(start)
				log.Printf("took %s", elapsed)
				printStats()
			case <-ctx.Done():
				return
			}
		}
	}()

	for i := 1; i <= *messages; i++ {
		randReq := map[string]interface{}{
			"http_scheme": "http",
			"http_proto":  "HTTP/1.0",
			"http_method": methods[rand.Intn(len(methods))],
			"remote_addr": hosts[rand.Intn(len(hosts))],
			"user_agent":  userAgents[rand.Intn(len(userAgents))],
			"uri":         fmt.Sprintf("%s://%s/%s", schemes[rand.Intn(len(schemes))], hosts[rand.Intn(len(hosts))], urls[rand.Intn(len(urls))]),
		}
		request := elastic.NewBulkIndexRequest().Index(*elasticIndex).Type("log").OpType("create").Id(fmt.Sprintf("%d", i)).Doc(randReq)
		proc.Add(request)
	}

	if err := proc.Flush(); err != nil {
		log.Fatal(err)
	}

	printStats()

	if err := proc.Close(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("total %s", elapsed)
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
