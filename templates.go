package main

// HTMLTemplate required in all templates
type HTMLTemplate struct {
	Title string
	JSON  string
}

// TemplateHome included paths "login", "register", "about"...
const TemplateHome string = "templates/home.html"
