package test

import (
	"context"
	"fmt"
	"github.com/guihouchang/go-elasticsearch/test/migrate"
	"github.com/guihouchang/go-elasticsearch/test/userdata"
	"github.com/olivere/elastic/v7"
	"testing"
)

func Test_Decode(t *testing.T) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)

	if err != nil {
		t.Fatal(err)
	}
	var boolQuery []elastic.Query
	boolQuery = append(boolQuery, elastic.NewTermQuery(userdata.FieldIntIiii, 1))
	ctx := context.Background()
	result, err := client.Search().Index(userdata.Label).
		Query(elastic.NewBoolQuery().Must(boolQuery...)).
		TrackTotalHits(true).Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_Migrate(t *testing.T) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)

	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	exists, err := client.IndexExists(migrate.UserDataMapping.Name).Do(ctx)
	if exists {

		// 先进行备份
		_, err := client.Reindex().Body(migrate.UserDataMapping.BackupBody()).Do(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// 删除旧索引
		_, err = client.DeleteIndex(migrate.UserDataMapping.Name).Do(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// 创建索引
		_, err = client.CreateIndex(migrate.UserDataMapping.Name).BodyJson(migrate.UserDataMapping.CreateBody()).Do(ctx)
		if err != nil {
			t.Fatal(err)
		}

		exist, err := client.IndexExists(migrate.UserDataMapping.BackupName()).Do(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if exist {
			// 恢复数据
			_, err = client.Reindex().Body(migrate.UserDataMapping.RecoveryBody()).Do(ctx)
			// 删除旧索引
			_, err = client.DeleteIndex(migrate.UserDataMapping.BackupName()).Do(ctx)
			if err != nil {
				t.Fatal(err)
			}
		}

	} else {
		_, err := client.CreateIndex(migrate.UserDataMapping.Name).BodyJson(migrate.UserDataMapping.CreateBody()).Do(context.Background())
		if err != nil {
			t.Fatal(err)
		}
	}
	//result, err := client.CreateIndex(migrate.UserDataMapping.Name).BodyJson(migrate.UserDataMapping.Body()).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}

}
