package main

import (
	_ "github.com/mailru/go-clickhouse"

	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	var (
		dsn           = flag.String("dsn", "http://127.0.0.1:8123/default", "")
		messages      = flag.Int("messages-num", 10000, "")
		chBulkActions = flag.Int("ch-bulk-actions", 1000, "")
		workers       = flag.Int("workers", runtime.GOMAXPROCS(-1), "")
	)
	flag.Parse()

	connect, err := sql.Open("clickhouse", *dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		log.Fatal(err)
	}
	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			date Date MATERIALIZED toDate(timestamp),
			timestamp DateTime,
			http_scheme String,
			http_proto String,
			http_method String,
			remote_addr String,
			user_agent String,
			uri String	
		) ENGINE = MergeTree(date, (timestamp), 32768)
	`)
	if err != nil {
		log.Fatalf("could not migrate: %v", err)
	}

	messagesPerWorker := *messages / *workers
	log.Printf("workers: %d, message per worker: %d\n", *workers, messagesPerWorker)

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(*workers)
	for w := 1; w <= *workers; w++ {
		go func(id int) {
			defer wg.Done()

			batchCommits := messagesPerWorker / *chBulkActions
			for i := 0; i < batchCommits; i++ {
				tx, err := connect.Begin()
				if err != nil {
					log.Fatal(err)
				}

				stmt, err := tx.Prepare(`INSERT INTO logs (timestamp, http_scheme, http_proto, http_method, remote_addr, user_agent, uri) VALUES (?, ?, ?, ?, ?, ?, ?)`)
				if err != nil {
					log.Fatal(err)
				}

				for j := 0; j < *chBulkActions; j++ {
					_, err = stmt.Exec(
						time.Now(),
						"http",
						"HTTP/1.0",
						methods[rand.Intn(len(methods))],
						hosts[rand.Intn(len(hosts))],
						userAgents[rand.Intn(len(userAgents))],
						fmt.Sprintf("%s://%s/%s", schemes[rand.Intn(len(schemes))], hosts[rand.Intn(len(hosts))], urls[rand.Intn(len(urls))]),
					)
					if err != nil {
						log.Fatal(err)
					}
				}

				if err := tx.Commit(); err != nil {
					log.Fatal(err)
				}
				log.Printf("[worker %d] commit %d items (%d/%d)", id, *chBulkActions, i+1, batchCommits)
			}
		}(w)
	}
	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("total %s", elapsed)

	// rows, err := connect.Query("SELECT country_code, os_id, browser_id, categories, action_day, action_time FROM example")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for rows.Next() {
	// 	var (
	// 		country               string
	// 		os, browser           uint8
	// 		categories            []int16
	// 		actionDay, actionTime time.Time
	// 	)
	// 	if err := rows.Scan(&country, &os, &browser, &categories, &actionDay, &actionTime); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s", country, os, browser, categories, actionDay, actionTime)
	// }
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
	"Mozilla/5.0 Windows NT 10.0; Win64; x64 AppleWebKit/537.36 KHTML, like Gecko Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 Windows NT 6.1; Win64; x64 AppleWebKit/537.36 KHTML, like Gecko Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 Macintosh; Intel Mac OS X 10_12_6 AppleWebKit/537.36 KHTML, like Gecko Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 Macintosh; Intel Mac OS X 10_12_6 AppleWebKit/604.1.38 KHTML, like Gecko Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 Windows NT 10.0; Win64; x64; rv:56.0 Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 Macintosh; Intel Mac OS X 10_13 AppleWebKit/604.1.38 KHTML, like Gecko Version/11.0 Safari/604.1.38",
}

var schemes = []string{"http", "https"}

var methods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodHead}
