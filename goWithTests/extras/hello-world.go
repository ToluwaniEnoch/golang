package main

import (
	"fmt"
	"sync"
	// "time"
	// "net/http"
	// "strconv"
	// "reflect"
)

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type ConsoleWriter struct {}

// func (cw ConsoleWriter) Write(data []byte) (int, error){
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

func sayHello(){
	fmt.Println("Hello")
}

var wg = sync.WaitGroup{}

// var i int = 42
func main() {
	var msg = "Hello"

	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	} (msg)

	msg = "Goodbye"
	wg.Wait()


	// go sayHello()
	// time.Sleep(100 * time.Millisecond)


	// fmt.Println(sayMessage("Hello world", 3))

	// var a int = 42
	// var b *int = &a // * is a pointer to an integer
	// fmt.Println(a, *b) // is a deference to find the memory location being
	// //pointed to and retrieve thr value

	// a = 27 //changes a and b to 27
	// fmt.Println(a, *b)

	// *b = 14 // changes a and b to 14
	// fmt.Println(a, *b)


	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }



	// fmt.Println("start")
	// fmt.Println("middle")
	// defer fmt.Println("end")

	// s := []int{1,2,3}
	// for _,v := range s {
	// 	fmt.Println( v)
	// }

	// for i :=0; i<5; i++ {
	// 	fmt.Println(i)
	// }


	// switch 20 {
	// case 1, 5, 10 :
	// 	fmt.Println("one, five, ten")
	// case 2, 4, 6:
	// 	fmt.Println("two, four, six")
	// default: 
	// fmt.Println("another number")
	// }


	// if false {
	// 	fmt.Println("The test is true")
	// }

	// t := reflect.TypeOf(Animal{})
	// field, _ := t.FieldByName("Name")
	// fmt.Println(field.Tag)
	

	// b := Bird{
	// 	Animal: Animal{Name: "Emu", Origin: "Austrialia"},
	// 	SpeedKPH: 48,
	// 	CanFly: false,
	// }

	// fmt.Println(b)

	//anonymous struct
	// aDoctor := struct{name string} {name: "Tim Lahm"}
	// anotherDoctor := aDoctor
	// anotherDoctor.name = "Tom Baker"
	// fmt.Println(aDoctor)
	// fmt.Println(anotherDoctor)


	//structs
	// aDoctor := Doctor{
	// 	number: 3,
	// 	actorName: "Tim Ford",
	// 	companions: []string {
	// 		"ABC", "JFK", "SJM",
	// 	},
	// }


	// fmt.Println(aDoctor)
	// fmt.Println(aDoctor.companions)



	
	//maps
	// statePopulations := make(map[string]int)
	// statePopulations = map[string]int{
	// 	"California": 33,
	// 	"Texas": 323,
	// 	"Florida": 4353,
	// }

	// for k,v := range statePopulations{
	// 	fmt.Println(k, v)
	// }
	// statePopulations["Georgia"] = 3442

	// fmt.Println(statePopulations)
	// delete(statePopulations, "Georgia")

	// fmt.Println(statePopulations)
	
	// if _, isPresent := statePopulations["Texas"]; isPresent {
	// 	fmt.Println("is present")
	// }
	// fmt.Println(statePopulations)
	// fmt.Println(len(statePopulations))
	// fmt.Println(isPresent)
	// fmt.Println(statePopulations["Texas"])

	// sp := statePopulations
	// fmt.Println(sp)


	//slices
	// a := []int{1, 2, 3}
	// fmt.Printf("Length: %v \n", len(a))
	// fmt.Printf("Capacity: %v \n", cap(a)) //len of array backing the slice

	// f := make([]int, 3)

	// f := []int{}
	// f = append(f, 1)
	// f = append(f, 2,3,4,5)
	// f = append(f, []int{6,7,8,9}...)
	// fmt.Println(f)

	// b := append(f[:2], f[4:]...)
	// fmt.Println(b)

	// fmt.Printf("Length: %v \n", len(f))
	// fmt.Printf("Capacity: %v \n", cap(f)) //len of array backing the slice

	//arrays
	// grades := [3]int {73, 34, 90}
	// gradesB := [...]int {73, 34, 90}
	// var students [3]string
	// fmt.Printf("Students:%v \n", students)

	// students[0] = "Lisa"
	// students[1] = "Rebecca"
	// students[2] = "Jacob"
	// fmt.Printf("Students:%v \n", students)
	// fmt.Printf("Students length:%v \n", len(students))

	// fmt.Printf("Grades:%v \n", grades)
	// fmt.Printf("Grades:%v \n", gradesB)



	// const myConst int = 14
	// fmt.Printf("%v, %T\n", myConst, myConst)
	// var n bool = true
	// fmt.Printf("%v, %T\n",n ,n)

	// n := 1 == 1
	// m := 1 == 2

	// fmt.Printf("%v, %T\n",n ,n)
	// fmt.Printf("%v, %T\n",m ,m)

	// fmt.Println("hello world")

	// var i int = 42
	// i := 42
	// fmt.Println(i)
	// fmt.Printf("%v, %T", i, i)

	// var i int = 42
	// fmt.Printf("%v, %T\n", i, i)

	// var j string
	// //convert int to string
	// j = strconv.Itoa(i)
	// fmt.Printf("%v, %T\n", j, j)

}

// type Doctor struct {
// 	number int
// 	actorName string
// 	companions []string
// }

// type Animal struct {
// 	Name string `required max:"100"`
// 	Origin string
// }

// type Bird struct {
// 	Animal
// 	SpeedKPH float32
// 	CanFly bool
// }

// func sayMessage(msg string, index int) string {
// 	fmt.Println(msg)
// 	fmt.Println("index of :: ", index)

// 	return msg
// }
