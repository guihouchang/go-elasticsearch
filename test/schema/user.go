package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"github.com/guihouchang/go-elasticsearch/schema/annotation"
	esfield "github.com/guihouchang/go-elasticsearch/schema/field"
)

type UserData struct {
	ent.Schema
}

func (UserData) Fields() []ent.Field {
	return []ent.Field{
		esfield.Text("id").Analyzer("ik_max_word").Comment("这是一个ID"),
		esfield.Keyword("keyword_kkkk").Comment("这是一个keyword"),
		esfield.Byte("byte_bbbb").Comment("这是一个byte"),
		esfield.Short("short_ssss").Comment("这是一个short"),
		esfield.Int("int_iiii").Comment("这是一个int"),
		esfield.Long("long_llll").Comment("这是一个long"),
		esfield.Float("float_ffff").Comment("这是一个Float"),
		esfield.Double("double_ddddd").Comment("这是一个double"),
		esfield.Bool("bool_bbbb").Comment("这是一个bool"),
		esfield.Date("date_dddd").Format("yyyy-MM-dd HH:mm:ss").Comment("这是一个date"),
		esfield.Strings("stings_sssss").Comment("这是一个strings"),
		esfield.Ints("ints_iiiiii").Comment("这是一个ints"),
		esfield.Int64s("int64s_iiiiii").Comment("这是一个int64s"),
		esfield.Floats("floats_llllll").Comment("这是一个floats"),
	}
}

// Annotations of the schema.
func (UserData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotation.Setting{Settings: map[string]interface{}{
			"number_of_shards":   1,
			"number_of_replicas": 2,
		}},
		entsql.Annotation{Table: "test_user_data"},
	}
}
