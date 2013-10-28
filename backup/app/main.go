package main

import (
	"fmt"
	gameloop "github.com/GlenKelley/go-glutil/gameloop"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"

	"github.com/neurocase/physicstest"
	"time"
	"math"

)


var starttime = time.Now()
var elapsedtime = 0
var lastcollision = 0.0

var cannondeg = 0.0

var canRadian = 0.0
var canDegree = 0.0

var bulRadian = 0.0
var bulDegree = 0.0

var en = physicstest.Entity{0,-1,-1, 1,"blue", true}
var cannon = physicstest.Entity{180,-9.5,-9.5,2,"orange", true}
var bullet = physicstest.Entity{0,0,0,0.4,"red", true}


var wallA = physicstest.Line{-10,10,10,10,"blue"}
var wallB = physicstest.Line{10,10,10,-10,"blue"}
var wallC = physicstest.Line{-10,-10,10,-10,"blue"}
var wallD = physicstest.Line{-10,10,-10,-10,"blue"}
//var wallC = physicstest.Line{-4,-4,4,4,"red"}

var BumpX = -9.5
var BumpY = -9.5

var worldObj = physicstest.SizeableTri{-2,-2, 0,4, 4,-2, "lblue" }

var bulletProjDist = 0.0
var bulletSpeed = 0.2

var winHeight = 480.0
var winWidth = 480.0

var mousex = 0.0
var mousey = 0.0


var worldmousex = 0.0
var worldmousey = 0.0

func ResolutionToWorld(reswidth, resheight float64)(float64, float64){
ex := (reswidth * 2 / winWidth - 1)*10
wy := (1- resheight * 2 / winHeight)*10
return ex, wy
}

func fireCannon(){
	bulletProjDist = 0.0
	bulRadian = canRadian
	BumpX = -9.5
	BumpY = -9.5
}

func DetCollision(wall physicstest.Line){
		//boundary
		bnd := 0.25
			//A MUST be higher than B

		if wall.Bx ==  wall.Ax{
			if bullet.Ypos > wall.By-bnd && bullet.Ypos < wall.Ay+bnd{
				if math.Abs(bullet.Xpos - wall.Bx) <0.2{
					//fmt.Println("Collision Detected")
					CalcBounce(0)
				}
			}
		}else{

			if bullet.Xpos < wall.Bx-bnd && bullet.Xpos > wall.Ax+bnd{
				if math.Abs(bullet.Ypos - wall.By) <0.2{
					//fmt.Println("Collision Detected")
					CalcBounce(3.14159265)
				}
			}
		}
}


func DegtoRad(deg float64)(float64){
	rad := deg*(math.Pi/180)
	return rad
}

func RadtoDeg(rad float64)(float64){
	deg := rad*(180/math.Pi)
	return deg
}

func CheckHighest(a, b float64)(float64, float64){
	if b < a{
		h := b
		b = a
		a = h
	}
	return a,b
}

func DetCollisionLine(max,may,mbx,mby float64)(bool){

	CheckHighest(max, mbx)
	CheckHighest(may, mby)

	if bullet.Ypos > may  &&	bullet.Ypos < mby{
		if bullet.Xpos > max  &&	bullet.Xpos < mbx{
		
			return true
		}
	}
	return false
}

func GetTriValues(tri physicstest.SizeableTri)(float64,float64,float64,float64,float64,float64){
	
	Ax := tri.Ax
	Ay := tri.Ay

	Bx := tri.Bx
	By := tri.By

	Cx := tri.Cx
	Cy := tri.Cy

	return Ax,Ay,Bx,By,Cx,Cy
}

func Ignore(a float64){
	if a - a == 254{
		fmt.Println("just fuck off and ignore the unused variable when I'm testing shit head")
	}
}

func DetCollisionTri(tri physicstest.SizeableTri){

	//Ax, Ay
	//Bx, By
	//Cx, Cy 

	Ax, Ay, Bx, By, Cx, Cy := GetTriValues(tri)

	Ignore(Cx)
	Ignore(Cy)

	an := GiveAngle(Ax,Ay,Bx,By)
	if DetCollisionLine(Ax,Ay,Bx,By){
			
			an = DegtoRad(an)
			fmt.Println("Line Collision Detected", an)
			CalcBounce(an)
	}

	/*an = GiveAngle(Bx,By,Cx,Cy)
	if DetCollisionLine(Bx,By,Cx,Cy){
			fmt.Println("Line Collision Detected")
			an = DegtoRad(an)
			CalcBounce(an)
	}


	an = GiveAngle(Ax,Ay,Cx,Cy)
	if DetCollisionLine(Ax,Ay,Cx,Cy){
			fmt.Println("Line Collision Detected")
			an = DegtoRad(an)
			CalcBounce(an)
	}	*/

}


