package tiangan

func (t TianGan) String() string {
	return []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}[t]
}

func ConvetTainGan(tainGan string) TianGan {
	return map[string]TianGan{
		"甲": 0, "乙": 1, "丙": 2, "丁": 3, "戊": 4, "己": 5, "庚": 6, "辛": 7, "壬": 8, "癸": 9,
	}[tainGan]
}
