package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"github.com/guihouchang/go-elasticsearch/schema/annotation"
	esfield "github.com/guihouchang/go-elasticsearch/schema/field"
)

type UserData struct {
	ent.Schema
}

func (UserData) Fields() []ent.Field {
	return []ent.Field{
		esfield.Text("id").Analyzer("ik_max_word").SearchAnalyzer("ik_max_word"),
		esfield.Keyword("keyword_kkkk"),
		esfield.Byte("byte_bbbb"),
		esfield.Short("short_ssss"),
		esfield.Int("int_iiii"),
		esfield.Long("long_llll"),
		esfield.Float("float_ffff"),
		esfield.Double("double_ddddd"),
		esfield.Bool("bool_bbbb"),
		esfield.Date("date_dddd").Format("yyyy-MM-dd HH:mm:ss"),
	}
}

// Annotations of the schema.
func (UserData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotation.Setting{Settings: map[string]interface{}{
			"number_of_shards":   1,
			"number_of_replicas": 2,
		}},
	}
}
