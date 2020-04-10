package main

import (
	"fmt"
	"github.com/Gidraff/algorithms-golang/ds"
	"github.com/Gidraff/algorithms-golang/numeric"
	"github.com/Gidraff/algorithms-golang/sorting"
)

func main() {
	numList := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(numeric.NumInList(numList, 5))

	fmt.Print("Fizz Buzz Algo")
	fmt.Println()
	numeric.FizzBuzz(15)

	fmt.Println("DecToBase =>:")
	fmt.Println(numeric.DecToBase(10, 2))

	fmt.Println("BaseToBase =>:")
	fmt.Println(numeric.BaseToBase("10101", 2, 10))

	fmt.Println("FindTwoThatSum =>:")
	fmt.Println(numeric.FindTwoThatSum([]int{1, 2, 3, 4, 5}, 9))

	fmt.Println("Factors of Given number =>:")
	fmt.Println(numeric.Factor([]int{3, 5, 7, 9}, 9))

	fmt.Println("Fibonacci Sequence:")
	fmt.Println(numeric.Fibonacci(6))

	fmt.Println("GCD =>:")
	fmt.Println(numeric.GCD(20, 15))

	fmt.Println("Linked lists: ==>")
	fmt.Println("Create a playlist")
	playlist := ds.CreatePlaylist("Raff-Best")

	fmt.Print("Adding songs to playlist...\n\n")
	playlist.AddSong("Magode choke", "Zzero")
	playlist.AddSong("Suzzana", "Sauti sol")
	playlist.AddSong("whiley flow", "stormyz")
	playlist.AddSong("Ill mind", "Hopsin")
	playlist.AddSong("Hello", "Adele")

	playlist.ShowAllSongs()
	fmt.Println()

	playlist.StartPlaying()
	fmt.Printf("Now playing %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)

	playlist.NextSong()
	playlist.NextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)

	playlist.PreviousSong()
	fmt.Println("Changing tn previous song...")
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)

	playlist.PreviousSong()
	fmt.Println("Changing tn previous song...")
	fmt.Printf("Now playing: %s by %s\n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
	fmt.Println("Linked list implementation")
	link := &ds.List{}
	link.Insert(1)
	link.Insert(3)
	link.Insert(3)
	fmt.Printf("Head: %v\n", link.Head.Key)
	fmt.Printf("Tail: %v\n", link.Tail.Key)
	link.Display()
	link.Reverse()

	fmt.Println("BubbleSort")
	toBeSorted := []int{23, 13, 34, 37, 16, 8, 9}
	sorting.BubbleSortInt(toBeSorted)

}
