package main

import (
	"fmt"
	"log"

	"go.i3wm.org/i3/v4"
)

func main() {
	tree, _ := i3.GetTree()
	recv := i3.Subscribe(i3.WindowEventType)
	set_width := fmt.Sprintf("resize set width %d px", int(float64(tree.Root.Rect.Width)*.618))
	set_height := fmt.Sprintf("resize set height %d px", int(float64(tree.Root.Rect.Height)*.618))
	for recv.Next() {
		_, err := i3.RunCommand(set_width)
		if err != nil {
			log.Printf("failed to set with: %s\n", err)
		}
		_, err = i3.RunCommand(set_height)
		if err != nil {
			log.Printf("failed to set height: %s\n", err)
		}

	}
	log.Fatal(recv.Close())
}
