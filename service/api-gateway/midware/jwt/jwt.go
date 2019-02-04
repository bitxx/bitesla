package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"yaichain.com/aichain/AIChain-blockchain-prime/aichain-common/errs"
	"yaichain.com/aichain/AIChain-blockchain-prime/aichain-common/logger"
	"yaichain.com/aichain/AIChain-blockchain-prime/aichain-common/result"
	"yaichain.com/aichain/AIChain-blockchain-prime/aichain-common/util"
	"yaichain.com/aichain/AIChain-blockchain-prime/aichain-core/front/aichain-main/conf"
	"yaichain.com/aichain/AIChain-blockchain-prime/aichain-core/front/aichain-main/constants"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		logger.Info("head:", c.Request.Header)
		logger.Info("body", c.Request.Body)

		token := c.GetHeader(constants.Token)
		code, err := validData(token)

		if code != errs.Success {
			res := result.NewResult()
			defer c.JSON(http.StatusOK, res)
			logger.Error(err)
			res.Code = code
			res.Msg = errs.GetMsg(code)
			res.Data = data
			c.Abort()
		}

		c.Next()
	}
}

func validData(token string) (code int, err error) {
	if token == "" {
		return errs.TokenEmptyErr, nil
	}

	claims, err := util.ParseToken(token, conf.CurrentConfig.ServerConf.JwtSecret)
	if err != nil {
		return errs.TokenErr, err
	}
	if time.Now().Unix() > claims.ExpiresAt {
		return errs.TokenExpire, nil
	}
	return errs.Success, nil
}
