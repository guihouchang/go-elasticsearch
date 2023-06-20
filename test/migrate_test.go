package test

import (
	"context"
	"fmt"
	"github.com/guihouchang/go-elasticsearch/test/migrate"
	"github.com/olivere/elastic"
	"testing"
)

func Test_Migrate(t *testing.T) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)

	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	//exists, err := client.IndexExists(migrate.UserDataMapping.Name).Do(ctx)

	//return
	result, err := client.PutMapping().Index(migrate.UserDataMapping.Name).BodyJson(migrate.UserDataMapping.AlterBody()).Do(ctx)
	//result, err := client.CreateIndex(migrate.UserDataMapping.Name).BodyJson(migrate.UserDataMapping.Body()).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(result)
}
