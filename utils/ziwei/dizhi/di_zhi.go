package dizhi

func (d DiZhi) String() string {
	return []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}[d]
}

func ConvetDiZhi(dizhi string) DiZhi {
	return map[string]DiZhi{
		"子": 0, "丑": 1, "寅": 2, "卯": 3, "辰": 4, "巳": 5, "午": 6, "未": 7, "申": 8, "酉": 9, "戌": 10, "亥": 11,
	}[dizhi]
}
