package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mochi-yu/webapp-auth-practice/ent"
)

func Transaction(db *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tx, err := db.Tx(ctx)
		if err != nil {
			ctx.JSON(500, "failed to create transaction")
			ctx.Abort()
			return
		}

		defer func() {
			if 400 <= ctx.Writer.Status() {
				tx.Rollback()
				return
			}
			tx.Commit()
		}()

		ctx.Set("tx", tx)
		ctx.Next()
	}
}
