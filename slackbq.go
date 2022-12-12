package slackbq

import (
	"cloud.google.com/go/bigquery"
	"context"
	"encoding/json"
	"net/http"
)

func CreateDataset(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Dataset string `json:"dataset"`
		Project string `json:"project"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		panic(err)
	}
	err := save(d.Project, d.Dataset)
	if err != nil {
		panic(err)
	}
}
func save(projectID, datasetID string) error {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	meta := &bigquery.DatasetMetadata{
		Location: "US",
	}
	if err := client.Dataset(datasetID).Create(ctx, meta); err != nil {
		panic(err)
	}
	return nil
}
