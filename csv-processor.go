package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

func csvReaders(inputColumns []string, fileLocation string) {
	for _, column := range inputColumns {
		fmt.Printf("Finding values on column = %v\n", column)
		csvReader(column, fileLocation)
	}
}

func csvReader(inputColumn, fileLocation string) {
	startTime := time.Now()

	// get all the files from the specific folder
	files, err := ioutil.ReadDir(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	noColumn := make(chan struct{}, len(files))
	results := make(chan map[string]bool, len(files))

	// reading all the files from dir concurrently
	for _, file := range files {
		wg.Add(1)
		// and read it concurrently to get the data from specific column
		go func(file os.FileInfo) {
			defer wg.Done()
			f, err := os.Open(path.Join(dir, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			datas, ok := readFile(inputColumn, f)

			if !ok {
				noColumn <- struct{}{}
				return
			}

			results <- datas

		}(file)

	}

	wg.Wait()

	// check if we got the column or not
	select {
	case <-noColumn:
		fmt.Printf("The column name = %v doesnt exist\n", inputColumn)
		return

	default: // do nothing and continue
	}

	close(results)
	close(noColumn)

	//receive results and determine which size datas is the smallest
	theSameValue := getSameValues(results)

	fmt.Printf("final result = %+v\n", theSameValue)
	fmt.Printf("final result size = %+v\n", len(theSameValue))
	fmt.Printf("time consume = %+v\n", time.Since(startTime).Seconds())

}

func readFile(inputColumn string, f *os.File) (map[string]bool, bool) {
	r := csv.NewReader(bufio.NewReader(f))

	isColumnExist := false
	columnIndex := -1
	datas := make(map[string]bool)
	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		readRecords(records, datas, inputColumn, &columnIndex, &isColumnExist)
	}

	return datas, isColumnExist
}

func readRecords(records []string, datas map[string]bool, inputColumn string, columnIndex *int, isColumnExist *bool) {

	if len(records) > 0 {
		//find index column
		if *columnIndex == -1 {
			for index, record := range records {
				if record == inputColumn {
					*isColumnExist = true
					*columnIndex = index
				}
			}

		} else if *columnIndex >= 0 {
			datas[records[*columnIndex]] = true
		}
	}
}

// getting the same value from all the datas
func getSameValues(results chan map[string]bool) []string {
	var datas = make([]map[string]bool, len(results))
	minIndex := -1
	minSize := int(MaxUint >> 1)
	i := 0
	for values := range results {
		sizeValues := len(values)
		if sizeValues < minSize && sizeValues > 0 {
			minSize = sizeValues
			minIndex = i
		}
		datas[i] = values
		i++
	}

	// getting the same value from all the datas
	var theSameValue []string
	for value, _ := range datas[minIndex] {
		// check if all value is exist
		isExistAll := true
		for _, data := range datas {
			isExistAll = isExistAll && data[value]
		}

		if isExistAll {
			theSameValue = append(theSameValue, value)
		}
	}

	return theSameValue
}
