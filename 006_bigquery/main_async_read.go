package main

import (
	"os"
	"golang.org/x/net/context"
	"cloud.google.com/go/bigquery"
	"log"
	"fmt"
)

/*
	Samples for how to use BigQuery.

	Dependencies:
		$ go get -u cloud.google.com/go/bigquery

	Environment Variables:
		- BIGQUERY_PROJECT_ID: 	Set your GCP project id.

	References:
		- https://cloud.google.com/bigquery/docs/reference/libraries
		- https://godoc.org/cloud.google.com/go/bigquery
 */

// TODO 認証方法についてはいくつかあるみたい.
// https://cloud.google.com/docs/authentication?hl=ja

func main() {

	// Context for BigQuery client.
	ctx := context.Background()

	// Set the auth file path.
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/etc/mytest/bigquery-sample-auth.json")

	// Get a project id.
	projectId := os.Getenv("BIGQUERY_PROJECT_ID")

	// Create a client for BigQuery.
	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a client. %v", err)
	}
	fmt.Println(client)

	//q := client.Query("SELECT year, SUM(number) as num FROM `bigquery-public-data`.usa_names.usa_1910_2013 WHERE name = \"William\" GROUP BY year ORDER BY year")
	//q.QueryConfig.UseStandardSQL = true
	q := client.Query(`
		SELECT year, SUM(number) as num
		FROM [bigquery-public-data:usa_names.usa_1910_2013]
		WHERE name = "William"
		GROUP BY year
		ORDER BY year
	`)
	q.QueryConfig.UseLegacySQL = true

	// Execute a query async.
	job, err := q.Run(ctx)
	if err != nil {
		log.Fatalf("Failed to start a job. %v", err)
	}
	fmt.Println("job:", job)

	// Get a jobID.
	jobID := job.ID()
	fmt.Println("Job ID is ", jobID)

	// Create a job from a jobID.
	job_retreived, err := client.JobFromID(ctx, jobID)
	if err != nil {
		log.Fatalf("Failed to retreive a job from jobID. %v", err)
	}
	fmt.Println("job_retreived:", job_retreived)

}

