package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type board struct {
	values [][]int
}

func collapseNums(nums []int) []int {
	// base case, only one number left
	if len(nums) == 1 {
		return nums
	} else if nums[0] == nums[1] {
		nums[0] = 0
		a := append(nums[1:2], nums[2:]...)
		return collapseNums(a)
	} else {
		return append(collapseNums(nums[0:1]), collapseNums(nums[1:])...)
	}
}

func collapseNums2(nums [][]int) [][]int {
	// base case, only one number left
	if len(nums) == 1 {
		return nums
	} else if nums[0][0] == nums[1][0] {
		//copy to new var
		a := nums
		// add value to inner slice

		a[0] = append(a[0], a[0][0])
		// cocatonate into new slice, removing the second entry
		//var b [][]int
		b := append(a[0:1], a[2:]...)

		return collapseNums2(b)
	} else {
		return append(collapseNums2(nums[0:1]), collapseNums2(nums[1:])...)
	}
}

func collapseNums3(nums []int, final *[]int) []int {
	// base case, only one number left
	if len(nums) == 1 {
		return nums
	} else if nums[0] == nums[1] {
		*final = append((*final), 0)
		a := append(nums[1:2], nums[2:]...)
		return collapseNums3(a, final)
	} else {
		return append(collapseNums3(nums[0:1], final), collapseNums3(nums[1:], final)...)
	}
}

func collapseNums4(nums [][]int, output *[]int) [][]int {
	// base case, only one number left
	if len(nums) == 1 {
		return nums
	} else if nums[0][0] == nums[1][0] {
		//add a zero to final
		*output = append((*output), 0)

		//copy to new var
		a := nums
		//fmt.Println("a: ", a)
		// add value to inner slice
		//fmt.Println("a1 before: ", a[1])
		a[1] = append(a[1], a[0][0:]...)
		//fmt.Println("a1 after: ", a[1])
		// cocatonate into new slice, removing the second entry
		//var b [][]int

		b := append(a[1:2], a[2:]...)
		//fmt.Println("b : ", b)

		return collapseNums4(b, output)
	} else if nums[0][0] != 0 && nums[1][0] == 0 { // need to check if next val is a zero, and swap
		// swap them
		tmp := nums[0]
		nums[0] = nums[1]
		nums[1] = tmp
		// call func
		return append(collapseNums4(nums[0:1], output), collapseNums4(nums[1:], output)...)
	} else {
		return append(collapseNums4(nums[0:1], output), collapseNums4(nums[1:], output)...)
	}
}

func processRow(rowVals []int) []string {
	// make the 2D holder
	toCollapse := make([][]int, 4)
	for i, v := range rowVals {
		toCollapse[i] = []int{v}
	}

	final := make([]int, 0)
	a := make([]int, 0)
	toReturn := make([]string, 0)

	//fmt.Println("tocolaps: ", toCollapse)
	j := collapseNums4(toCollapse, &a)
	//fmt.Println("j: ", j)

	for _, v := range j {
		//fmt.Println(v)
		total := 0
		for _, v1 := range v {
			total += v1
		}
		final = append(final, total)
	}

	fmt.Println("final: ", final)
	a = append(a, final...)

	for i := 0; i < len(a); i++ {
		toReturn = append(toReturn, strconv.Itoa(a[i]))
	}
	//toReturn = append(toReturn, strconv.Itoa(a))

	fmt.Println("toReturn: ", toReturn)

	//matrix := [4][4]int{}
	//fmt.Println(board)
	return toReturn

}

// SERVER HANDLER STUFF

type dataToServe struct {
	data [][]int
}

