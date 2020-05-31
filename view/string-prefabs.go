package view

import (
	"sort"

	strs "../model/strings"
)

var br string
var Greetings string
var StartInfo string
var actions string
var Errors map[int]string

func init() {
	actions = "\n"
	actionStrings := strs.ActionStrings

	keys := make([]int, 0)
	for k := range actionStrings {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, ind := range keys {
		actions += actionStrings[ind] + "\n"
	}

	br = Colors[Blue] + strs.BR + Colors[Reset]
	Greetings = Colors[Cyan] + strs.GREETINGS + Colors[Reset]
	StartInfo = Colors[Green] + strs.STARTINFO + actions + Colors[Reset]

	Errors = make(map[int]string)
	Errors[int(strs.PROC_IMPOS)] = Colors[Red] + strs.Errors[strs.PROC_IMPOS] + Colors[Reset]
	Errors[int(strs.NO_ACTION)] = Colors[Red] + strs.Errors[strs.NO_ACTION] + Colors[Reset]
}
