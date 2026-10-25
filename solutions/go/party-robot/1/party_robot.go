package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	welcomeMessage := "Welcome to my party, "
	return welcomeMessage + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// panic("Please implement the HappyBirthday function")
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcome := fmt.Sprintf("Welcome to my party, %s!", name)
	tableInfo := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	neightborInfo := fmt.Sprintf("You will be sitting next to %s.", neighbor)

	return welcome + "\n" + tableInfo + "\n" + neightborInfo
}

// "Welcome to my party, Chihiro! You have been assigned to table 022. Your table is straight ahead, exactly 9.2 meters from here. You will be sitting next to Akachi Chikondi."
// "Welcome to my party, Chihiro! You have been assigned to table 022. Your table is straight ahead, exactly 9.2 meters from here. You will be sitting next to Akachi Chikondi."
