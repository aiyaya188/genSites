package controlers

/*
import (
	"html/template"
	"os"
)

//ParseTtemple 替换模版中的变量
func ParseTemple() {
	//替换变量
	//替换数组
	//title := `abcdefe`
	//keywords := `12345678`
	temple := `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>{{.Title}}</title>
    </head>
    <body>
        {{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
    </body>
</html>`
	t, _ := template.New("webpage").Parse(temple)

	// 定义传入到模板的数据，并在终端打印
	data := struct {
		Title string

		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}
	t.Execute(os.Stdout, data)
}
*/
