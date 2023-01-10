package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func updateBalance(updateAddr string, balance string) {
	var balances []string

	file, err := os.OpenFile("blocks/genesisblock.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	scanFile, err := os.Open("blocks/genesisblock.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(scanFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		balances = append(balances, fileScanner.Text())
	}

	if err := os.Truncate("blocks/genesisblock.txt", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

	defer file.Close()
	defer scanFile.Close()

	for _, line := range balances {
		fmt.Println(line)
		addr, _, _ := strings.Cut(line, ",")
		if addr == updateAddr {
			fmt.Println(line)
			if _, err = file.WriteString(updateAddr + "," + balance + "\n"); err != nil {
				panic(err)
			}
		} else {
			if _, err = file.WriteString(line + "\n"); err != nil {
				panic(err)
			}
		}
	}
}

func appendNewAddress(updateAddr string) {
	file, err := os.OpenFile("blocks/genesisblock.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	if _, err = file.WriteString(updateAddr + ",0\n"); err != nil {
		panic(err)
	}
	defer file.Close()
}

func getBalance(fromAddr string) int {
	readFile, err := os.Open("blocks/genesisblock.txt")
	if err != nil {
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	strBalance := "-1"
	for fileScanner.Scan() {
		line := fileScanner.Text()
		addr, amount, _ := strings.Cut(line, ",")
		if addr == fromAddr {
			strBalance = amount
		}
	}
	readFile.Close()
	balance, _ := strconv.Atoi(strBalance)
	return balance
}

func transactionValidation(received string) (string, string, string, int, int) {
	fmt.Println("ADDR: ", received)
	fromAddr, rest, _ := strings.Cut(received, ",")
	toAddr, strAmount, _ := strings.Cut(rest, ",")
	amount, _ := strconv.Atoi(strAmount)
	balance := getBalance(fromAddr)
	toBalance := getBalance(toAddr)
	if balance < 0 {
		return "NOT FOUND", toAddr, fromAddr, 0, 0
	} else if toBalance < 0 {
		appendNewAddress(toAddr)
	}
	if balance >= amount {
		return "OK", toAddr, fromAddr, balance - amount, toBalance + amount
	}
	return "ERROR", toAddr, fromAddr, 0, 0
}
