package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type CoreContext struct {
	PM_OIDC_LOGOUT_URL      string
	PM_OIDC_CLIENT_ID       string
	PM_POST_LOGOUT_REDIRECT_URI string
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
		PM_OIDC_LOGOUT_URL:          os.Getenv("PM_OIDC_LOGOUT_URL"),
		PM_OIDC_CLIENT_ID:           os.Getenv("PM_OIDC_CLIENT_ID"),
		PM_POST_LOGOUT_REDIRECT_URI: os.Getenv("PM_POST_LOGOUT_REDIRECT_URI"),
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
