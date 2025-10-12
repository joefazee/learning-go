package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type EmailData struct {
	RecipientName string
	SenderName    string
	Subject       string
	Body          string
	Items         []string // demo a loop
	UnreadCount   int
}

func main() {

	fmt.Println("--- Text template example ---")

	emailTemplate := `
Subject: {{ .Subject }}

{{.Body}}

{{if .Items}}
   Related Items:
{{range .Items}}
	- {{.}}
{{end}}
{{end}}

{{if gt .UnreadCount 0}}
You have {{.UnreadCount}} unreads.
{{else}}
You have no messages
{{end}}


- Thanks
{{.SenderName}}
`
	tmpl, err := template.New("email-message").Parse(emailTemplate)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data := EmailData{
		RecipientName: "Alice",
		SenderName:    "Bob's Auto-Responder",
		Subject:       "Your Weekly Update",
		Body:          "Here is the update you requested. We hope you find it useful.",
		Items:         []string{"Report A", "Document B", "Summary C"},
		UnreadCount:   0,
	}

	var output strings.Builder

	err = tmpl.Execute(&output, data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(strings.ToUpper(output.String()))
}
