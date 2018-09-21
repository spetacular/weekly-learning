package main

import (
	"fmt"
)

//物品数量
const AMOUNT = 5

//空间
const ROOM = 10

type item struct {
	name   string
	value  int
	weight int
}

var items [AMOUNT]item

var c [AMOUNT + 1][ROOM + 1]int

func maxValue(items [AMOUNT]item) int {
	for i := 0; i <= ROOM; i++ {
		c[0][i] = 0
	}

	for i := 1; i <= AMOUNT; i++ {
		for j := 1; j <= ROOM; j++ {
			//选择物品i之后的受益
			//weightSeletI := items[i-1].value + c[i][j-items[i-1].weight]
			//不选择物品i之后的受益
			//weightNotSelectI := c[i-1][j]

			if items[i-1].weight <= j && items[i-1].value+c[i][j-items[i-1].weight] > c[i-1][j] {
				c[i][j] = items[i-1].value + c[i][j-items[i-1].weight]
			} else {
				c[i][j] = c[i-1][j]
			}
		}
	}

	return c[AMOUNT][ROOM]
}

func main() {
	items[0] = item{name: "a", value: 1, weight: 2}
	items[1] = item{name: "b", value: 3, weight: 3}
	items[2] = item{name: "c", value: 5, weight: 4}
	items[3] = item{name: "d", value: 9, weight: 7}
	items[4] = item{name: "e", value: 6, weight: 5}
	m := maxValue(items)
	fmt.Println("Max weight is : ", m)
	fmt.Println("The best choices as following : ")
	remainSapce := ROOM
	for i := AMOUNT; i >= 1; i-- {
		if remainSapce >= items[i-1].weight {
			if c[i][remainSapce]-c[i-1][remainSapce-items[i-1].weight] == items[i-1].value {
				fmt.Println("item " + items[i-1].name + " is selected")
				remainSapce = remainSapce - items[i-1].weight
			}
		}
	}

}
