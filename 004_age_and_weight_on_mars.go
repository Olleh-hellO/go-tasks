/*Рік на Марсі триває 687 земних днів. Його гравітація на 38% менше, ніж на Землі.
Визначити ваш вік та вагу на Марсі.*/

package main

import "fmt"

func main() {

	fmt.Println("Моя вага на Марсі становить", 85*0.3783, "кілограм") // результат 32,1555
	fmt.Println("Мій вік на Марсі становить", 39*365/687, "років")    // результат 20.7205
}
