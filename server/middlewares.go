package server

import "fmt"

type jtMiddlewareFunc func(jtHandlerFunc) jtHandlerFunc

func mAuth() jtMiddlewareFunc {
	return func(next jtHandlerFunc) jtHandlerFunc {
		return func(c jtContext) {
			if !c.client.isAuthorized {
				c.sendResult(errorNoAuth)
				return
			}
			next(c)
		}
	}
}

func mArgsEqual(count int) jtMiddlewareFunc {
	return func(next jtHandlerFunc) jtHandlerFunc {
		return func(c jtContext) {
			if len(c.args) != count {
				c.sendResult(fmt.Sprintf(errorWrongArguments, c.command))
				return
			}
			next(c)
		}
	}
}

func mArgsMoreThan(count int) jtMiddlewareFunc {
	return func(next jtHandlerFunc) jtHandlerFunc {
		return func(c jtContext) {
			if len(c.args) < count {
				c.sendResult(fmt.Sprintf(errorWrongArguments, c.command))
				return
			}
			next(c)
		}
	}
}

func mArgsCountEven() jtMiddlewareFunc {
	return func(next jtHandlerFunc) jtHandlerFunc {
		return func(c jtContext) {
			if len(c.args)%2 != 0 {
				c.sendResult(fmt.Sprintf(errorWrongArguments, c.command))
				return
			}
			next(c)
		}
	}
}
