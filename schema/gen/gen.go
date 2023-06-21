package gen

import (
	"bytes"
	"embed"
	_ "embed"
	"entgo.io/ent/entc/load"
	entfield "entgo.io/ent/schema/field"
	"fmt"
	"github.com/go-openapi/inflect"
	"github.com/guihouchang/go-elasticsearch/schema/field"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"
)

//go:embed template/const.tmpl template/struct.tmpl template/migrate.tmpl template/client.tmpl
var templateFile embed.FS

var (
	acronyms = make(map[string]struct{})
	rules    = ruleset()
)

var tpl, err = template.New("").Funcs(template.FuncMap{
	"ToCamelCase":  pascal,
	"ToUnderscore": ToUnderscore,
	"Type":         Type,
	"ESType":       ESType,
	"ToLower":      strings.ToLower,
	"RendProperty": RendProperty,
}).ParseFS(templateFile, "template/*.tmpl")

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
	case entfield.TypeJSON:

	}

	return nil
}

func ESType(t entfield.Type) string {
	return field.Type(t).String()
}

func Type(fd *load.Field) string {
	t := fd.Info.Type
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
		switch t {
		case entfield.TypeJSON:
			// json 格式需要单独处理
			return fd.Info.Ident
		}

		return t.String()
	}

	return "string"
}

func ruleset() *inflect.Ruleset {
	rules := inflect.NewDefaultRuleset()
	// Add common initialism from golint and more.
	for _, w := range []string{
		"ACL", "API", "ASCII", "AWS", "CPU", "CSS", "DNS", "EOF", "GB", "GUID",
		"HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "KB", "LHS", "MAC", "MB",
		"QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "SSO", "TCP",
		"TLS", "TTL", "UDP", "UI", "UID", "URI", "URL", "UTF8", "UUID", "VM",
		"XML", "XMPP", "XSRF", "XSS",
	} {
		acronyms[w] = struct{}{}
		rules.AddAcronym(w)
	}
	return rules
}

func pascalWords(words []string) string {
	for i, w := range words {
		upper := strings.ToUpper(w)
		if _, ok := acronyms[upper]; ok {
			words[i] = upper
		} else {
			words[i] = rules.Capitalize(w)
		}
	}
	return strings.Join(words, "")
}

func isSeparator(r rune) bool {
	return r == '_' || r == '-' || unicode.IsSpace(r)
}

// pascal converts the given name into a PascalCase.
//
//	user_info 	=> UserInfo
//	full_name 	=> FullName
//	user_id   	=> UserID
//	full-admin	=> FullAdmin
func pascal(s string) string {
	words := strings.FieldsFunc(s, isSeparator)
	return pascalWords(words)
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

func Gen(path string) {
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
