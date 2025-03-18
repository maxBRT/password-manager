package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func addPassword(filePath string, s []string) error {
	fileExits := false

	if _, err := os.Stat(filePath); err == nil {
		fileExits = true
	}

	if fileExits {
		f, err := os.Open(filePath)
		check(err)

		reader := csv.NewReader(f)

		records, err := reader.ReadAll()
		check(err)

		for _, row := range records {
			if slices.Contains(row, s[0]) {
				f.Close()
				return errors.New("Password already exist for app")
			}
		}
		f.Close()
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	if !fileExits {
		header := []string{"App", "Password"}
		err = writer.Write(header)
		check(err)
	}

	row := s
	err = writer.Write(row)
	check(err)

	return nil
}

func getPassword(filePath string, app string) ([]string, error) {
	var s []string
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	check(err)

	for _, row := range records {
		if slices.Contains(row, strings.ToLower(app)) {
			return row, nil
		}
	}
	return s, errors.New("Password not found.")
}

func deletePassword(filePath string, app string) error {
	pswFound := false
	f, err := os.Open(filePath)
	defer f.Close()
	check(err)

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	check(err)

	var filteredRecord [][]string
	filteredRecord = append(filteredRecord, records[0])

	for _, row := range records[1:] {
		if slices.Contains(row, strings.ToLower(app)) {
			pswFound = true
			continue
		} else {
			filteredRecord = append(filteredRecord, row)
		}
	}
	if pswFound {
		err := os.Remove(filePath)
		check(err)

		newF, err := os.Create(filePath)
		check(err)

		writer := csv.NewWriter(newF)
		defer writer.Flush()

		err = writer.WriteAll(filteredRecord)
		check(err)
		return nil
	} else {
		return errors.New("Password not found.")
	}
}

func updatePassword(filePath string, updatedRow []string) error {
	pswFound := false
	f, err := os.Open(filePath)
	defer f.Close()
	check(err)

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	check(err)

	var filteredRecord [][]string
	filteredRecord = append(filteredRecord, records[0])

	for _, row := range records[1:] {
		if slices.Contains(row, strings.ToLower(updatedRow[0])) {
			pswFound = true
			continue
		} else {
			filteredRecord = append(filteredRecord, row)
		}
	}
	if pswFound {
		filteredRecord = append(filteredRecord, updatedRow)

		err := os.Remove(filePath)
		check(err)

		newF, err := os.Create(filePath)
		check(err)

		writer := csv.NewWriter(newF)
		defer writer.Flush()

		err = writer.WriteAll(filteredRecord)
		check(err)

		return nil
	} else {
		return errors.New("Password not found.")
	}
}
