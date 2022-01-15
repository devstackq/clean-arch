package main

import (
	"log"

	"github.com/devstackq/go-clean/config"
	"github.com/devstackq/go-clean/server"
	"github.com/spf13/viper"
)

// connect doker mongo

func Josephus(items []int, k int) []int {
	// Your code: [1,2,3,4,5,6,7]w, k =?
	// if i > currLenItem;  ->
	//send CurrItem -> set LastIdx -> next Elem in circle
	//
	// k = k - 1
	last := 0
	// currentIdx := 0
	currentItems := items
	var res []int
	for i := 1; i < len(items)+1; i++ {
		// last = len(currentItems) + k - i //
		last = (k * i) - i
		//last +=
		log.Print(last)
		//5,10
		if last == len(currentItems) { //9-3-l
			// last = last - i - k
			// last = last - len(currentItems)
			//set index
			last = len(currentItems) // 3
		} else if last > len(currentItems) {
			// last=
		}

		log.Print(last, i, len(currentItems), currentItems, res)

		res = append(res, currentItems[last])
		//update items
		// separate; join
		//1234567; 123

		copy(currentItems[last:], currentItems[last+1:]) //2 : copy; from , where  data
		currentItems[len(currentItems)-1] = 0            //remove last
		currentItems = currentItems[:len(currentItems)-1]

		// log.Print(currentItems, len(currentItems), res, last, i, k)

	}
	return res
}
func main() {
	// log.Print(Josephus([]int{1, 2, 3, 4, 5, 6, 7}, 3), "res")

	if err := config.Init(); err != nil {
		log.Println(err, "viper")
		return
	}
	app := server.NewApp()
	if err := app.Run(viper.GetString("port")); err != nil {
		log.Println(err)
		return
	}

}
