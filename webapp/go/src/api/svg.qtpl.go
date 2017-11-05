// This file is automatically generated by qtc from "svg.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line svg.qtpl:1
package main

//line svg.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line svg.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line svg.qtpl:3
func StreamSVG(qw422016 *qt422016.Writer, r *Room) {
	//line svg.qtpl:3
	qw422016.N().S(`<?xml version="1.0" standalone="no"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd"><svg xmlns="http://www.w3.org/2000/svg" version="1.1" baseProfile="full" width="`)
	//line svg.qtpl:3
	qw422016.N().D(r.CanvasWidth)
	//line svg.qtpl:3
	qw422016.N().S(`" height="`)
	//line svg.qtpl:3
	qw422016.N().D(r.CanvasHeight)
	//line svg.qtpl:3
	qw422016.N().S(`" style="width:`)
	//line svg.qtpl:3
	qw422016.N().D(r.CanvasWidth)
	//line svg.qtpl:3
	qw422016.N().S(`px;height:`)
	//line svg.qtpl:3
	qw422016.N().D(r.CanvasHeight)
	//line svg.qtpl:3
	qw422016.N().S(`px;background-color:white;" viewBox="0 0 `)
	//line svg.qtpl:3
	qw422016.N().D(r.CanvasWidth)
	//line svg.qtpl:3
	qw422016.N().S(` `)
	//line svg.qtpl:3
	qw422016.N().D(r.CanvasHeight)
	//line svg.qtpl:3
	qw422016.N().S(`">`)
	//line svg.qtpl:3
	for _, s := range r.Strokes {
		//line svg.qtpl:3
		qw422016.N().S(`<polyline id="`)
		//line svg.qtpl:3
		qw422016.E().V(s.ID)
		//line svg.qtpl:3
		qw422016.N().S(`" stroke="rgba(`)
		//line svg.qtpl:3
		qw422016.N().D(s.Red)
		//line svg.qtpl:3
		qw422016.N().S(`,`)
		//line svg.qtpl:3
		qw422016.N().D(s.Green)
		//line svg.qtpl:3
		qw422016.N().S(`,`)
		//line svg.qtpl:3
		qw422016.N().D(s.Blue)
		//line svg.qtpl:3
		qw422016.N().S(`,`)
		//line svg.qtpl:3
		qw422016.N().F(s.Alpha)
		//line svg.qtpl:3
		qw422016.N().S(`)" stroke-width="`)
		//line svg.qtpl:3
		qw422016.N().D(s.Width)
		//line svg.qtpl:3
		qw422016.N().S(`" stroke-linecap="round" stroke-linejoin="round" fill="none" points="`)
		//line svg.qtpl:3
		for _, p := range s.Points {
			//line svg.qtpl:3
			qw422016.N().F(p.X)
			//line svg.qtpl:3
			qw422016.N().S(`,`)
			//line svg.qtpl:3
			qw422016.N().F(p.Y)
			//line svg.qtpl:3
			qw422016.N().S(` `)
			//line svg.qtpl:3
		}
		//line svg.qtpl:3
		qw422016.N().S(`"></polyline>`)
		//line svg.qtpl:3
	}
	//line svg.qtpl:3
	qw422016.N().S(`</svg>`)
//line svg.qtpl:3
}

//line svg.qtpl:3
func WriteSVG(qq422016 qtio422016.Writer, r *Room) {
	//line svg.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line svg.qtpl:3
	StreamSVG(qw422016, r)
	//line svg.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line svg.qtpl:3
}

//line svg.qtpl:3
func SVG(r *Room) string {
	//line svg.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
	//line svg.qtpl:3
	WriteSVG(qb422016, r)
	//line svg.qtpl:3
	qs422016 := string(qb422016.B)
	//line svg.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
	//line svg.qtpl:3
	return qs422016
//line svg.qtpl:3
}
