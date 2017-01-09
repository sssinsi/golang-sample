package main

import (
	"html/template"
	"log"
	"os"
)

// const issueList = template.Must(template.New("issuelist").Parse(`
// <h1>{{.TotalCount}} issues</h1>
// <table>
// <tr style='text-align: left'>
//   <th>#</th>
//   <th>State</th>
//   <th>User</th>
//   <th>Title</th>
// </tr>
// {{range .Items}}
// <tr>
//   <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
//   <td>{{.State}}}</td>
//   <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
//   <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
// </tr>
// {{end}}
// </table>
// `))

func main() {
	const templ = `<p>A:{{.A}}</p><p>B{{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        //信頼さないプレーンテキスト
		B template.HTML //信頼されたHTML
	}

	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