// Make handler by adding ServeHttp method
func (m dataToServe) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// first thing we do is parse the form
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// data := struct {
	// 	Method        string
	// 	URL           *url.URL
	// 	Submission    url.Values
	// 	Header        http.Header
	// 	ContentLength int64
	// }{
	// 	req.Method,
	// 	req.URL,
	// 	req.Form,
	// 	req.Header,
	// 	req.ContentLength,
	// }

	// Write response
	w.Header().Set("Joe-Key", "this is from joe")              // you can make what ever headers you want
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // text/plain tells the browser to not interpret html
	w.WriteHeader(200)
	// grab the Form field from request
	// this will give you url values - map with key= string and value = []string
	// r.Form
	// tpl.ExecuteTemplate(w, "form.html", data)
	tpl.ExecuteTemplate(w, "index.html", m.data)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	//fmt.Println("in index handler")
	err := req.ParseForm()
	if err != nil {
		log.Fatalln("error with get", err)
	}

	//Check if get method has newgame parameter
	if req.Method == "GET" {

		//check for new game request param
		newGame := req.URL.Query()
		//fmt.Println(newGame)

		//fmt.Println("get: ", newGame.Get("newgame"))
		if newGame.Get("newgame") == "New Game" {
			//fmt.Println("inNew game")
			//make new board and execute template
			newBoard := makeNewBoard()

			// write response
			w.Header().Set("New Game", "True")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(200)

			// execute template with new board setup
			tpl.ExecuteTemplate(w, "index.html", newBoard)
			return
		} else {
			// do nothing
			w.Header().Set("New Game", "False")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(200)
			return
		}
	}

	// POST will be used when the client slides the board pieces
	// if the method is post, pull out body and check which direction to slide the tiles
	if req.Method == "POST" {
		//fmt.Println("body: ", req.Body)
		decoder := json.NewDecoder(req.Body)
		var reqBod requestBody
		// fmt.Println()
		err := decoder.Decode(&reqBod)
		if err != nil {
			panic(err)
		}
		//log.Println(reqBod.Direction)

		// pull out the values of the POST body
		// reqBody := req.PostForm
		// fmt.Println("body: ", reqBody)
		// fmt.Printf("reqbody type: %T", reqBody)

		// if reqBody.Get("direction") == "right" {
		// 	fmt.Println("in GET: ")
		// 	fmt.Println("values", reqBody.Get("values"))

		// 	fmt.Printf("type: %T", reqBody.Get("values"))

		// }
		if reqBod.Direction == "right" {

			// fmt.Println("in the post, driection: ", reqBody.Get("direction"))
			fmt.Println("direction: ", reqBod.Direction)
			fmt.Println("values: ", reqBod.Values)

			// build back up the board and thenn call the slide tiles function
			newBoard := reconstructBoard(reqBod.Values)

			fmt.Println("newboard: ", newBoard)

			responseArray := make([]string, 0)
			// call slide right function
			fmt.Println("newboard before: ", newBoard)
			for _, v := range newBoard {
				responseArray = append(responseArray, processRow(v)...)
			}

			// convert back to string

			fmt.Printf("%T", responseArray)

			fmt.Println("newboard after: ", responseArray)

			// respBod := &responseBody{
			// 	NewBoard: newBoard,
			// }

			//marshal into JSON
			respBod, err := json.Marshal(responseArray)

			fmt.Println("respBod: ", string(respBod))

			if err != nil {
				fmt.Println("error in json marshall")
			}
			//fmt.Println("response: ", respBod2)
			//add to response body
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(respBod))

			// send values back to ajax

			// Write response

		}
		//fmt.Println(reqBody)
		// do something
	}

}

func reconstructBoard(values []string) [][]int {
	newBoard := make([][]int, 4)
	for i := 0; i < 4; i++ {
		newBoard[i] = make([]int, 0)
		//fmt.Println(i)
		for j := (4 * i); j < (i*4 + 4); j++ {
			//fmt.Println("j: ", j)
			currentVal, _ := strconv.Atoi(values[j])
			newBoard[i] = append(newBoard[i], currentVal)
		}
		//fmt.Println("new j")
	}
	return newBoard
}

// request data type
type requestBody struct {
	Direction string
	Values    []string
}

type responseBody struct {
	NewBoard [][]int
}

// make global var for out pointer to template
var tpl *template.Template

// call init func to parse all templates
func init() {
	tpl = template.Must(template.ParseFiles(("static/index.html")))

}

func makeNewBoard() [][]int {
	board := make([][]int, 4)
	for i := 0; i < 4; i++ {
		board[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			board[i][j] = rand.Intn(4)
		}
	}
	fmt.Println("borad", board)
	return board
}

func main() {

	// set seed for random num generator
	rand.Seed(time.Now().UTC().UnixNano())

	// make an empty matrix
	board := makeNewBoard()

	// b := []int{2, 3, 2, 3}
	// a = append(a, collapseNums3(b, &a)...)
	// //fmt.Println(collapseNums3(b, &a))
	// fmt.Println(a)
	// //fmt.Println(collapseNums(b))

	// m := make([][]int, 4)
	// for i := 0; i < len(m); i++ {
	// 	//m[i][0] = 0
	// 	m[i] = make([]int, 1)
	// 	//m[i][0] = 2
	// }

	// m[0][0] = 2
	// m[1][0] = 2
	// m[2][0] = 3
	// m[3][0] = 3

	//var m = []int{2, 2, 3, 3}

	for _, v := range board {
		processRow(v)
	}

	// server stuff
	// d := dataToServe{
	// 	data: board,
	// }

	//fmt.Println(d)
	http.HandleFunc("/", indexHandler)

	//load static files
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("static/scripts"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
	//<link rel="stylesheet" href="/css/styles.css">
}
