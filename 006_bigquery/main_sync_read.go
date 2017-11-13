package main

import (
	"os"
	"golang.org/x/net/context"
	"cloud.google.com/go/bigquery"
	"log"
	"fmt"
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
	q := client.Query(`
		SELECT year, SUM(number) as num
		FROM [bigquery-public-data:usa_names.usa_1910_2013]
		WHERE name = "William"
		GROUP BY year
		ORDER BY year desc
		LIMIT 3
	`)
	q.QueryConfig.UseLegacySQL = true
	itr, err := q.Read(ctx)
	if err != nil {
		log.Fatalf("Failed to read the bigquery result. %v", err)
	}

	// Fetch rows.
	//for {
	//	var values []bigquery.Value
	//	err := itr.Next(&values)
	//	if err == iterator.Done {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalf("Somethin happended. %v", err)
	//	}
	//	fmt.Printf("year=%d, sum=%v\n", values[0], values[1])
	//}

	// Or, You can also use user-defined struct.
	type Count struct {
		Year int
		Num  int
	}
	for {
		var c Count
		err := itr.Next(&c)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Somethin happended. %v", err)
		}
		fmt.Println(c)
	}

}


























