//package main
package mlib
import "errors"

type MusicEntry struct {
    Id string
    Name string
    Artist string
    Source string
    Type string
}

type MusicManager struct {
    musics []MusicEntry
}

func NewMusicManager() * MusicManager{
	return &MusicManager{make([]MusicEntry,0)}
}

func (m* MusicManager) Len() int{
	return len(m.musics)
}

func (m* MusicManager) Get(index int)(music* MusicEntry, err error){
	if index < 0 || index >= m.Len(){
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m* MusicManager) Find(name string) *MusicEntry{
	if m.Len() == 0 {
		return nil
	}

	for _, music := range m.musics{
		if music.Name == name{
			return &music
		}
	}

	return nil
}

func (m* MusicManager) Add(music * MusicEntry){
	m.musics = append(m.musics,*music)
}

func (m* MusicManager) Remove(index int) * MusicEntry{
	if index < 0 || index >= m.Len(){
		return nil
	}
	removedMusic := &m.musics[index]
	//删除元素
	m.musics = append(m.musics[:index],m.musics[index+1:]...)
	return removedMusic
}