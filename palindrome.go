package main

import (
	"fmt"

	"github.com/FS-Frost/palindrome"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type palindromeCheker struct {
	app.Compo
	input  string
	output string
}

func (p *palindromeCheker) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Verificador de palíndromos"),
		app.Input().
			Type("text").
			Value(p.input).
			Placeholder("Ingresa una palabra...").
			AutoFocus(true).
			OnChange(p.OnInputChange),
		app.Br(),
		app.Br(),
		app.Text(p.output),
	)
}

func (p *palindromeCheker) OnInputChange(ctx app.Context, e app.Event) {
	p.input = ctx.JSSrc().Get("value").String()
	isPalindrome := palindrome.IsPalindrome(p.input)
	var result string

	if isPalindrome {
		result = "es palíndromo :)"
	} else {
		result = "no es palíndromo :/"
	}

	p.output = fmt.Sprintf("\"%s\" %s", p.input, result)
}
