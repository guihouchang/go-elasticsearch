package test

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/guihouchang/go-elasticsearch/test/migrate"
	"github.com/guihouchang/go-elasticsearch/test/userdata"
	"github.com/olivere/elastic/v7"
	"reflect"
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

	datas, err := UnserizerUserData(result)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(datas)
}

func compareMaps(map1, map2 interface{}) bool {
	// 比较类型
	if reflect.TypeOf(map1) != reflect.TypeOf(map2) {
		return false
	}

	// 比较长度
	v1 := reflect.ValueOf(map1)
	v2 := reflect.ValueOf(map2)
	if v1.Len() != v2.Len() {
		return false
	}

	// 比较键值对
	for _, key := range v1.MapKeys() {
		value1 := v1.MapIndex(key)
		value2 := v2.MapIndex(key)

		if !value2.IsValid() || !reflect.DeepEqual(value1.Interface(), value2.Interface()) {
			return false
		}
	}

	return true
}

func getProperties(mapping map[string]interface{}) (map[string]interface{}, error) {
	if mappingData, ok := mapping[migrate.UserDataMapping.Name]; ok {
		if tmpData, ok := mappingData.(map[string]interface{}); ok {
			if tmpData, ok := tmpData["mappings"]; ok {
				if tmpData, ok := tmpData.(map[string]interface{}); ok {
					if tmpData, ok := tmpData["properties"]; ok {
						if properties, ok := tmpData.(map[string]interface{}); ok {
							return properties, nil
						}
					}
				}

			}

		}

	}

	return nil, fmt.Errorf("properties not found")
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
		mapping, err := client.GetMapping().Index(migrate.UserDataMapping.Name).Do(ctx)
		if err != nil {
			t.Fatal(err)
			return
		}
		properties, err := getProperties(mapping)
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(properties, migrate.UserDataMapping.Properties) {
			// 先进行备份
			_, err = client.Reindex().Body(migrate.UserDataMapping.BackupBody()).Do(ctx)
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
