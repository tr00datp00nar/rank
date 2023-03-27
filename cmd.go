package rank

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `rank`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_rank),
	Description: help.D(_rank),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		fmt.Print("Items to rank:\n\n")

		r := []string{}
		for {
			item := readline()
			if item == "" {
				break
			}
			r = append(r, item)
		}
		check(scanner.Err())
		BinaryInsertionSort(r, LessPrompt)
		for i, s := range r {
			fmt.Printf("%d. %s\n", i+1, s)
		}
		return nil
	},
}
