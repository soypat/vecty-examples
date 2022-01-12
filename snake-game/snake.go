package main

import (
	"math/rand"
	"time"

	"github.com/soypat/vecty-examples/util/jlog"
	"github.com/soypat/vecty-examples/util/svg"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	_ "github.com/hexops/vecty/event"
)

func main() {
	vecty.SetTitle("GopherJS • TodoMVC")
	p := &Model{
		snek: []vec{
			{svgW / 2, svgH / 2},
			{svgW / 2, svgH/2 + 1},
			{svgW / 2, svgH/2 + 2},
		},
		food: vec{svgW / 4, svgH / 4},
		msgs: make(chan msg),
	}

	go func() {
		// Update every second.
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			p.msgs <- msg{}
		}
	}()

	vecty.RenderInto("body", p)
	for a := range p.msgs {
		p.update(a)
	}
}

const (
	svgPix = 6
	svgH   = 20
	svgW   = 30
)

type direction int

const (
	dirNone direction = iota
	dirUp
	dirDown
	dirLeft
	dirRight
)

type msg struct {
	dir direction
}

type Model struct {
	vecty.Core

	state gamestate
	snek  []vec
	food  vec
	msgs  chan msg
}
type vec struct {
	x, y int
}

type gamestate int

const (
	play gamestate = iota
	lose
)

func (m *Model) update(action msg) {
	var facing direction
	// Set current direction if non-keydown update.
	switch {
	case m.state == lose:
		//Do nothing after loss
		return
	case m.snek[0].x == m.snek[1].x && m.snek[0].y > m.snek[1].y:
		// snake head below body
		facing = dirDown

	case m.snek[0].x == m.snek[1].x && m.snek[0].y < m.snek[1].y:
		// snake head above body
		facing = dirUp

	case m.snek[0].y == m.snek[1].y && m.snek[0].x < m.snek[1].x:
		// snake head left of body
		facing = dirLeft
	case m.snek[0].y == m.snek[1].y && m.snek[0].x > m.snek[1].x:
		// snake head right of body
		facing = dirRight

	default:
		panic("impossible position")
	}

	if action.dir == dirNone {
		action.dir = facing
	}

	m.state = play
	head := m.snek[0]
	switch action.dir {
	// TODO(soypat): need to prevent going in opposite direction of facing dir.
	case dirDown:
		m.snek = append([]vec{{head.x, head.y + 1}}, m.snek...)

	case dirUp:
		m.snek = append([]vec{{head.x, head.y - 1}}, m.snek...)

	case dirLeft:
		m.snek = append([]vec{{head.x - 1, head.y}}, m.snek...)

	case dirRight:
		m.snek = append([]vec{{head.x + 1, head.y}}, m.snek...)
	}
	newHead := m.snek[0]
	for i, v := range m.snek[1:] {
		if v.x == newHead.x && v.y == newHead.y {
			jlog.Debug("lose bodypart/vec", i, v, "to head position", newHead)
			m.state = lose
		}
	}
	if newHead == m.food {
		m.food = vec{rand.Intn(svgW), rand.Intn(svgH)}
		jlog.Debug("new food at", m.food)
	} else {
		m.snek = m.snek[:len(m.snek)-1]
	}

	jlog.Debug(m.snek)
	vecty.Rerender(m)
}

func (m *Model) Render() vecty.ComponentOrHTML {
	s := &svg.SVG{
		Height: svgH * svgPix,
		Width:  svgW * svgPix,
	}
	food := svg.NewRect("blue", m.food.x*svgPix, m.food.y*svgPix, svgPix, svgPix)
	s.Add(food)
	for i := range m.snek {
		r := svg.NewRect("red", m.snek[i].x*svgPix, m.snek[i].y*svgPix, svgPix, svgPix)
		s.Add(r)
	}
	jlog.Debug(s)
	return elem.Body(
		vecty.If(m.state == play, elem.Paragraph(
			vecty.Markup(vecty.UnsafeHTML("Use Arrow keys to move.")),
		)),
		vecty.If(m.state == lose, elem.Paragraph(
			vecty.Markup(vecty.UnsafeHTML("You have lost.")),
		)),
		vecty.Markup(
			event.KeyUp(func(e *vecty.Event) {
				action := msg{}
				switch e.Value.Get("key").String() {
				case "ArrowUp":
					action.dir = dirUp
				case "ArrowDown":
					action.dir = dirDown
				case "ArrowLeft":
					action.dir = dirLeft
				case "ArrowRight":
					action.dir = dirRight
				default:
					jlog.Debug("do nothing, invalid keydown")
					return
				}
				jlog.Debug("update ready")
				m.msgs <- action
			}),
		),
		elem.Div(
			s.Render(),
		),
	)
}
