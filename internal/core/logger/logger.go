package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/0ScPro0/go-todolist/internal/core/config"
)

// Logger wraps the zap logger with additional file handle management
// This provides structured logging capabilities with both console and file outputs
type Logger struct {
	*zap.Logger // Embedded zap logger for all logging methods (Info, Debug, Error, etc.)

	file *os.File // File handle to the log file, needed for proper cleanup
}

func FromContext(ctx context.Context) *Logger {
	log, ok := ctx.Value("log").(*Logger)
	if !ok {
		panic("No logger in context")
	}
	
	return log
}

// NewLogger creates and initializes a new logger instance based on the application configuration
// It sets up dual logging: to stdout (console) and to a timestamped log file
// Returns the initialized logger or an error if initialization fails
func NewLogger(cfg *config.Config) (*Logger, error) {
	// Create an atomic level configuration that can be changed at runtime
	zapLvl := zap.NewAtomicLevel()
	
	// Parse the log level from config (e.g., "debug", "info", "error")
	// UnmarshalText converts string representation to zap's internal level type
	if err := zapLvl.UnmarshalText([]byte(cfg.Logger.Level)); err != nil {
		return nil, fmt.Errorf("Unmarshal log level: %w", err)
	}

	// Create the log directory if it doesn't exist
	// 0755 permissions: owner can read/write/execute, group/others can read/execute
	if err := os.MkdirAll(cfg.Logger.Folder, 0755); err != nil {
		return nil, fmt.Errorf("mkdir log folder: %w", err)
	}

	// Generate a unique timestamp for the log file name
	// Using UTC time and ISO-like format with colons replaced by hyphens for filesystem compatibility
	// Format includes microseconds for uniqueness in rapid succession log creation
	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logFilePath := filepath.Join(
		cfg.Logger.Folder,
		fmt.Sprintf("%s.log", timestamp),
	)

	// Create or open the log file with write-only permissions
	// O_CREATE: create file if doesn't exist
	// O_WRONLY: write-only mode
	// 0644 permissions: owner can read/write, group/others can read only
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Open log file: %w", err)
	}

	// Create encoder configuration for human-readable log format (development style)
	// This produces logs like: "2024-01-01T12:00:00.000000 INFO some message"
	zapConfig := zap.NewDevelopmentEncoderConfig()
	
	// Override the default time format to include microseconds for precise timing
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")

	// Create a console encoder that produces human-readable output (not JSON)
	zapEncoder := zapcore.NewConsoleEncoder(zapConfig)

	// Create a multi-writer core that writes to both stdout and the log file simultaneously
	// zapcore.NewTee allows writing to multiple destinations with the same encoder and level
	core := zapcore.NewTee(
		// First output: standard output (console) for real-time visibility
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), zapLvl),
		// Second output: log file for persistence and later analysis
		zapcore.NewCore(zapEncoder, zapcore.AddSync(logFile), zapLvl),
	)

	// Create the zap logger with caller information enabled
	// AddCaller adds file name and line number to each log entry for debugging
	zapLogger := zap.New(core, zap.AddCaller())

	// Wrap the zap logger with our file handle for proper cleanup
	return &Logger{
		Logger: zapLogger,
		file:   logFile,
	}, nil
}

func (l *Logger) With (field ...zap.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(field...),
		file: l.file,
	}
}

// Close properly shuts down the logger by closing the log file
// This ensures all buffered log entries are flushed to disk
// Should be called defer logger.Close() when the application shuts down
func (l *Logger) Close() {
	if err := l.file.Close(); err != nil {
		// Fallback to fmt.Println since the logger might be unavailable during shutdown
		// This ensures we don't silently fail to report the closure error
		fmt.Println("Failed to close application logger:", err)
	}
}