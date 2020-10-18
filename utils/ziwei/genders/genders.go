package genders

func (g Gender) String() string {
	return []string{"男", "女", "陽男", "陰男", "陽女", "陰女"}[g]
}
