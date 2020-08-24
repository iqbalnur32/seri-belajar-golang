package main
import "fmt"

/*
* Reusable Function In Golang
*/
type typeDataFunction = func(a, b string) string

func main(){

	var explicit typeDataFunction = func(a, b string) string {

		return fmt.Sprintf("%s %s", a, b)
	}

	implict := func(a, b string) string {

		return fmt.Sprintf("%s %s", b, a)
	}

	functionTypeCustomer(explicit)
	functionTypeCustomer(implict)
	functionTypeCustomer(func(a, b string) string {
		return fmt.Sprintf("%s %s!", a, b)
	})

	// Rause Multiple struct
	fmt.Println("++++++++ Multiple Rause struct in interface +++++")
	var a Pedagang
	a = Pempek{}
	fmt.Println(a.Speak())
}

func functionTypeCustomer(fn typeDataFunction) {
	s := fn("hallo", "Iqbal")
	fmt.Println(s)
}

/*
	Reuse Multiple struct in interface
*/

type Pedagang interface {
	Speak() string
}

type TukangTahu struct {}

func (d TukangTahu) Speak() string {
	return "Suara abang Tukang Tahu Ialah : Tahuu !!!!, Tahu !!!!"
}

type Pempek struct {}

func (d Pempek) Speak() string {
	return "Pempek Palembang Kapal Selam 200an Mntap lo"
}


