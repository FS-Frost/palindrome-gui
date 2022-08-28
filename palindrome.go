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

func (p *palindromeCheker) OnAppUpdate(ctx app.Context) {
	ctx.Reload()
}

func (p *palindromeCheker) OnNav(ctx app.Context) {
	go startVersionChecking(ctx)
}

func (p *palindromeCheker) Render() app.UI {
	return app.Div().Class("main columns").Body(
		app.Div().Class("column").Body(
			app.H1().Class("title is-2").Text("Verificador de palíndromos"),
			app.Input().
				Class("input").
				Type("text").
				Placeholder("Ingresa una palabra...").
				AutoFocus(true).
				OnKeyup(p.OnKeyup),
			app.P().Text(p.output).Class("mt-3"),
			app.Div().Class("mt-3").Body(
				app.Text("Potenciado por: "),
				app.A().
					Text(LIB_URL).
					Href(LIB_URL).
					Target("_blank"),
			),
		),
	)
}

func (p *palindromeCheker) OnKeyup(ctx app.Context, e app.Event) {
	input := ctx.JSSrc().Get("value").String()
	if input == "" {
		p.output = ""
		return
	}

	isPalindrome := palindrome.IsPalindrome(input)
	result := "no es palíndromo :/"

	if isPalindrome {
		result = "es palíndromo :)"
	}

	p.output = fmt.Sprintf("\"%s\" %s", input, result)
}
