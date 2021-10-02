package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type hello struct {
	app.Compo

	name string // Field where the username is stored
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello "),
			app.Text(h.name), // The name field used in the title
		),

		// The input HTML element that get the username.
		app.Input().
			Value(h.name).             // The name field used as current input value
			OnChange(h.OnInputChange), // The event handler that will store the username
	)
}

func (h *hello) OnInputChange(ctx app.Context, e app.Event) {
	h.name = ctx.JSSrc().Get("value").String() // Name field is modified
}
