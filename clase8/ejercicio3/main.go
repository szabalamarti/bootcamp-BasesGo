package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Client struct {
	Name        string
	Id          int
	PhoneNumber string
	Home        string
}

type ClientList struct {
	Clients []Client
	lastId  int
}

// Custom errors
var ErrEmptyFields = errors.New("client has empty fields")
var ErrClientExists = errors.New("client already exists")
var ErrInvalidData = errors.New("invalid data in file")

// Verify that no fields in the client are empty
func (c Client) verifyNonZero() error {
	if c.Name == "" || c.PhoneNumber == "" || c.Home == "" {
		return ErrEmptyFields
	}
	return nil
}

// String representation of ClientList
func (c ClientList) String() string {
	return fmt.Sprintf("Clients: %v\nLast ID: %d\n", c.Clients, c.lastId)
}

// Add a client to the list
func (c *ClientList) addClient(client Client) {
	client.Id = c.lastId
	c.Clients = append(c.Clients, client)
	c.lastId++
}

// Load clients from a file
func loadClients(fileName string) (clientList ClientList, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			err = fmt.Errorf("%v", r)
		}
	}()

	clientList = ClientList{
		Clients: []Client{},
		lastId:  0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		if len(values) != 4 {
			panic(ErrInvalidData)
		}

		id, err := strconv.Atoi(values[1])
		if err != nil {
			panic(ErrInvalidData)
		}

		client := Client{
			Name:        values[0],
			Id:          id,
			PhoneNumber: values[2],
			Home:        values[3],
		}

		clientList.Clients = append(clientList.Clients, client)
		if id > clientList.lastId {
			clientList.lastId = id
		}
	}

	if err := scanner.Err(); err != nil {
		panic(ErrInvalidData)
	}

	return
}

// Check if a client exists in the list
func clientExists(clientList ClientList, client Client) error {
	for _, c := range clientList.Clients {
		if c.Name == client.Name && c.Id == client.Id && c.PhoneNumber == client.PhoneNumber && c.Home == client.Home {
			return ErrClientExists
		}
	}
	return nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Load clients from file
	clientList, err := loadClients("ejercicio3/customers.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(clientList)

	// Create a new client
	client := Client{
		Name:        "Juan",
		Id:          0,
		PhoneNumber: "123456789",
		Home:        "Calle 123",
	}

	// Check if the client already exists
	if err := clientExists(clientList, client); err != nil {
		panic(err)
	}

	// Verify that the client has no empty fields
	if err := client.verifyNonZero(); err != nil {
		panic(err)
	}
	fmt.Println("New client has no empty fields")

	// Add the client to the list
	clientList.addClient(client)
	fmt.Println(clientList)
}
