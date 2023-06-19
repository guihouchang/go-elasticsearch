package schema

import (
	"entgo.io/ent"
	esfield "github.com/guihouchang/go-elasticsearch/schema/field"
)

type UserData struct {
	ent.Schema
}

func (UserData) Fields() []ent.Field {
	return []ent.Field{
		esfield.Text("text_tttt").Analyzer("aaaa").SearchAnalyzer("sss"),
		esfield.Keyword("keyword_kkkk"),
		esfield.Byte("byte_bbbb"),
		esfield.Short("short_ssss"),
		esfield.Int("int_iiii"),
		esfield.Long("long_llll"),
		esfield.Float("float_ffff"),
		esfield.Double("double_ddddd"),
		esfield.Bool("bool_bbbb"),
		esfield.Date("date_dddd").Format("Y-m-d H:m:s"),
	}
}
