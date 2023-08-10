package test

import (
	"entgo.io/ent/entc/load"
	entfield "entgo.io/ent/schema/field"
	"fmt"
	"github.com/guihouchang/go-elasticsearch/schema/field"
	"github.com/guihouchang/go-elasticsearch/schema/gen"
	"html/template"
	"math"
	"strings"
	"testing"
	"unicode"
)

var tpl, err = template.New("").Funcs(template.FuncMap{
	"ToCamelCase":  ToCamelCase,
	"ToUnderscore": ToUnderscore,
	"Type":         Type,
	"ESType":       ESType,
	"ToLower":      strings.ToLower,
	"RendProperty": RendProperty,
}).ParseFiles(
	"../schema/gen/template/const.tmpl",
	"../schema/gen/template/struct.tmpl",
	"../schema/gen/template/migrate.tmpl",
	"../schema/gen/template/client.tmpl",
)

func ToUnderscore(camelCase string) string {
	var result strings.Builder
	for i, r := range camelCase {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return strings.ToLower(result.String())
}

func ToCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})

	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}

func RendProperty(fd *load.Field) []template.HTML {
	switch fd.Info.Type {
	case entfield.TypeString:
		strPropertyList := make([]template.HTML, 0)
		if analyzer, ok := fd.SchemaType["analyzer"]; ok {
			strPropertyList = append(strPropertyList,
				template.HTML(fmt.Sprintf(`"%s":"%s"`, "analyzer", analyzer)))
		}

		if search, ok := fd.SchemaType["search_analyzer"]; ok {
			strPropertyList = append(strPropertyList,
				template.HTML(fmt.Sprintf(`"%s":"%s"`, "search_analyzer", search)))
		}

		return strPropertyList
	case entfield.TypeTime:
		strPropertyList := make([]template.HTML, 0)
		if format, ok := fd.SchemaType["format"]; ok {
			strPropertyList = append(strPropertyList,
				template.HTML(fmt.Sprintf(`"%s":"%s"`, "format", format)))
		}
		return strPropertyList
	}

	return nil
}

func ESType(t entfield.Type) string {
	return field.Type(t).String()
}

func Type(t entfield.Type) string {
	if t.String() == "invalid" {
		esType := field.Type(t)
		switch esType {
		case field.TypeKeyword:
			return "string"
		case field.TypeHalfFloat:
			return "float32"
		case field.TypeScaledFloat:
			return "double"
		}
	} else {
		return t.String()
	}

	return "string"
}

func TestFloat(t *testing.T) {
	total := 10000
	pageSize := 26

	totalPage := math.Ceil(float64(total) / float64(pageSize))
	fmt.Println(totalPage)
}

func Test_Load(t *testing.T) {

	path := "./schema"
	gen.Gen(path)

}
