package view

type Color int

var Colors map[Color]string

const (
	Reset Color = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	Gray
	White
)

func init() {
	Colors = make(map[Color]string)
	Colors[Reset] = "\033[0m"
	Colors[Red] = "\033[31m"
	Colors[Green] = "\033[32m"
	Colors[Yellow] = "\033[33m"
	Colors[Blue] = "\033[34m"
	Colors[Purple] = "\033[35m"
	Colors[Cyan] = "\033[36m"
	Colors[Gray] = "\033[37m"
	Colors[White] = "\033[97m"
}
