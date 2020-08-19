package main

import (
	"bufio"
	"fmt"
	"golang-/day11/mlib"
	"golang-/day11/mp"
	"os"
	"strconv"
	"strings"
)

var lib *mlib.MusicManager
var id int = 1

//管理模块句柄实现
func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&mlib.MusicEntry{strconv.Itoa(id),
				tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognied lib commond:", tokens[1])
	}
}

//播放模块句柄实现
func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e, _ := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The Music", tokens[1], "does not exist")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	//打印操作菜单
	fmt.Println(`
		Enter following commands to control the player:
		lib list -- View the existing music lib
		lib add <name><artist><source><type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> -- Play the specified music
	`)
	lib = mlib.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter Command-> ")

		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		//输入q或者e时退出播放器
		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
