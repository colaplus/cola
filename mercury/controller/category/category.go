package category

import (
	"github.com/gin-gonic/gin"
	"mercury/dal/db"
	"mercury/util"
)

func GetCategoryListHandle(c *gin.Context) {
	categoryList, err := db.GetCategoryList()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	util.ResponseSuccess(c, categoryList)
}
