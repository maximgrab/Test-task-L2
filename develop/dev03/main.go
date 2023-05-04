package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func numericalSort(data []string) ([]string, error) {
	var (
		result  []string
		nums    []int
		tempStr string
	)
	for _, str := range data {
		tempStr = strings.TrimSuffix(str, "\n")
		tempStr = strings.Replace(tempStr, "\r", "", 1)
		n, err := strconv.Atoi(tempStr)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	sort.Ints(nums)
	for _, str := range nums {
		result = append(result, fmt.Sprintf("%s\n", strconv.Itoa(str)))
	}
	return result, nil
}
func sortByColomn(data []string, colomn int) ([]string, error) {
	var (
		splitedData [][]string
		result      []string
	)
	for _, str := range data {
		spl := strings.Split(str, " ")
		if len(spl) < colomn {
			return nil, errors.New("Not enough colomns")
		}
		splitedData = append(splitedData, spl)
	}
	sort.Slice(splitedData, func(i, j int) bool {
		return splitedData[i][colomn] < splitedData[j][colomn]
	})
	for _, str := range splitedData {
		result = append(result, strings.Join(str, " "))
	}
	return result, nil
}

var (
	kFlag = flag.Int("k", 0, "Sort by colomn")
	nFlag = flag.Bool("n", false, "compare according to string numerical value")
	rFlag = flag.Bool("r", false, "reverse the result of comparisons")
	uFlag = flag.Bool("u", false, "output only the first of an equal run")
)

func main() {
	fmt.Println(os.Args[len(os.Args)-1])
	flag.Parse()
	file, err := os.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	data := strings.Split(string(file), "\n")

	if *kFlag > 0 {
		res, err := sortByColomn(data, *kFlag)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		fmt.Print(res)
	} else if *nFlag {
		res, err := numericalSort(data)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		fmt.Print(res)
	} else if *rFlag {
		res, err := numericalSort(data)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		fmt.Print(res)
	}
}
