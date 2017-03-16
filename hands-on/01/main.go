package main

import (
"net/http"
"html/template"
"io"
)

var tpl *template.Template

func init() {
tpl = template.must(template.ParseGlob("template/*"))
