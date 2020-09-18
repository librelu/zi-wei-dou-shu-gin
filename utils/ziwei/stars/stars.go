package stars

func (sn StarName) String() string {
	return []string{"紫微", "天府", "天機", "太陽", "武曲", "天同", "廉貞", "太陰", "貪狼", "巨門", "天相", "天梁", "七殺", "破軍"}[sn]
}
