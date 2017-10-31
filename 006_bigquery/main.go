package main

import (
	"os"
	//"fmt"
	"golang.org/x/net/context"
	//"golang.org/x/oauth2/jwt"
	//"google.golang.org/api/bigquery/v2"
	//"log"

	"cloud.google.com/go/bigquery"
	"log"
)

/*
	Samples for how to use BigQuery.

	Dependencies:
		$ go get -u cloud.google.com/go/bigquery

	Environment Variables:
		- BIGQUERY_PROJECT_ID: 	Set your GCP project id.
		- BIGQUERY_API_KEY: 	Your API Key (PEM).
		- BIGQUERY_EMAIL: 		Your Email Address for auth.

	References:
		- https://cloud.google.com/bigquery/docs/reference/libraries?hl=ja#client-libraries-usage-go

 */

// TODO 認証方法についてはいくつかあるみたい.

func main() {

	ctx := context.Background()


	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "~/gcp/bigquery-sample-auth.json")

	//
	//client, err := bigquery.NewClient(ctx)
	//if err != nil {
	//	log.Fatalf("Failed to create a client: %v", err)
	//}
	//
	//client.Q

	//
	//projectId := os.Getenv("BIGQUERY_PROJECT_ID")
	//fmt.Println("ProjectId:", projectId)
	//
	//fmt.Println(os.Getenv("GIBQUERY_EMAIL"))
	//fmt.Println(os.Getenv("BIGQUERY_API_KEY"))
	//
	//cfg := jwt.Config{
	//	Email: 		os.Getenv("GIBQUERY_EMAIL"),
	//	PrivateKey: []byte(os.Getenv("BIGQUERY_API_KEY")),
	//	Scopes: 	[]string{bigquery.BigqueryScope},
	//	TokenURL: 	"https://accounts.google.com/o/oauth2/token",
	//}
	//client := cfg.Client(ctx)
	//conn, err := bigquery.New(client)
	//if err != nil {
	//	log.Fatalf("Failed to create a client: %v", err)
	//}
	//fmt.Println(conn)
	//
	//// Query.
	//result, err := conn.Jobs.Query(projectId, &bigquery.QueryRequest{
	//	Query: `
	//		SELECT year, SUM(number) as num
	//		FROM [bigquery-public-data:usa_names.usa_1910_2013]
	//		WHERE name = "William"
	//		GROUP BY year
	//		ORDER BY year
	//	`,
	//}).Do()
	//
	//if err != nil {
	//	log.Fatalf("Query execute error. %v", err)
	//}
	//
	//log.Printf("result: %v\n", result)
	//
	//// Get results.
	//for _, row := range result.Rows {
	//	for _, cell := range row.F {
	//		fmt.Print(cell.V)
	//		fmt.Print(",")
	//	}
	//	fmt.Print("\n")
	//}









}


























