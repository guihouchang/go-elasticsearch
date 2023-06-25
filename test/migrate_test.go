package test

import (
	"context"
	"fmt"
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

	datas, _, err := UnserizerUserData(result)
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
		elastic.SetURL("http://120.76.136.97:9200"),
		elastic.SetSniff(false),
	)

	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	sc := migrate.NewSchema(client)
	err = sc.Create(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
