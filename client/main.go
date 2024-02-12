
package main

import (
   "io/ioutil"
   "net/http"
   "fmt"
   "encoding/json"
)

import "bytes"

type responseBody struct {
    Grid [3][3]int `json:"grid"`
	Result string `json:"result"`
	Computer int `json:"computer"`
}

type requestBody struct {
    Grid [3][3]int `json:"grid"`
	Move int `json:"move"`
}

func getMarker( i int) string{
	if i == 1 {
		return "X"
	} else if i == 0{
		return "O"
	}

	return "*"
}

func PrintGrid(grid [3][3]int){
	for i := 0; i<3 ;i++{
		for j:=0; j<3 ; j++{
			fmt.Print(getMarker(grid[i][j]))
			fmt.Print(" | ")
		}
		fmt.Println()
	}

}

func main() {
   resp, err1 := http.Get("http://localhost:8080/start")
   if err1 != nil {

   }
   defer resp.Body.Close()
   var response responseBody
   body, err2 := ioutil.ReadAll(resp.Body)
   if err2 != nil{
	   
   }
   if err := json.Unmarshal(body,&response); err != nil {
    panic(err)
  }
  var data requestBody
  data.Grid = response.Grid


  PrintGrid(response.Grid)
  for true {
	  fmt.Print("Enter your choice from 1 to 9: ")
	  var input int
	  fmt.Scanln(&input)
	  if input < 1 || input > 9 {
		  fmt.Println("Enter valid number")
		  continue
	  }

	  fmt.Println("Client Move: ",input)
	  data.Move = input
	  requestByte, _ := json.Marshal(data)


	  Moveresp,error := http.Post("http://localhost:8080/move","application/json",bytes.NewReader(requestByte))

	  if error != nil{

	  }
	  defer Moveresp.Body.Close()
	  body, err2 = ioutil.ReadAll(Moveresp.Body)
	  var resp responseBody
	  if err := json.Unmarshal(body,&resp); err != nil {
		panic(err)
	  }
	  
	  fmt.Println("Computer Move: ",resp.Computer)
	  PrintGrid(resp.Grid)
	  if(resp.Result != ""){
		  fmt.Println(resp.Result)
		  break
	  }
	  data.Grid = resp.Grid

  }


}