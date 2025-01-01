package middleware

import (
	"company-name/pkg/localization"
	"github.com/gin-gonic/gin"
)

func LocalizationMiddleware(c *gin.Context) {
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = "en"
	}
	localization.SetLang(lang)
	c.Next()
	return
}
