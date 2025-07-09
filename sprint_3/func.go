package sprint_3

import "fmt"

func printFriendsCount(friendsCount int) {
	if friendsCount == 0 {
		fmt.Println("У тебя нет друзей")
	} else if friendsCount == 1 {
		fmt.Println("У тебя", friendsCount, "друг")
	} else if friendsCount >= 2 && friendsCount <= 4 {
		fmt.Println("У тебя", friendsCount, "друга")
	} else if friendsCount >= 5 && friendsCount < 20 {
		fmt.Println("У тебя", friendsCount, "друзей")
	} else {
		fmt.Println("Ого, сколько у тебя друзей! Целых ", friendsCount)
	}
}

func sprint_3() {
	// Ниже закончите цикл, в котором будет вызываться функция printFriendsCount
	// с аргументом от 0 до 20 включительно
	for i := 0; i <= 20; i++ {
		// вызов функции printFriendsCount
		printFriendsCount(i)
	}

}
