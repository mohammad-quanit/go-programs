package main

import (
	"encoding/json"
	"fmt"
)

// *********************** Example 1 - Result **********************

type Result struct {
	Code    int                    `json:"code"`    // field tag `code`
	Message string                 `json:"msg"`     // field tag `code`
	Data    map[string]interface{} `json:"my Data"` // field tag `code`
}

func example1() {
	// By convention, Go uses the same title cased attribute names as are present in the case insensitive
	// JSON properties. So the Code attribute in our Result struct will map to the code,
	// or Code or cODe JSON property.
	jsonStr := `{"cODe":200,"msg":"success","my Data":{"url":"https:\/\/mp.weixin.qq.com\/cgi-bin\/showqrcode?ticket=gQHQ7jwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAyX3pqS0pMZlA4a1AxbEJkemhvMVoAAgQ5TGNYAwQsAQAA"}}`
	var res Result
	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Code, res.Data["url"])
}

// *********************** Example 2 - Bird **********************

type Dimensions struct {
	Height int
	Width  int
}
type Bird struct {
	Species     string
	Description string
	Dimensions  Dimensions
	// Dimensions  struct {
	// 	Height int
	// 	Width  int
	// }
}

// JSON Arrays
func example2() {
	birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"}, {"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	err := json.Unmarshal([]byte(birdJson), &birds)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Birds: %+v\n\n", birds[0])
}

// Nested Objects
func example3() {
	birdJson := `[{"sPECies":"Duck","description":"likes to hang on water", "dimensions": {"height":24, "width": 10}}, {"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	err := json.Unmarshal([]byte(birdJson), &birds)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Birds: %+v\n\n", birds)
}

// Primitive Types
func example4() {
	numberJson := "3"
	floatJson := "3.1412"
	stringJson := `"bird"`

	var n int
	var pi float64
	var str string

	json.Unmarshal([]byte(numberJson), &n)
	fmt.Println(n)

	json.Unmarshal([]byte(floatJson), &pi)
	fmt.Println(pi)

	json.Unmarshal([]byte(stringJson), &str)
	fmt.Println(str)
}

// JSON decodes to map - unstructured data

func example5() {
	birdsJson := `{"birds":{"pigeon": "likes to perch on rocks", "eagle": "bird of prey"}, "animals":"none"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdsJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := result["birds"].(map[string]interface{})
	fmt.Println(birds)
	for key, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}
}

func main() {
	// example1()
	// example2()
	// example3()
	// example4()
	example5()
}
