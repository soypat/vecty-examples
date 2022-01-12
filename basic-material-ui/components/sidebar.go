package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/drawer"
	"github.com/vecty-components/material/ul"
	router "marwan.io/vecty-router"
)

type DemoLink struct {
	Name  string
	URL   string
	Image string
	Page  Page
}

func NewComponentSidebar(list []DemoLink) *drawer.D {
	links := append([]DemoLink{
		{
			Name: "Home",
			URL:  "/",
		},
	}, list...)

	items := make([]vecty.ComponentOrHTML, len(links))
	for i, link := range links {
		items[i] = &ul.Item{
			Primary: router.Link(
				link.URL,
				link.Name,
				router.LinkOptions{},
			),
		}
	}

	return &drawer.D{
		Root: vecty.Markup(
			prop.ID("drawer"),
			vecty.Class("drawer", "mdc-top-app-bar--fixed-adjust"),
		),
		Type: drawer.Dismissible,
		Content: &ul.L{
			Items: items,
		},
	}
}
