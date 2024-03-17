package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds application configuration data
type AppConfig struct {
	UseCache 						bool
	TemplateCache 			map[string]*template.Template
	InfoLog 						*log.Logger
	InProd							bool
	Session							*scs.SessionManager
}