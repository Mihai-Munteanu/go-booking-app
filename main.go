package main

import (
	// in case we move helper file into the helper folder, it will be considered a different package and therefore should be imported as a package
	// "booking-app/helper"
	"fmt"
	"sync"
	"time"
	// no more use of strconv package
	// "strconv"
	// no more use of strings package
	// "strings"
);

const conferenceTickets int = 50
var conferenceName string = "Go conference"
var remainingTickets uint = 50
// we change the slice type from string to map, we don't need curly brackets, but instead we need to use function make which should also include slice type and the next argument should be the initial number of entries
// var bookings = []string{}
// var bookings = make([]map[string]string, 0 )
// we change the from maps type to UserData struct type
var bookings = make([]UserData, 0 )


// strunct allow us to create key value pairs with different types.
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

// this variable is created in order for main thread to wait other treads
var wg = sync.WaitGroup{}

func main () {

	greedUser()

	// we delete the for loop in order to view that the main tread does not wait for other threads. In order to wait other treads we should create a WaitGroup
	// for {
		firstName, lastName, email, userTickets := getUserInput()

		// we have used helper.ValidateUserInput as in this scenario we consider helper file to be into helper package and thus we imported the function ValidateUserInput from that package. Please note that the function is written with caps lock, this means that a function is public and can be accessed bu other package. Same apply for the build in packages fmt, strings
		// isValidName, isValidEmail,  isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		isValidName, isValidEmail,  isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail &&  isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			// this Add function  from wait group add the following function which you would like to be waited
			wg.Add(1)
			// by including the go in front of the function we make the function to work concurrent
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstName()
			fmt.Printf("The first names of bookings are: %v \n", firstNames)


			noTicketsAvailable := remainingTickets == 0

			if noTicketsAvailable {
				// end the program
				fmt.Println("Conference is booked out, come back next year")
				// break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name is to short")
			}

			if !isValidEmail {
				fmt.Println("email address does not contain @")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}

			// continue
		}
		// at the end of the main thread we should include the function wait from wg in order to wait for all added function into wg
		wg.Wait()
	}
// }

func greedUser() {
	fmt.Printf ("Welcome to our %v bookings application\n", conferenceName)
	fmt.Printf ("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println ("Get your ticket here to attend")
}

func getFirstName() []string {
	firstNames := []string{}

	// we change the iteration from a slice of string into an iteration of maps
	// for _, booking := range bookings {
	// 	var names = strings.Fields(booking)
	// 	firstNames = append(firstNames, names[0])
	// }
	// change from iteration of map into an iteration of struct
	// for _, booking := range bookings {
	// 	firstNames = append(firstNames, booking["firstName"])
	// }
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask for user input
	fmt.Println("Please intern your first name")
	fmt.Scan(&firstName)

	fmt.Println("Please intern your last name")
	fmt.Scan(&lastName)

	fmt.Println("Please intern your email address")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//  create a map for a user the variable user data is type map with string type for key and string type for values. in order to crete a empty map, we should wrap the type into the make function
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// // as user tickets is an uint and not a string we can not include it into the map as the map can only accept the same type for all its key value pair. Thus we convert the uint into the string
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	// change the userData from map type to struct type
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)


	fmt.Printf("Thank you %v %v for bookings %v tickets. You will receive an email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lasName string, email string) {
	// we simulate the sleep 10 seconds
	time.Sleep(10 * time.Second)
	// if we would like to save the string we should use the method fmt.Sprintf
	var ticket = fmt.Sprintf("%v tickests for %v %v ", userTickets, firstName, lasName)
	fmt.Println("########################")
	fmt.Printf("sending ticket:\n %v \n to email address to %v\n", ticket, email)
	fmt.Println("########################")

	// here we instruct the wg that the function is done and should be removed from the waiting list of the main thread
	wg.Done()
}





func oldCode () {

	// fmt.Printf ("ConferenceTickets is %T remainingTickets is %T  and conferenceName is %T  \n", conferenceTickets, remainingTickets, conferenceName)


	// array
	// var bookings [50]string
	// slice

	// array
	// bookings[0] = firstName + " " + lastName
	// slice



	// fmt.Printf("hole slice:  %v \n", bookings)
	// fmt.Printf("first slice:  %v \n", bookings[0])
	// fmt.Printf("type of slice:  %T \n", bookings)
	// fmt.Printf("length of an slice:  %v \n", len(bookings))

	// fmt.Printf("hole array:  %v \n", bookings)
	// fmt.Printf("first array:  %v \n", bookings[0])
	// fmt.Printf("type of array:  %T \n", bookings)
	// fmt.Printf("length of an array:  %v \n", len(bookings))

	// fmt.Printf("Thank you %v %v for bookings %v tickets. You will receive an email at %v \n", firstName, lastName, userTickets, email)
	// fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)

	// fmt.Printf("User %v booked %v tickets", firstName, userTickets)
}