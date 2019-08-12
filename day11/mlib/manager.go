//package main
package mlib

import (
	"errors"
	"fmt"
)

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= m.Len() {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) (*MusicEntry, int) {
	if m.Len() == 0 {
		return nil, -1
	}

	for i, music := range m.musics {
		if music.Name == name {
			return &music, i
		}
	}

	return nil, -1
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= m.Len() {
		return nil
	}
	removedMusic := &m.musics[index]
	//删除元素
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return removedMusic
}

//通过歌名删除歌曲
func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	removedMusic, index := m.Find(name)
	if removedMusic == nil || index == -1 {
		fmt.Println("Want to delete the song does not exist")
		return nil
	}

	rm := m.Remove(index)
	fmt.Printf("%p", rm, "  ", "\n%p", removedMusic)
	return removedMusic
}
