package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var m=make([]([]uint8), dy)
    for y, _ :=range m {
        m[y]=make([]uint8, dx)
        for x, _:=range m[y] {
            m[y][x]=uint8(x*y)
        }
    }
    
  return m 
}

func main() {
	pic.Show(Pic)
}