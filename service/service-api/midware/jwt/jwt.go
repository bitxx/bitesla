package jwt

import (
	"bytes"
	"github.com/bitxx/bitesla/common/errs"
	"github.com/bitxx/bitesla/common/logger"
	"github.com/bitxx/bitesla/common/util/jwt"
	"github.com/bitxx/bitesla/service/service-api/conf"
	"github.com/bitxx/bitesla/service/service-api/constants"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		token := c.GetHeader(constants.Token)
		userId, code, err := validData(token)

		if code != errs.Success {
			logger.Error(err)

			//TODO 此处代码重复，有空再调
			res := result.NewResult()
			defer c.JSON(http.StatusOK, res)
			res.Code = code
			res.Msg = errs.GetMsg(code)
			res.Data = data
			c.Abort()
			return
		}

		bts, _ := c.GetRawData()
		reqStr := string(bts)
		if strings.Contains(reqStr, constants.CurrentLoginUserID) {
			code = errs.RequestHeadCurrUserIdErr

			//TODO 此处代码重复，有空再调
			res := result.NewResult()
			defer c.JSON(http.StatusOK, res)
			res.Code = code
			res.Msg = errs.GetMsg(code)
			res.Data = data
			c.Abort()
			return
		}

		tmp := strings.Replace(reqStr, "{", "{\""+constants.CurrentLoginUserID+"\":"+strconv.FormatInt(userId, 10)+",", 1)
		c.Request.Body = ioutil.NopCloser(bytes.NewBufferString(tmp))
		c.Next()
	}
}

func validData(token string) (userId int64, code int, err error) {
	if token == "" {
		return 0, errs.TokenEmptyErr, nil
	}

	claims, err := jwt.ParseToken(token, conf.CurrentConfig.ServerConf.JwtSecret)
	logger.Info("当前访问用户的id:", claims.UserId)
	if err != nil {
		return 0, errs.TokenErr, err
	}
	if time.Now().Unix() > claims.ExpiresAt {
		return 0, errs.TokenExpire, nil
	}
	return claims.UserId, errs.Success, nil
}
