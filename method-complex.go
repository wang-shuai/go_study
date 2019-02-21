package main

import "fmt"

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct{
	width,height,depth float64
	color Color
}

type BoxList []Box

func(b Box) Volume() float64{
	return b.width * b.height * b.depth
}

func(b *Box) SetColor(c Color){
	b.color = c
}

func(bl BoxList) BiggestColor() Color{
	big := 0.00
	c := Color(WHITE)
	for _,b := range bl{
		if bv := b.Volume(); bv>big{
			big = bv
			c = b.color
		}
	}
	return c
}

func (bl BoxList) PaintItBlack(){
	for i,_ := range bl{
		bl[i].SetColor(BLACK)
	}
}

func(c Color) String() string{
	strings := []string{"WHITH","BLACK","BLUE","RED","YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4,4,4,RED},
		Box{10,10,1,YELLOW},
		Box{1,1,20,BLACK},
		Box{10,20,2,BLUE},
		Box{10,30,1,WHITE},
		Box{20,20,20,YELLOW},
	}

	fmt.Printf("we have %d boxes in our set \n",len(boxes))
	fmt.Println("the volume of the first one is ",boxes[0].Volume())
	fmt.Println("the color of the last one is ",boxes[len(boxes)-1].color.String())
	fmt.Println("the biggest one is ",boxes.BiggestColor().String())

	boxes.PaintItBlack()
	fmt.Println("the color of the second on is ",boxes[1].color.String())

}