package session

import (
	"github.com/labstack/echo"
)

const ECHO_SESSION_KEY = "echo-session"

func Echo(c *echo.Context) Store {
	store, ok := c.Get(ECHO_SESSION_KEY).(Store)
	if !ok {
		panic("get session store failed")
	}

	return store
}

func EchoSessioner(options ...Options) echo.MiddlewareFunc {
	opt := prepareOptions(options)
	manager, err := NewManager(opt.Provider, &opt.Config)
	if err != nil {
		panic(err)
	}
	go manager.GC()

	return func(h echo.HandlerFunc) echo.HandlerFunc {

		return func(c *echo.Context) error {

			sess, err := manager.SessionStart(c.Response(), c.Request())
			if err != nil {
				return err
			}

			s := store{
				RawStore: sess,
				Manager:  manager,
			}

			c.Set(ECHO_SESSION_KEY, s)

			if err := h(c); err != nil {
				c.Error(err)
			}

			sess.SessionRelease(c.Response())

			return nil
		}
	}
}
