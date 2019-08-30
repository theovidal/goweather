package goweather

import (
	"fmt"
	"time"
)

func ExampleAPI_Current() {
	api, err := NewAPI("", "en", "metric")
	if err != nil {
		panic(err)
	}
	data, err := api.Current("Lyon,FR")
	if err != nil {
		panic(err)
	}

	fmt.Println(data.City.Name, ":", data.Conditions.Description)
	if data.Conditions.Clouds < 20 {
		println("What a beautiful moment ! Go outside !")
	}
}

func ExampleAPI_Forecast() {
	api, err := NewAPI("", "en", "metric")
	if err != nil {
		panic(err)
	}
	data, err := api.Forecast("Lyon,FR")
	if err != nil {
		panic(err)
	}

	fmt.Println(data.City.Name)
	println("---------------")

	for _, condition := range data.Conditions {
		conditionTime := time.Unix(int64(condition.Timestamp), 0)
		fmt.Printf("At %d:%d : %s", conditionTime.Hour(), conditionTime.Minute(), condition.Description)
	}
}
