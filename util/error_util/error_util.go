package error_util

import (
	"github.com/gogf/gf/net/ghttp"
	"kong1/util/response_util"
)

func CommonCheckError(r *ghttp.Request , err error)  {
	if err != nil {
		response_util.JsonExit(r, 1, err.Error())
	}
}
