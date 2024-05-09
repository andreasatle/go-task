package templates

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

type Template struct {
	Template *template.Template
	Execute func(*gin.Context, gin.H)
}
type Templates struct {
	Index Template
	Contact Template
	Resume Template
}
func CreateTemplates() Templates {

	templates := Templates{}
	templates.Index.Template = template.Must(template.New("index").ParseFiles("templates/navbar.tmpl", "templates/base.tmpl", "templates/index.tmpl"))
	templates.Index.Execute = func(c *gin.Context, data gin.H){
		templates.Index.Template.ExecuteTemplate(c.Writer, "index.tmpl", data)
	}

	templates.Resume.Template = template.Must(template.New("resume").ParseFiles("templates/navbar.tmpl", "templates/base.tmpl", "templates/resume.tmpl"))
	templates.Resume.Execute = func(c *gin.Context, data gin.H){
		templates.Resume.Template.ExecuteTemplate(c.Writer, "resume.tmpl", data)
	}

	templates.Contact.Template = template.Must(template.New("contact").ParseFiles("templates/navbar.tmpl", "templates/base.tmpl", "templates/contact.tmpl"))
	templates.Contact.Execute = func(c *gin.Context, data gin.H){
		templates.Contact.Template.ExecuteTemplate(c.Writer, "contact.tmpl", data)
	}

	return templates 
}