/*Розрахувати час, що знадобиться на подорож до Марсу. Швидкість світла була б ідеальною
для польоту. Відстань від Землі до Марса може відрізнятися залежно від того, в даний момент
часу планети знаходяться на сонячній орбіті.*/

package main

import "fmt"

func main() {

	const speedLight = 299792 // км/сек

	var distanceMin = 56000000 // мінімальна відстань до Марсу

	var distanceMax = 401000000 // максимальна відстань до Марсу

	fmt.Println(distanceMin/speedLight, "секунд") // мінімальний час 186 секунд

	fmt.Println(distanceMax/speedLight, "секунд") // максимальний час 1337 секунд
}
