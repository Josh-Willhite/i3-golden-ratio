package main

import (
	"fmt"
	"log"
	"strings"

	"go.i3wm.org/i3/v4"
)

func main() {
	recv := i3.Subscribe(i3.WindowEventType)
	for recv.Next() {
		ev := recv.Event().(*i3.WindowEvent)
		log.Printf("change: %s", ev.Change)
		log.Printf("name: %s", ev.Container.Name)

	}
	log.Fatal(recv.Close())

	// Focus or start Google Chrome on the focused workspace.
	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("TREE-NAME: %s\n", tree.Root.Name)

	ws := tree.Root.FindFocused(func(n *i3.Node) bool {
		return n.Type == i3.WorkspaceNode
	})
	if ws == nil {
		log.Fatalf("could not locate workspace")
	}

	// res := ws.FindChild(func(n *i3.Node) bool {
	// 	return strings.ContainsAny(n.Name, "emacs")
	// })

	chrome := ws.FindChild(func(n *i3.Node) bool {
		return strings.ContainsAny(n.Name, "calibre")
		// return strings.HasSuffix(n.Name, "- Emacs")
	})

	if chrome != nil {
		_, err = i3.RunCommand(fmt.Sprintf(`[con_id="%d"] focus`, chrome.ID))
	} else {
		_, err = i3.RunCommand(`exec google-chrome`)
	}
	if err != nil {
		log.Fatal(err)
	}
}
