package dizhi

type DiZhi uint

const (
	Zi = DiZhi(iota)
	Chou
	Yin
	Mao
	Chen
	Si
	Wu
	Wei
	Shen
	You
	Xu
	Hai
)

var Time = map[int]DiZhi{
	23: Zi,
	24: Zi,
	0:  Zi,
	1:  Chou,
	2:  Chou,
	3:  Yin,
	4:  Yin,
	5:  Mao,
	6:  Mao,
	7:  Chen,
	8:  Chen,
	9:  Si,
	10: Si,
	11: Wu,
	12: Wu,
	13: Wei,
	14: Wei,
	15: Shen,
	16: Shen,
	17: You,
	18: You,
	19: Xu,
	20: Xu,
	21: Hai,
	22: Hai,
}
