package view

import (
	"html/template"
	"strings"
)

// Defina todas as funções necessárias
func Join(slice []string, sep string) string {
	return strings.Join(slice, sep)
}

func ContainsInt(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Crie o FuncMap com todas as funções
var funcMap = template.FuncMap{
	"join":        Join,
	"containsInt": ContainsInt,
}

// Parseie os templates uma vez
var Templates = template.Must(template.New("").Funcs(funcMap).ParseGlob("view/templates/*.html"))
