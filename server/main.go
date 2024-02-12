package main

import "github.com/gin-gonic/gin"
import (b "board")
import "fmt"

type requestBody struct {
    Grid [3][3]int `json:"grid"`
	Move int `json:"move"`
}

func main() {
	
	r := gin.Default()
	r.GET("/start", func(c *gin.Context) {
		var bb b.Board
		c.JSON(200, gin.H{
			"grid": bb.Init(),
		})
	})

	r.POST("/move", func(c *gin.Context) {
		req := new(requestBody)
		err := c.Bind(&req)
		if err != nil{

		}
		var bb b.Board
		bb.Grid = req.Grid
		req.Move--;
		bb.MakeMove(req.Move / 3,req.Move % 3,1)
		isGameOver,_ := bb.IsGameOver()
		if(isGameOver){
			c.JSON(200,gin.H{
				"grid": bb.Grid,
				"result": "Player Wins"})
		}

		depth := bb.CalculateDepth()
		if depth == 0{
			c.JSON(200,gin.H{
				"grid": bb.Grid,
				"result": "Draw",})
		} 
		res := bb.Minimax(depth,false)
		bb.MakeMove(res[0],res[1],0)
		isGameOver,_  = bb.IsGameOver()
		if(isGameOver){
			fmt.Println(isGameOver)

			c.JSON(200,gin.H{
				"grid": bb.Grid,
				"result": "Computer wins",
				"computer": (3 * res[0]) + res[1] + 1})
		} else{
			c.JSON(200, gin.H{
				"grid" : bb.Grid,
				"result": "",
				"computer": (3 * res[0]) + res[1] + 1})
		}


		
	})

	r.Run() 
}