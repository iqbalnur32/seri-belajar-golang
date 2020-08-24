package main

import "fmt"

/*
	Pass Value Refrence By Interface with pointer
*/
func main(){
	var result Pedagang
	ParamsPedagangInterface(&result)
	fmt.Println(result)
}

type Pedagang struct {
	ID string
	Nama string
	Jualan string
}

func ParamsPedagangInterface(result interface{}) {
	// switch a := result.(type) {
	// case *Pedagang:

	// 	*a = Pedagang{
	// 		ID: "1",
	// 		Nama: "Sutanso",
	// 		Jualan : "Pengusaha Mineral",
	// 	}
	// }

	data := result.(*Pedagang)
	*data = Pedagang{
		ID: "1",
		Nama: "Sutanso",
		Jualan : "Pengusaha Mineral",
	}
}