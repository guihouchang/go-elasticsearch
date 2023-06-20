package test

import (
	"bytes"
	"entgo.io/ent/entc/load"
	entfield "entgo.io/ent/schema/field"
	"fmt"
	"github.com/guihouchang/go-elasticsearch/schema/field"
	"github.com/guihouchang/go-elasticsearch/schema/gen"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
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

func genClient() error {
	pkgName, err := os.Getwd()
	if err != nil {
		return err
	}

	moduleName, err := executeGoList()
	if err != nil {
		return err
	}

	pkgName = filepath.Base(pkgName)

	var b bytes.Buffer
	target := fmt.Sprintf("%s.go", "client")
	err = tpl.ExecuteTemplate(&b, "client.tmpl", struct {
		PkgName    string
		ModuleName string
	}{
		PkgName:    pkgName,
		ModuleName: moduleName,
	})
	if err != nil {
		return err
	}

	err = os.WriteFile(target, b.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func executeGoList() (string, error) {
	cmd := exec.Command("go", "list")
	cmd.Env = append(os.Environ(), "GO111MODULE=on")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	moduleName := strings.TrimSpace(string(output))
	return moduleName, nil
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

func Test_Load(t *testing.T) {

	path := "./schema"
	gen.Gen(path)

}

func getPackagePath(filePath string) (string, error) {
	fset := token.NewFileSet()

	astFile, err := parser.ParseFile(fset, filePath, nil, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}

	pkgPath := astFile.Name.Name

	return pkgPath, nil
}
