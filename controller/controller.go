package controller

import (
	"strconv"

	strs "../model/strings"
)

var result string
var isErr bool

func Do(text string, action string) (string, bool) {
	ind, err := strconv.Atoi(action)
	if err != nil {
		return "", true
	}
	a := strs.ActionFuncs[ind]
	result, isErr = a.(func(string) (string, bool))(text)
	return result, isErr
}

func ShouldContinue(action string) (bool, int) {
	ind, err := strconv.Atoi(action)
	if err != nil {
		return false, int(strs.NO_ACTION)
	}
	if len(strs.ActionFuncs) == ind {
		strs.ActionFuncs[ind].(func(int))(0)
	}
	if _, ok := strs.ActionFuncs[ind]; !ok {
		return false, int(strs.NO_ACTION)
	}
	return true, 0
}
