package startype

func (sn StarType) String() string {
	return []string{
		"十四主星",
		"年干系諸星",
		"四化",
		"年支系諸星",
		"月系星",
		"時系諸星",
		"身命",
	}[sn]
}
