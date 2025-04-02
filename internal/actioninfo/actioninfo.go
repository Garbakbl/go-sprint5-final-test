package actioninfo

import "fmt"

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Println(err)
		}
		result, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}
