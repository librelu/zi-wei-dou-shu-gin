package mingju

type MingJuType int

const (
	Jin = MingJuType(iota)
	Shui
	Huo
	Tu
	Mu
)

var JuShuMap = map[MingJuType]uint{
	Jin:  4,
	Shui: 2,
	Huo:  6,
	Tu:   5,
	Mu:   3,
}
