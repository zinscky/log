package log

import (
	"fmt"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

func TestLog(t *testing.T) {
	log := New(Info, "test-app")

	log.Debug("DEBUG LOG")
	log.Info("INFO LOG")
	log.Warn("WARN LOG")
	log.Error("ERROR LOG")
	require.NotEmpty(t, log.LogStr)
	require.Contains(t, log.LogStr[0], "INFO LOG")
	require.Contains(t, log.LogStr[1], "WARN LOG")
	require.Contains(t, log.LogStr[2], "ERROR LOG")
	require.Equal(t, 3, len(log.LogStr))

	fmt.Println(log.String())
}

func TestFileLogger(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().Header.Set("X-Request-ID", "12345")
	ctx.Request().Header.Set("X-App-Name", "fiber-fn")
	log := NewLogger(ctx, Debug)
	log.Debug("DEBUG LOG")
	log.Info("INFO LOG")
	log.Warn("WARN LOG")
	log.Error("ERROR LOG")
	logs, err := os.ReadFile("fiber-fn-12345.log")
	require.Nil(t, err)
	require.NotEmpty(t, logs)
	require.Contains(t, string(logs), "DEBUG LOG")
	fmt.Println(string(logs))
	app.ReleaseCtx(ctx)
	os.Remove("fiber-fn-12345.log")
}
