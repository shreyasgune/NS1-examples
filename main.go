package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/ns1/ns1-go.v2/rest/model/monitor"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"

	api "gopkg.in/ns1/ns1-go.v2/rest"
)

func main() {
	k := os.Getenv("NS1_KEY")
	if k == "" {
		fmt.Println("NS1_KEY environment variable is not set, giving up")
		os.Exit(1)
	}

	httpClient := &http.Client{Timeout: time.Second * 10}
	client := api.NewClient(httpClient, api.SetAPIKey(k))
	log.Print(client)
}

// GetZonesNS1 gets you the zones
func GetZonesNS1(client *api.Client) []*dns.Zone {
	zones, _, err := client.Zones.List()
	if err != nil {
		log.Fatal(err)
	}
	return zones
}

// GetHostedZonesByNameNS1 gets you the DNS records for a given zone
func GetHostedZonesByNameNS1(client *api.Client, zone string) *dns.Zone {
	activeZone, _, err := client.Zones.Get(zone)
	if err != nil {
		log.Fatal(err)
	}
	return activeZone
}

// GetRecordsNS1 gets you the DNS records for a given zone and domain
func GetRecordsNS1(client *api.Client, zone, domain, recordType string) *dns.Record {
	records, _, err := client.Records.Get(zone, domain, recordType)
	if err != nil {
		log.Fatal(err)
	}
	return records
}

// CreateRecordNS1 gets you the hosted zones
func CreateRecordNS1(client *api.Client, rec *dns.Record) *http.Response {
	res, err := client.Records.Create(rec)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// DeleteRecordNS1 gets you the hosted zones
func DeleteRecordNS1(client *api.Client, zone, domain, recordType string) *http.Response {
	res, err := client.Records.Delete(zone, domain, recordType)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// UpdateRecordsNS1 updates existing record
func UpdateRecordsNS1(client *api.Client, rec *dns.Record) *http.Response {
	res, err := client.Records.Update(rec)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// GetHealthcheckStatusNS1 gets you healthcheck status
func GetHealthcheckStatusNS1(client *api.Client, healthcheckID string) *monitor.Job {
	healthcheck, _, err := client.Jobs.Get(healthcheckID)
	if err != nil {
		log.Fatal(err)
	}
	return healthcheck
}

// GetAllHealthChecksNS1 gets you a list of existing healthchecks
func GetAllHealthChecksNS1(client *api.Client) []*monitor.Job {
	healthchecks, _, err := client.Jobs.List()
	if err != nil {
		log.Fatal(err)
	}
	return healthchecks
}

// CreateHealthcheckNS1 creates a healthcheck
func CreateHealthcheckNS1(client *api.Client, mj *monitor.Job) *http.Response {
	res, err := client.Jobs.Create(mj)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// UpdateHealthcheckNS1 updates an existing healthcheck
func UpdateHealthcheckNS1(client *api.Client, mj *monitor.Job) *http.Response {
	res, err := client.Jobs.Update(mj)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// DeleteHealthcheckNS1 deletes a healthcheck
func DeleteHealthcheckNS1(client *api.Client, healthcheckID string) *http.Response {
	res, err := client.Jobs.Delete(healthcheckID)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
