package main

import (
	"os"
	"golang.org/x/net/context"
	"cloud.google.com/go/bigquery"
	"log"
	"fmt"
	"time"
	"google.golang.org/api/iterator"
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

	// Query complete (2.0s elapsed, 126 MB processed)
	//q := client.Query(`
	//	SELECT year, SUM(number) as num
	//	FROM [bigquery-public-data:usa_names.usa_1910_2013]
	//	WHERE name = "William"
	//	GROUP BY year
	//	ORDER BY year desc
	//	LIMIT 3
	//`)

	// Query complete (7.6s elapsed, 1.61 GB processed)
	q := client.Query(`
		SELECT label_name, COUNT(1) cnt
		FROM [bigquery-public-data:open_images.labels]
		WHERE confidence >= 0.85
		GROUP BY label_name
		ORDER BY cnt DESC
		LIMIT 3
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
	//job_retreived, err := client.JobFromID(ctx, jobID)
	//if err != nil {
	//	log.Fatalf("Failed to retreive a job from jobID. %v", err)
	//}
	//fmt.Println("job_retreived:", job_retreived)

	// Wait until the job will be finished.
	for {
		println("Getting status.")
		jobStatus, err := job.Status(ctx)
		if err != nil {
			log.Fatalf("jos.Status failed. %v", err)
		}
		if jobStatus.Err() != nil {
			log.Fatalf("jos.Status failed. %v", jobStatus.Err())
		}
		if jobStatus.Done() {
			break
		}
		println("Waiting 1s...")
		time.Sleep(1)
	}

	// Get the result.
	println("Getting results.")
	it, err := job.Read(ctx)
	if err != nil {
		log.Fatalf("job.Read() failed. %v", err)
	}
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Retreiving an iterator failed. %v")
		}
		fmt.Printf("col1=%v, col2=%v\n", values[0], values[1])
	}

	println("Finished.")
}


























