package ask

import (
	"github.com/gin-gonic/gin"
	"mercury/common"
	"mercury/filter"
	"mercury/id_gen"
	"mercury/util"
)

func QuestionSubmitHandle(c *gin.Context) {
	var question common.Question
	err := c.BindJSON(&question)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	_, hit := filter.Replace(question.Caption, "***")
	if hit {
		util.ResponseError(c, util.ErrCodeCaptionHit)
		return
	}

	_, hit = filter.Replace(question.Content, "***")
	if hit {
		util.ResponseError(c, util.ErrCodeContentHit)
		return
	}

	question.QuestionId, err = id_gen.GetId()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
	}

	util.ResponseSuccess(c, nil)
	return
}
