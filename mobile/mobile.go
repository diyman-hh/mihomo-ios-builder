package mobile

import (
	"os"
	"runtime/debug"
	"time"

	C "github.com/metacubex/mihomo/constant"
	"github.com/metacubex/mihomo/hub"
	"github.com/metacubex/mihomo/hub/executor"
	"github.com/metacubex/mihomo/log"
)

// Start starts the mihomo kernel with the given home directory and configuration string.
// It returns an empty string on success, or an error message on failure.
func Start(homeDir string, configContent string) string {
	// Set Home Directory
	C.SetHomeDir(homeDir)

	// Ensure home directory exists
	if _, err := os.Stat(homeDir); os.IsNotExist(err) {
		_ = os.MkdirAll(homeDir, 0777)
	}

	// Initialize Logger
	// Default log level might be info, can be configured in configContent usually
	// But we ensure basic setup
	
	// Recover from panics during startup
	defer func() {
		if r := recover(); r != nil {
			log.Errorln("Recovered from panic in Start: %v\nStack: %s", r, debug.Stack())
		}
	}()

	// Parse configuration and start
	// We pass no extra options for now, relying on config content
	err := hub.Parse([]byte(configContent))
	if err != nil {
		return err.Error()
	}

	return ""
}

// Stop shuts down the mihomo kernel.
func Stop() {
	executor.Shutdown()
	// Force GC to free memory?
	// runtime.GC()
	// Wait a bit?
	time.Sleep(100 * time.Millisecond)
}
