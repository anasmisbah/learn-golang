package main

import (
	"fmt"
	"strings"
	"strconv"
)

// TUGAS 1
func ArrayMerge(arrayA, arrayB []string) []string  {
	merged := append(arrayA, arrayB...)
	uniqueArray := []string{}
	for i := 0; i < len(merged); i++ {
		isNotDuplicate := true;
		for j := 0; j < len(uniqueArray); j++ {
			
				if merged[i] == uniqueArray[j]  {
					isNotDuplicate = false
				}
			
		}
		if isNotDuplicate {
			uniqueArray =append(uniqueArray, merged[i])
		}
	}
	return uniqueArray
}

func tugas_1() {
	// Test case
	fmt.Println(ArrayMerge([]string{"king","devil jin","akuma"},[]string{"eddie","steve","geese"})) 

	fmt.Println(ArrayMerge([]string{"sergei","jin"},[]string{"jin","steve","bryan"})) 
	
	fmt.Println(ArrayMerge([]string{"alisa","yoshimitsu"},[]string{"devil jin","yoshimitsu","alisa","law"})) 

	fmt.Println(ArrayMerge([]string{},[]string{"devil jin","sergei"})) 

	fmt.Println(ArrayMerge([]string{"hwoarang"},[]string{}))
	
	fmt.Println(ArrayMerge([]string{},[]string{}))
}

// TUGAS 2
func Mapping(slice []string) map[string]int  {
	unique := getUnique(slice)
	countMap := make(map[string]int)
	for _, uniqueVal := range unique {
		countMap[uniqueVal] = 0;
		for _, sliceValue := range slice {
			if sliceValue == uniqueVal {
				countMap[uniqueVal]++;
			}
		}
	}
	return countMap
}

func getUnique(slice []string) []string {
	uniqueArray := []string{}
	for i := 0; i < len(slice); i++ {
		isNotDuplicate := true;
		for j := 0; j < len(uniqueArray); j++ {
			
				if slice[i] == uniqueArray[j]  {
					isNotDuplicate = false
				}
			
		}
		if isNotDuplicate {
			uniqueArray =append(uniqueArray, slice[i])
		}
	}
	return uniqueArray
}

func tugas_2() {
	fmt.Println(Mapping([]string{"asd","qwe","asd","adi","qwe","qwe"}))
	fmt.Println(Mapping([]string{"asd","qwe","asd"}))
	fmt.Println(Mapping([]string{}))
}

// TUGAS 3
func munculSekali(angka string) []int {
	numbers := strings.Split(angka,"")
	// fmt.Println(len(numbers))
	uniqueArray := []int{}
	for i := 0; i < len(numbers); i++ {
		isNotDuplicate := true;
		for j := 0; j < len(numbers); j++ {
			
				if numbers[i] == numbers[j] && i!=j {
					isNotDuplicate = false
				}
			
		}
		if isNotDuplicate {
			val,_:=strconv.Atoi(numbers[i])
			uniqueArray =append(uniqueArray, val)
		}
	}
	return uniqueArray
}

func tugas_3() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}

func main() {
	fmt.Println("=== TUGAS 1 ===")
	tugas_1()
	fmt.Println("===============")
	fmt.Println("=== TUGAS 2 ===")
	fmt.Println()
	tugas_2()
	fmt.Println("===============")
	fmt.Println("=== TUGAS 3 ===")
	tugas_3()	
	fmt.Println("===============")
}