package main

import (
	"fmt"
	"strconv"
)

func countDistinctValues(stringMap map[string]string) map[float64]float64 {
	distinctValues := make(map[float64]float64)
	for _, valueStr := range stringMap {
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			fmt.Printf("Error parsing float: %v\n", err)
			continue
		}
		distinctValues[value] = 0
	}
	return distinctValues
}

func main() {
	stringMap := map[string]string{
		"00": "0.301010",
		"01": "0.301010",
		"02": "0.301010",
		"03": "0.301010",
		"04": "0.301010",
		"05": "0.301010",
		"06": "0.301010",
		"07": "0.625970",
		"08": "0.625970",
		"09": "0.625970",
		"10": "0.950930",
		"11": "0.950930",
		"12": "0.625970",
		"13": "0.625970",
		"14": "0.625970",
		"15": "0.950930",
		"16": "0.950930",
		"17": "0.950930",
		"18": "0.950930",
		"19": "0.950930",
		"20": "0.950930",
		"21": "0.625970",
		"22": "0.625970",
		"23": "0.301010",
	}

	distinctValues := countDistinctValues(stringMap)
	fmt.Println("Distinct values:", distinctValues)
}
