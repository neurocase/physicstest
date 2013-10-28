package physicstest

import (
//	"fmt"
	"github.com/go-gl/gl"
	//"math"

)


func Hello() string{

	return "drawengine"
}

func DrawLine(Ln Line){

ax := Ln.Ax
ay := Ln.Ay

bx := Ln.Bx
by := Ln.By

col := Ln.Colour

switch col{
	case "grey":
			gl.Color3d(0.5,0.5,0.5)
	case "red":
			gl.Color3d(1, 0, 0)
	case "green":
			gl.Color3d(0, 1, 0)
	case "blue":
			gl.Color3d(0, 0, 1)
	case "lblue":
		gl.Color3d(0.3, 0.3, 1)
	case "orange":
		gl.Color3d(1, 0.5, 0)
	case "yellow":
			gl.Color3d(1, 1, 0)
	case "purple":
			gl.Color3d(1, 0, 1)
	default:
			gl.Color3d(1, 1, 1)
}
	if ay != by{
		gl.PushMatrix()
		gl.Begin(gl.TRIANGLES)
		gl.Vertex3d(ax-0.2, ay, 0)
		gl.Vertex3d(ax+0.2, ay, 0)
		gl.Vertex3d(ax-0.2, by, 0)
		gl.End()
		gl.PopMatrix()

		gl.PushMatrix()
		gl.Begin(gl.TRIANGLES)
		gl.Vertex3d(ax+0.2, ay, 0)
		gl.Vertex3d(ax+0.2, by, 0)
		gl.Vertex3d(ax-0.2, by, 0)
		gl.End()
		gl.PopMatrix()
	}else{

		gl.PushMatrix()
		gl.Begin(gl.TRIANGLES)
		gl.Vertex3d(ax, ay+0.2, 0)
		gl.Vertex3d(bx, by+0.2, 0)
		gl.Vertex3d(bx, by-0.2, 0)
		gl.End()
		gl.PopMatrix()

		gl.PushMatrix()
		gl.Begin(gl.TRIANGLES)
		gl.Vertex3d(ax, ay+0.2, 0)
		gl.Vertex3d(ax, ay-0.2, 0)
		gl.Vertex3d(bx, by-0.2, 0)
		gl.End()
		gl.PopMatrix()

	}


}



func DrawSizeableTri(Tr SizeableTri){
ax := Tr.Ax
bx := Tr.Bx
cx := Tr.Cx

ay := Tr.Ay
by := Tr.By
cy := Tr.Cy

col := Tr.Colour

switch col{
	case "grey":
			gl.Color3d(0.5,0.5,0.5)
	case "red":
			gl.Color3d(1, 0, 0)
	case "green":
			gl.Color3d(0, 1, 0)
	case "blue":
			gl.Color3d(0, 0, 1)
	case "lblue":
		gl.Color3d(0.3, 0.3, 1)
	case "orange":
		gl.Color3d(1, 0.5, 0)
	case "yellow":
			gl.Color3d(1, 1, 0)
	case "purple":
			gl.Color3d(1, 0, 1)
	default:
			gl.Color3d(1, 1, 1)
}

	gl.PushMatrix()

//	gl.Translated(x, y, 0)
//	gl.Scaled(s, s, s)

//	gl.Rotated(r, 0, 0, 1)

	gl.Begin(gl.TRIANGLES)
	gl.Vertex3d(ax, ay, 0)
	gl.Vertex3d(bx, by, 0)
	gl.Vertex3d(cx, cy, 0)
	gl.End()

	gl.PopMatrix()

}

func DrawEntity(en Entity){

x := en.Xpos
y := en.Ypos
r := en.Rot
s := en.Size
col := en.Colour


switch col{
	case "grey":
			gl.Color3d(0.5,0.5,0.5)
	case "red":
			gl.Color3d(1, 0, 0)
	case "green":
			gl.Color3d(0, 1, 0)
	case "blue":
			gl.Color3d(0, 0, 1)
	case "lblue":
		gl.Color3d(0.3, 0.3, 1)
	case "orange":
		gl.Color3d(1, 0.5, 0)
	case "yellow":
			gl.Color3d(1, 1, 0)
	case "purple":
			gl.Color3d(1, 0, 1)
	default:
			gl.Color3d(1, 1, 1)
}

	gl.PushMatrix()

	gl.Translated(x, y, 0)
	gl.Scaled(s, s, s)
	//gl.Rotated(r*180/math.Pi, 0, 0, 1)
	gl.Rotated(r, 0, 0, 1)

	enx := [3]float64{-0.8, 0.8, 0}
	eny := [3]float64{1, 1, -1}


	//in OpenGL 3.2, the vertices below would be stored on the GPU
	//so all you would need to do is say DrawShape(A) or something

	gl.Begin(gl.TRIANGLES)
	gl.Vertex3d(enx[0], eny[0], 0)
	gl.Vertex3d(enx[1], eny[1], 0)
	gl.Vertex3d(enx[2], eny[2], 0)
	gl.End()

	gl.PopMatrix()

}