package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {

	defer fileClose()
	log.Println("가변 인자 사용하기")
	log.Println("중간값 찾기")

	result, err := inRange(-10, 10, 3, 5, 100, 2, 6, 7, 3, 4)
	if err != nil {
		log.Panic(err)
	}

	log.Println(result)
}

func inRange(min float64, max float64, nums ...float64) ([]float64, error) {
	if min > max {
		return nil, fmt.Errorf("최소값이 더 큰데?")
	}
	var result []float64
	for _, num := range nums {
		if min <= num && num < max {
			result = append(result, num)
		}
	}
	return result, nil
}
