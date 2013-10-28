package physicstest

type Entity struct{
Rot, Xpos, Ypos, Size float64
Colour string
IsAlive bool
} 

type SizeableTri struct{
	Ax, Ay, Bx, By, Cx, Cy float64
	Colour string
}

type Line struct{
	Ax,Ay,Bx,By float64
	Colour string
}


type GridLoc struct{
Xpos, Ypos float64
}



type Node struct{

Xpos, Ypos, Hval, Fval, Gval float64
IsOccupied bool

Parent []GridLoc
//Parent List
}

