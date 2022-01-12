package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/vecty-examples/basic-material-ui/components"
	"github.com/soypat/vecty-examples/basic-material-ui/views"
	"github.com/soypat/vecty-examples/util/jlog"
	"github.com/vecty-components/material/app"
	"github.com/vecty-components/material/base"
	router "marwan.io/vecty-router"
)

func main() {
	base.SetViewport()
	base.AddStyles()

	// Add custom styles here

	base.Boot()
	body := &Body{}
	vecty.RenderBody(body)
}

type Body struct {
	vecty.Core
}

// Render renders the <body> tag with the App as its children
func (b *Body) Render() vecty.ComponentOrHTML {
	jlog.Trace("Render Body")
	return elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
		),
		&CatalogPage{},
	)
}

type CatalogPage struct {
	vecty.Core
}

func (c *CatalogPage) Render() vecty.ComponentOrHTML {
	jlog.Trace("Render Catalog")
	vecty.SetTitle("Material Components Web | Catalog")

	links := []components.DemoLink{
		{
			Name:  "Top 5 anime betrayals",
			URL:   "/example",
			Image: "/favico.png",
			Page: components.Page{
				Title:       "These anime betrayals are not the ones you'd expect!",
				Description: "Betrayals are bittersweet, like marmite",
				Content:     loremipsum,
			},
		},
	}

	sidebar := components.NewComponentSidebar(links)
	a := &app.A{
		RootMarkup: vecty.Markup(
			vecty.Class("panel"),
		),
		ChildMarkup: vecty.Markup(
			vecty.Class("content"),
		),
		Sidebar: sidebar,
		Appbar:  components.NewHeaderBar(sidebar),
		Routes:  genRoutes(links),
	}
	return a
}

func genRoutes(links []components.DemoLink) []vecty.ComponentOrHTML {
	routeOpt := router.NewRouteOpts{ExactMatch: true}
	// Create main site link
	routes := []vecty.ComponentOrHTML{
		router.NewRoute("/", views.NewComponentImageList(links), routeOpt),
	}
	for _, l := range links {
		routes = append(routes, router.NewRoute(l.URL, &l.Page, routeOpt))
	}
	return routes
}

const loremipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce sodales ante nec pellentesque viverra. Donec tempor ac risus vel maximus. Phasellus id vulputate eros. Duis quis pharetra massa. Phasellus nec nisi id odio convallis interdum nec eu tellus. Vestibulum rhoncus ex ut interdum fermentum. Nunc ornare ultricies varius. Praesent vestibulum tellus eu nulla ultrices, in ultrices turpis sodales. Nam consectetur metus purus, nec varius purus imperdiet eget. Duis urna magna, dictum sed placerat a, bibendum porta nulla. Proin eget lacus odio. Ut id dolor eu risus scelerisque vulputate eget et sem. Quisque gravida tincidunt lacus, vel porta est vulputate sed. In accumsan nisl sagittis ante ornare bibendum. Nam rhoncus quam a libero euismod, vitae rutrum orci molestie.

Etiam consequat est ligula, eget tempor eros dapibus sollicitudin. Fusce hendrerit, arcu id pretium tincidunt, enim est convallis enim, vel vestibulum dolor erat a tortor. Aenean commodo nisi est, a venenatis dui viverra eu. Aliquam elementum bibendum eleifend. Maecenas posuere felis et aliquet rutrum. Cras fringilla tortor turpis. Vestibulum consectetur enim porta mauris tempor, quis vestibulum diam suscipit. Vivamus porta arcu pharetra maximus tincidunt. Mauris condimentum posuere ligula et venenatis. Phasellus at mi sed velit dictum pellentesque eu a lacus. Aenean sit amet odio ultrices, hendrerit lectus eget, fermentum enim. Nullam nec tincidunt lorem. Lorem ipsum dolor sit amet, consectetur adipiscing elit.

Cras ut scelerisque libero, facilisis mattis dolor. Phasellus eu scelerisque erat. Praesent ante dolor, suscipit vel placerat nec, dignissim sit amet massa. Donec magna odio, vehicula vel facilisis malesuada, porttitor ac odio. In in nisi turpis. Pellentesque viverra fringilla ex at lobortis. Duis mi metus, suscipit vel lacus et, blandit scelerisque diam. Vivamus pellentesque semper nisi nec vestibulum. Maecenas ultrices arcu id pulvinar facilisis. Curabitur vitae ante posuere, convallis risus eget, eleifend tellus. Morbi a mauris in nisi finibus luctus at et arcu.

Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Nulla facilisi. Pellentesque ultricies elit quis quam dapibus sodales. Cras pellentesque mauris a congue dignissim. Proin vitae congue risus. Curabitur nec mollis lorem. In feugiat eros a libero tincidunt, et iaculis arcu finibus. Integer viverra accumsan sapien ac eleifend. Integer vitae mauris ut augue finibus maximus. Sed vel tellus egestas, viverra lacus et, tristique magna. Aenean elementum nec risus ac aliquet. Nullam ullamcorper blandit diam vel aliquet. Nullam sagittis augue non nunc elementum, vel porta leo accumsan. Suspendisse potenti.

Maecenas lacinia, ante ac convallis mattis, nunc urna volutpat nisl, ut dictum tellus dui a odio. Maecenas vitae suscipit dolor. Aliquam id venenatis sapien. Sed nec pretium dui. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Ut at efficitur lorem. Quisque vitae lacus nisl. Phasellus nunc felis, suscipit vel commodo in, mattis quis arcu. In placerat fermentum varius. Mauris sagittis, erat at malesuada hendrerit, felis sapien finibus ipsum, non pretium eros libero sit amet nisi. Mauris vulputate lectus dignissim diam posuere, quis efficitur tellus rhoncus. Maecenas malesuada fringilla diam, id ornare turpis semper at.`
