package boards

import (
	"github.com/zi-wei-dou-shu-gin/db/dao"
)

func NewBoardHandler(dao dao.DaoHandler) Handler {
	return handler{
		dao: dao,
	}
}
