package model

import "strconv"

type Gif struct {
	Name     string
	Width    int
	Height   int
	Geometry string
}

func (g *Gif) SetGeometry() string {
	g.Geometry = strconv.Itoa(g.Width) + "x" + strconv.Itoa(g.Height)
	return g.Geometry
}

