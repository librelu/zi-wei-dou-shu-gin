package utils

import (
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/dizhi"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/miaoxian"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/stars"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/startype"
	"github.com/zi-wei-dou-shu-gin/utils/ziwei/tiangan"
)

func SetQingYang(board *Board, tianGan tiangan.TianGan) *Board {
	qingYangMap := []dizhi.DiZhi{
		dizhi.Mao,
		dizhi.Chen,
		dizhi.Wu,
		dizhi.Wei,
		dizhi.Wu,
		dizhi.Wei,
		dizhi.You,
		dizhi.Xu,
		dizhi.Zi,
		dizhi.Chou,
	}
	qingYangLocation := qingYangMap[int(tianGan)]
	board.Blocks[qingYangLocation].Stars = append(board.Blocks[qingYangLocation].Stars, &Star{
		Name:     stars.QingYang.String(),
		StarType: startype.RightFuXing.String(),
		MiaoXian: getQingYangMiaoXian(int(qingYangLocation)).String(),
	})
	board.StarsMap[stars.QingYang] = int(qingYangLocation)
	return board
}

func getQingYangMiaoXian(index int) miaoxian.MiaoXian {
	miaoXianMap := []miaoxian.MiaoXian{
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.None,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.None,
		miaoxian.Ping,
		miaoxian.Miao,
		miaoxian.None,
		miaoxian.Xian,
		miaoxian.Miao,
		miaoxian.None,
	}
	return miaoXianMap[index]
}
