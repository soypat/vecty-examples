package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type Page struct {
	vecty.Core
	Title       string
	Description string
	Content     string
}

func (cp *Page) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("demo-content-transition"),
		),
		elem.Section(
			vecty.Markup(
				vecty.Class("component-catalog-panel"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("component-catalog-panel__hero-area"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("component-catalog-panel__header"),
					),
					elem.Heading1(
						vecty.Markup(
							vecty.Class(
								"component-catalog-panel__header-elements", "mdc-typography--headline3",
							),
						),
						vecty.Text(cp.Title),
					),
					elem.Heading6(
						vecty.Markup(
							vecty.Class(
								"component-catalog-panel__header-elements", "mdc-typography--headline6",
							),
						),
						vecty.Text(cp.Description),
					),
					elem.Paragraph(
						vecty.Markup(
							vecty.Class(
								"component-catalog-panel__header-elements", "mdc-typography--body1",
							),
						),
						vecty.Text(cp.Content),
					),
				),
			),
		),
	)
}
