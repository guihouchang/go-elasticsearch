package generate

import (
	"bytes"
	"entgo.io/ent/entc/load"
	entfield "entgo.io/ent/schema/field"
	"fmt"
	"github.com/guihouchang/go-elasticsearch/schema/field"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

var GenCmd = &cobra.Command{
	Use:     "generate [flags] path",
	Short:   "generate go code for the schema directory",
	Example: "ent-elastic generate  ./es/schema",
	Run:     Gen,
}

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

func HtmlUnescapeString(s string) template.HTML {
	return template.HTML(s)
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

func genClient() error {
	pkgName, err := os.Getwd()
	if err != nil {
		return err
	}

	pkgName = filepath.Base(pkgName)
	var b bytes.Buffer
	target := fmt.Sprintf("%s.go", "client")
	err = tpl.ExecuteTemplate(&b, "client.tmpl", struct {
		PkgName string
	}{
		PkgName: pkgName,
	})
	if err != nil {
		return err
	}

	return os.WriteFile(target, b.Bytes(), os.ModePerm)
}

func genMigrate(scs []*load.Schema) error {
	// 创建migrate目录
	err = os.MkdirAll("migrate", os.ModePerm)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	target := fmt.Sprintf("migrate/%s.go", "schema")
	err = tpl.ExecuteTemplate(&b, "migrate.tmpl", scs)
	if err != nil {
		return err
	}

	return os.WriteFile(target, b.Bytes(), os.ModePerm)
}

func genStruct(sc *load.Schema) error {
	pkgName, err := os.Getwd()
	if err != nil {
		return err
	}

	pkgName = filepath.Base(pkgName)

	var b bytes.Buffer
	target := fmt.Sprintf("%s.go", strings.ToLower(sc.Name))
	err = tpl.ExecuteTemplate(&b, "struct.tmpl", struct {
		PkgName string
		Schema  *load.Schema
	}{
		PkgName: pkgName,
		Schema:  sc,
	})
	if err != nil {
		return err
	}

	return os.WriteFile(target, b.Bytes(), os.ModePerm)
}

func genConst(sc *load.Schema) error {
	name := strings.ToLower(sc.Name)
	err = os.MkdirAll(name, os.ModePerm)
	if err != nil {
		return err
	}
	var b bytes.Buffer

	target := fmt.Sprintf("%s/%s.go", name, name)
	err := tpl.ExecuteTemplate(&b, "const.tmpl", sc)
	if err != nil {
		return err
	}

	return os.WriteFile(target, b.Bytes(), os.ModePerm)
}

func Gen(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalln(fmt.Errorf("schema error for args"))
	}

	path := args[0]
	template.Must(tpl, err)
	// 拿到定义描述信息，加载模板生成代码
	spec, err := (&load.Config{Path: path, BuildFlags: []string{}}).Load()
	if err != nil {
		log.Fatal(err)
	}

	//packageName := filepath.Base(spec.PkgPath)
	for _, sc := range spec.Schemas {
		err := genConst(sc)
		if err != nil {
			log.Fatal(err)
		}

		err = genStruct(sc)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = genMigrate(spec.Schemas)
	if err != nil {
		log.Fatal(err)
	}

	err = genClient()
	if err != nil {
		log.Fatal(err)
	}
}
