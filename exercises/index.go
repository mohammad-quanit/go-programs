package exercises

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToNumeric() {
	fmt.Println("----demo String To Numeric----")
	str_val1 := "5"
	str_val2 := "2.8769"

	var err error
	var int_val1 int64
	int_val1, err = strconv.ParseInt(str_val1, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("This is int value %d\n", int_val1)

	float_val2, err := strconv.ParseFloat(str_val2, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("This is float value %.2f\n", float_val2)
}

func NumericToString() {
	fmt.Println("----demo numeric to string----")
	num1 := 5.872

	str_num1 := fmt.Sprintf("this is stringified number '%f' \n", num1)
	fmt.Println(str_num1)
}

func StringParser() {
	fmt.Println("----demo String Parser----")
	data := "Berlin;Amsterdam;London;Tokyo"
	fmt.Println(data)

	cities := strings.Split(data, ";")
	fmt.Println(cities)

	for _, city := range cities {
		fmt.Println(city)
	}
}

func StringLength(str_data string) {
	fmt.Println("----demo String Length----")
	len_data := len(str_data)
	fmt.Printf("length %d\n", len_data)
}

func StringCopy() {
	// [start: end]
	fmt.Println("----demo copy data----")
	sample := "hello world, go!"
	fmt.Println(len(sample))
	fmt.Println(sample[0:])                          // copy all
	fmt.Println(sample[:5])                          // copy 5 characters
	fmt.Println(sample[3:8])                         // copy characters from index 3 until 8
	fmt.Println(sample[len(sample)-5 : len(sample)]) // 5 copy characters from
	fmt.Println(sample[6:11])                        // copy only world from string
}
