package stars

func (sn StarName) String() string {
	return []string{
		// 十四主星
		"紫微", "天府", "天機", "太陽", "武曲", "天同", "廉貞", "太陰", "貪狼", "巨門", "天相", "天梁", "七殺", "破軍",
		// 年干諸系星
		"祿存", "擎羊", "陀羅", "天魁", "天鉞", "天官", "天福",
		"旬空", "截空",
		//　年支諸系星
		"天哭", "天虛", "龍池", "鳳閣", "紅鸞", "天喜", "孤辰", "寡宿", "解神", "破碎", "天馬", "大耗", "天德", "劫殺", "華蓋", "咸池", "天才", "天壽", "天空",
		// 月系星
		"左輔", "右弼", "天刑", "天姚", "天巫", "天月", "陰煞",
		// 時系諸星
		"文昌",
	}[sn]
}
