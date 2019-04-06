package main

import "fmt"

func main() {

	//bikin variable
	handshome := "saya ganteng"
	fmt.Println("berisikan variable string dengan isinya : ", handshome)

	//dapatkan pointernya
	pointerHandsome := &handshome
	fmt.Println("dapatin pointer dari variable handsome yg sudah di referencing sebelumnnya ialah", pointerHandsome)

	//baca nilai

	ganteng := *pointerHandsome
	fmt.Println("membaca nilai pointer yang sebelumnnya di referencing dan skrng di dereferencing", ganteng)
}
