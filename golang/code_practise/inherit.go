package code_practise

type Car interface {
	IsOil()
	PrintBrandName()
}

type BYD struct {
}

func (byd *BYD) IsOil() {
	println("i am byd  do not need oil")
}
func (byd *BYD) PrintBrandName() {
	println("i am byd")
}
func (byd *BYD) Myself() {
	println("i am byd, Myself func")
}

type Audi struct {
}

func (ad *Audi) IsOil() {
	println("i am audi need oil")
}
func (ad *Audi) PrintBrandName() {
	println("i am audi")
}

func testPrint(car Car) {
	car.IsOil()
	car.PrintBrandName()
}
func Inherit() {
	byd := new(BYD)
	testPrint(byd)

	ad := new(Audi)
	testPrint(ad)

}
