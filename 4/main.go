// Package main provides ...
package main

import (
	"log"
	"net/http"
	"net/url"

	influx "github.com/influxdata/influxdb/client"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

// Client allows sending batches of Prometheus samples to InfluxDB.
type Client struct {
	client          *influx.Client
	database        string
	retentionPolicy string
}

// NewClient creates a new Client.
func NewClient(conf influx.Config, db string, rp string) *Client {
	c, err := influx.NewClient(conf)
	// Currently influx.NewClient() *should* never return an error.
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		client:          c,
		database:        db,
		retentionPolicy: rp,
	}
}

//fieldsFromMetric extracts InfluxDB fields from a Prometheus metric.
func fieldsFromMetric(m []*dto.LabelPair) map[string]interface{} {
	fields := make(map[string]interface{}, len(m))
	for _, l := range m {
		fields[l.GetName()] = l.GetValue()
	}
	return fields
}

// Send sends a batch of samples to InfluxDB via its HTTP API.
func (c *Client) Send(metricsFamily map[string]*dto.MetricFamily) error {
	points := make([]influx.Point, 0, len(metricsFamily))
	for _, metricFamily := range metricsFamily {
		metric := metricFamily.GetMetric()[0]

		points = append(points, influx.Point{
			Measurement: metricFamily.GetName(),
			Fields:      fieldsFromMetric(metric.GetLabel()),
			Tags: map[string]string{
				"type":  metricFamily.GetType().String(),
				"value": metricFamily.String(),
			},
		})
	}

	bps := influx.BatchPoints{
		Points:          points,
		Database:        c.database,
		RetentionPolicy: c.retentionPolicy,
	}
	_, err := c.client.Write(bps)
	return err
}

func main() {
	resp, err := http.Get("http://localhost:2379/metrics")
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	defer resp.Body.Close()

	parser := new(expfmt.TextParser)
	metrics, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		log.Fatalf("Pasing failed: %v", err)
	}

	connS := "http://localhost:8086"
	u, err := url.Parse(connS)
	if err != nil {
		log.Fatal(err)
	}

	conf := influx.Config{
		URL: *u,
	}

	client := NewClient(conf, "mydb", "")
	if err := client.Send(metrics); err != nil {
		log.Fatal(err)
	}
}
