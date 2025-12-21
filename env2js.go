package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type CoreContext struct {
	PAPERMERGE__OCR__LANG_CODES            string
	PAPERMERGE__OCR__DEFAULT_LANG_CODE     string
	PAPERMERGE__OCR__AUTOMATIC             bool
	PAPERMERGE__AUTH__OIDC_LOGOUT_URL      string
	PAPERMERGE__AUTH__POST_LOGOUT_REDIRECT_URI string
}

var (
	templateFileName = flag.String("f", "", "Template file")
)

func main() {
	flag.Parse()

	if *templateFileName == "" {
		fmt.Println("Error: empty template file")
		os.Exit(1)
	}

	templ, err := template.ParseFiles(*templateFileName)
	if err != nil {
		fmt.Printf("Error while parsing template file: %v", err)
		os.Exit(1)
	}

	data := CoreContext{
		PAPERMERGE__OCR__LANG_CODES:                os.Getenv("PAPERMERGE__OCR__LANG_CODES"),
		PAPERMERGE__OCR__DEFAULT_LANG_CODE:         os.Getenv("PAPERMERGE__OCR__DEFAULT_LANG_CODE"),
		PAPERMERGE__OCR__AUTOMATIC:                 getBoolEnv("PAPERMERGE__OCR__AUTOMATIC"),
		PAPERMERGE__AUTH__OIDC_LOGOUT_URL:          os.Getenv("PAPERMERGE__AUTH__OIDC_LOGOUT_URL"),
		PAPERMERGE__AUTH__POST_LOGOUT_REDIRECT_URI: os.Getenv("PAPERMERGE__AUTH__POST_LOGOUT_REDIRECT_URI"),
	}

	templ.Execute(os.Stdout, data)
}

func getBoolEnv(key string) bool {
	v := os.Getenv(key)

	switch strings.ToLower(v) {
	case "", "false", "no":
		return false
	case "true", "yes":
		return true
	}

	return false
}
