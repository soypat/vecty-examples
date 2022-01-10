package svg

import (
	"syscall/js"

	"github.com/hexops/vecty"
)

type SVG struct {
	vecty.Core
	Height, Width int
	things        []Obj
}

func log(a ...interface{}) {
	js.Global().Get("console").Call("log", a...)
}

func (s *SVG) Render() vecty.ComponentOrHTML {
	var childs []vecty.MarkupOrChild
	childs = append(childs, vecty.Markup(
		vecty.Namespace("http://www.w3.org/2000/svg"),
		vecty.Attribute("width", s.Width),
		vecty.Attribute("height", s.Height),
		vecty.Attribute("style", "border: thick solid black"),
	))
	for i := range s.things {
		childs = append(childs, s.things[i].HTML())
	}

	return vecty.Tag("svg", childs...)
}

func (s *SVG) Add(objs ...Obj) {
	s.things = append(s.things, objs...)
}

func NewRect(color string, coordx, coordy, height, width int) Rect {
	return Rect{
		color:  color,
		coordx: coordx,
		coordy: coordy,
		width:  width,
		height: height,
	}
}

type Obj interface {
	HTML() *vecty.HTML
}

type Rect struct {
	color          string
	coordx, coordy int
	width, height  int
}

func (o Rect) HTML() *vecty.HTML {
	return vecty.Tag("rect",
		vecty.Markup(
			vecty.Namespace("http://www.w3.org/2000/svg"),
			vecty.Attribute("x", o.coordx), vecty.Attribute("y", o.coordy),
			vecty.Attribute("width", o.width), vecty.Attribute("height", o.height),
			vecty.Attribute("style", "fill:"+o.color+";"),
		),
	)
}
