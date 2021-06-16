package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

type HouseListings struct {
	Houses []House
}

type House struct {
	NumOfRooms int
	City       string
	Address    string
	Price      int
}

var (
	reader          = bufio.NewReader(os.Stdin)
	errInvalidInput = errors.New("please enter an valid input")
	listings        HouseListings
	white           *color.Color = color.New(color.FgWhite)
	boldRed         *color.Color = color.New(color.FgRed, color.Bold)
	boldGreen       *color.Color = color.New(color.FgGreen, color.Bold)
)

func main() {
	var userChoice string

	fmt.Println("Welcome to the house listing system!")

	for userChoice != "q" {
		fmt.Print("\n")
		fmt.Println("Enter 1 to view all listings.")
		fmt.Println("Enter 2 to add a listing.")
		fmt.Println("Enter q to quit the application.")
		fmt.Print("Enter an option: ")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			printErr(errInvalidInput)
		}

		userChoice = userInput[0:1]

		switch userChoice {
		case "1":
			getListings()
		case "2":
			addListing()
		case "q":
			fmt.Println("Thank you for using the house listing system. Goodbye!")
			os.Exit(0)
		}
	}

}

func getListings() {
	fmt.Print("\n")
	for i := 0; i < len(listings.Houses); i++ {
		if i%2 == 0 {
			fmt.Printf("%s\t%s\t%d Rooms\t$%d\n", listings.Houses[i].Address, listings.Houses[i].City, listings.Houses[i].NumOfRooms, listings.Houses[i].Price)
		} else {
			fmt.Printf("%s\t%s\t%d Rooms\t$%d\n", listings.Houses[i].Address, listings.Houses[i].City, listings.Houses[i].NumOfRooms, listings.Houses[i].Price)
		}
	}
}

func addListing() {
	var h House
	var err error

	fmt.Print("Enter an address: ")
	address, aErr := reader.ReadString('\n')
	if aErr != nil {
		printErr(errInvalidInput)
	}
	h.Address = address[:len(address)-1]

	fmt.Print("Enter a city: ")
	city, cErr := reader.ReadString('\n')
	if cErr != nil {
		printErr(errInvalidInput)
	}
	h.City = city[:len(city)-1]

	fmt.Print("Enter the number of rooms: ")
	rooms, rErr := reader.ReadString('\n')
	if rErr != nil {
		printErr(errInvalidInput)
	}
	r := rooms[:len(rooms)-1]
	h.NumOfRooms, err = strconv.Atoi(r)
	if err != nil {
		printErr(err)
	}

	fmt.Print("Enter the price: ")
	price, pErr := reader.ReadString('\n')
	if pErr != nil {
		printErr(errInvalidInput)
	}
	p := price[:len(price)-1]
	h.Price, err = strconv.Atoi(p)
	if err != nil {
		printErr(err)
	}

	listings.Houses = append(listings.Houses, h)
	printSuccess("Added a new listing.")
}

func printSuccess(s string) {
	boldGreen.Print("Success! ")
	white.Println(s)
}

func printErr(err error) {
	boldRed.Print("Error! ")
	white.Println(err)
	os.Exit(1)
}
