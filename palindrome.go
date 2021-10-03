package main

import (
	"fmt"

	"github.com/FS-Frost/palindrome"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	LIB_URL = "https://www.github.com/FS-Frost/palindrome"
)

type palindromeCheker struct {
	app.Compo
	output string
}

func (p *palindromeCheker) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Verificador de palíndromos"),
		app.Input().
			Type("text").
			Placeholder("Ingresa una palabra...").
			AutoFocus(true).
			OnKeyup(p.OnKeyUp),
		app.Br(),
		app.Br(),
		app.Text(p.output),
		app.Br(),
		app.Br(),
		app.Div().Body(
			app.Text("Potenciado por: "),
			app.A().
				Text(LIB_URL).
				Href(LIB_URL),
		),
	)
}

func (p *palindromeCheker) OnKeyUp(ctx app.Context, e app.Event) {
	input := ctx.JSSrc().Get("value").String()

	if input == "" {
		p.output = ""
		return
	}

	isPalindrome := palindrome.IsPalindrome(input)
	var result string

	if isPalindrome {
		result = "es palíndromo :)"
	} else {
		result = "no es palíndromo :/"
	}

	p.output = fmt.Sprintf("\"%s\" %s", input, result)
}
