package gong

func (g Gong) String() string {
	return []string{"命宮", "父母", "福德", "田宅", "事業", "交友", "遷移", "疾病", "財帛", "子女", "夫妻", "兄弟"}[g]
}
