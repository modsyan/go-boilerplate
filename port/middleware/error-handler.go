package middleware

//func ErrorHandler() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Next()
//
//		if len(c.Errors) > 0 {
//			err := c.Errors.Last().Err
//
//			if locErr, ok := err.(*errors.LocalizedError); ok {
//				// Get localized message
//				message := locErr.Localize()
//
//				c.JSON(locErr.code, results.Error(
//					locErr.code,
//					message,
//					"",
//				))
//				return
//			}
//
//			// Handle regular errors
//			if httpErr, ok := err.(*errors.BaseError); ok {
//				c.JSON(httpErr.code, results.Error(
//					httpErr.code,
//					httpErr.Message,
//					httpErr.error,
//				))
//				return
//			}
//
//			// Default error response
//			c.JSON(500, results.Error(
//				500,
//				"Internal Server Error",
//				err.Error(),
//			))
//		}
//	}
//}