func GiveAngle(Ax,Ay,Bx,By float64)(float64){

	CheckHighest(Ax, Bx)
	CheckHighest(Ay, By)
	Rx := Bx - Ax
	Ry := By - Ay
	return math.Atan2(Rx,Ry)
}



var elapsed = 0.0

func CalcBounce(an float64){

	if (elapsed - lastcollision) > 188888888.0 {

	BumpX = bullet.Xpos
	BumpY = bullet.Ypos
	fmt.Println(bulRadian)
	bulRadian = 3.14159265-bulRadian+an
	bulletProjDist = 0.0
 	lastcollision = elapsed

 	}	
}

func main() {
	game := &Game{}
	err := gameloop.CreateWindow(int(winWidth), int(winHeight), "daz gl test", false, game)
	fmt.Println(err)
}

type Game struct {
//	Red float64
}

func (game *Game) Init(window *glfw.Window) {
	//Select the 'projection matrix'
	gl.MatrixMode(gl.PROJECTION)
	//Reset
	gl.LoadIdentity()
	//Scale everything down, to 1/10 scale
	gl.Scaled(0.1,0.1,0.1)
}

func (game *Game) Draw(window *glfw.Window) {

	nowtime := time.Now()
	et := nowtime.Sub(starttime)
	elapsed =  float64(et)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.ClearColor(0.2, 0.2, 0.2, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	physicstest.DrawSizeableTri(worldObj)


	physicstest.DrawLine(wallA)
	physicstest.DrawLine(wallB)
	physicstest.DrawLine(wallC)
	physicstest.DrawLine(wallD)
	// UNUSED CODE MAKES NON-EXISTANT ENEMIES FACE PLAYER

		hx := worldmousex - cannon.Xpos 
		hy := worldmousey - cannon.Ypos
		hhyp := (hx * hx) + (hy * hy)
		hhyp = math.Sqrt(hhyp)
		hsin := hy / hhyp
		canRadian = math.Asin(hsin)

		canDegree = canRadian*180/math.Pi

		if worldmousex < cannon.Xpos{
		canDegree = 0 - canDegree -90
		}else{
		canDegree = canDegree + 90	
		}
		if canDegree > 360{
			canDegree -= 360
		}else if canDegree < 0{
			canDegree += 360
		}

		/*if math.Abs(hx) < 4 && math.Abs(hy) < 4{

			//activate bullet
			//bullet := physicstest.Entity{degree,entary[i].Xpos,entary[i].Ypos,0.2,"grey", true}
			//bulletary = append(bulletary,bullet)
			forceX := bulletSpeed * cos(rot)
			forceY := bulletSpeed * sin(rot) 
		}*/
		bulletProjDist += bulletSpeed
		bullet.Xpos = BumpX + (math.Cos(bulRadian) * bulletProjDist)
		bullet.Ypos = BumpY + (math.Sin(bulRadian) * bulletProjDist)
		
		DetCollision(wallA)
		DetCollision(wallB)
		DetCollision(wallC)
		DetCollision(wallD)
		DetCollisionTri(worldObj)

		cannon.Rot = canDegree
		cannondeg = canDegree

	//bullet = physicstest.Entity{0,0,0,0.3,"grey", true}

	en = physicstest.Entity{225,worldmousex+0.2, worldmousey-0.2, 0.45,"yellow", true}
	physicstest.DrawEntity(en)
	physicstest.DrawEntity(cannon)
	physicstest.DrawEntity(bullet)

}

func (game *Game) Reshape(window *glfw.Window, width, height int) {
winHeight = float64(height)
winWidth = float64(width)
fmt.Println("reshape does not work", float64(width), float64(height))
}

func (game *Game) MouseClick(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if action == glfw.Press{

		switch button{

		case glfw.MouseButtonLeft:
			fireCannon()

		case glfw.MouseButtonRight:
			fmt.Println("candeg",cannondeg,"bpoj",bulletProjDist)

		}
	}
}

func (game *Game) MouseMove(window *glfw.Window, xpos float64, ypos float64) {
	mousex = xpos
	mousey = ypos
	worldmousex, worldmousey = ResolutionToWorld(xpos, ypos)
}

func (game *Game) KeyPress(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	fmt.Println("keypress", k)
	if action == glfw.Release {
		switch k {

		case 77:
			fmt.Println("m")
		case glfw.KeyEscape:
			window.SetShouldClose(true)
		}
	}
}

func (game *Game) Scroll(window *glfw.Window, xoff float64, yoff float64) {

}

func (game *Game) Simulate(time gameloop.GameTime) {
	//game.Red = math.Sin(time.Elapsed.Seconds())
}

func (game *Game) OnClose(window *glfw.Window) {

}

func (game *Game) IsIdle() bool {
	//if idle is true, the gameloop will
	//wait for user input before drawing
	return false
}

func (game *Game) NeedsRender() bool {
	//if render is false the game will not redraw the screen
	return true
}