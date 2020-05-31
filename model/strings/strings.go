package strings

import (
	"os"
	"strconv"

	"../../model/abstractions"
	"../../model/ciphers"
)

const (
	BR         = "=========================================="
	GREETINGS  = "Hello! =)"
	STARTINFO  = "Select a number of action:"
	ENC_MODE   = "encode"
	DEC_MODE   = "decode"
	EXIT_STR   = "Exit"
	PIG_LATIN  = "Pig Latin"
	VOWEL_CODE = "Vowel Code"
)

const (
	PLenc int = iota + 1
	VCenc
	VCdev
	Exit
)

var ActionStrings map[int]string
var ActionFuncs map[int]interface{}

func init() {
	ActionStrings = make(map[int]string)
	ActionStrings[PLenc] = "[" + strconv.Itoa(PLenc) + "] " + PIG_LATIN + " - " + ENC_MODE
	ActionStrings[VCenc] = "[" + strconv.Itoa(VCenc) + "] " + VOWEL_CODE + " - " + ENC_MODE
	ActionStrings[VCdev] = "[" + strconv.Itoa(VCdev) + "] " + VOWEL_CODE + " - " + DEC_MODE
	ActionStrings[Exit] = "[" + strconv.Itoa(Exit) + "] " + EXIT_STR

	var plEnc abstractions.Encoder = ciphers.PigLatin{Name: PIG_LATIN}
	var vcEnc abstractions.Encoder = ciphers.VowelCode{Name: VOWEL_CODE}
	var vcDec abstractions.Decoder = ciphers.VowelCode{Name: VOWEL_CODE}

	ActionFuncs = make(map[int]interface{})
	ActionFuncs[PLenc] = plEnc.Encode
	ActionFuncs[VCenc] = vcEnc.Encode
	ActionFuncs[VCdev] = vcDec.Decode
	ActionFuncs[Exit] = os.Exit
}
