package ds

import "fmt"

type Song struct {
	Name     string
	Artist   string
	Previous *Song // pointer, link to the previous song
	Next     *Song // pointer,link to the next song
}

type Playlist struct {
	Name       string
	Head       *Song // pointer (link to the head node)
	Tail       *Song // pointer (link to the tail node)
	NowPlaying *Song // current song
}

func CreatePlaylist(name string) *Playlist {
	return &Playlist{
		Name: name,
	}
}

func (p *Playlist) AddSong(name, artist string) error {
	s := &Song{
		Name:   name,
		Artist: artist,
	}
	if p.Head == nil {
		p.Head = s
	} else {
		//currentNode := p.Head
		//for currentNode.Next != nil {
		//	currentNode = currentNode.Next
		//}
		//currentNode.Next = s
		currentNode := p.Tail
		currentNode.Next = s
		s.Previous = p.Tail
	}
	p.Tail = s
	return nil
}

func (p *Playlist) ShowAllSongs() error {
	currentNode := p.Head
	if currentNode.Next == nil {
		fmt.Println("Playlist is empty")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

func (p *Playlist) StartPlaying() *Song {
	p.NowPlaying = p.Head
	return p.NowPlaying
}

func (p *Playlist) NextSong() *Song {
	p.NowPlaying = p.NowPlaying.Next
	return p.NowPlaying
}

func (p *Playlist) PreviousSong() *Song {
	p.NowPlaying = p.NowPlaying.Previous
	return p.NowPlaying
}

// add a song to playlist
// show all songs
// start playing a song
// next/prev song on the playlist
