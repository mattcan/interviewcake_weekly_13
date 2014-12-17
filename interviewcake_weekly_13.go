package main

import "fmt"

func main() {
	my_orders := []int{3,4,6,10,11,15}
	alices_orders := []int{1,5,8,12,14,19}

	merge_orders(my_orders, alices_orders)
}

func merge_orders(ordersa []int, ordersb []int) {
	all_orders := make([]int, 1)

	orderCount := (len(ordersa) + len(ordersb)) - 1 // cheap hack to not double print last number

	oAcounter, oBcounter := 0, 0

	for i := 0; i < orderCount; i++ {

		if ordersa[oAcounter] < ordersb[oBcounter] {
			all_orders = append(all_orders, ordersa[oAcounter])

			oAcounter++
			if oAcounter >= len(ordersa) {
				oAcounter = len(ordersa) - 1
			}

		} else {
			all_orders = append(all_orders, ordersb[oBcounter])

			oBcounter++
			if oBcounter >= len(ordersb) {
				oBcounter = len(ordersb) - 1
			}

		}
	}

	// tack on the last order by reversing the check
	if ordersa[oAcounter] > ordersb[oBcounter] {
		all_orders = append(all_orders, ordersa[oAcounter])
	} else {
		all_orders = append(all_orders, ordersb[oBcounter])
	}

	// remove first element, silly go
	all_orders = append(all_orders[:0],all_orders[0+1:]...)

	fmt.Println(all_orders)
}
