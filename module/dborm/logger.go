package dborm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/open-tdp/go-helper/logman"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type LogWrap struct {
	config *logger.Config
	logger *logman.Logger
}

func newLogger() logger.Interface {

	logger.Default = &LogWrap{
		logger: logman.Named("gorm"),
		config: &logger.Config{
			IgnoreRecordNotFoundError: false,
			SlowThreshold:             5 * time.Second,
		},
	}

	return logger.Default

}

func (lw LogWrap) LogMode(level logger.LogLevel) logger.Interface {

	lw.config.LogLevel = level
	return lw

}

func (lw LogWrap) Info(ctx context.Context, msg string, args ...any) {

	msg = fmt.Sprintf(msg, args...)
	lw.logger.Info(msg)

}

func (lw LogWrap) Warn(ctx context.Context, msg string, args ...any) {

	msg = fmt.Sprintf(msg, args...)
	lw.logger.Warn(msg)

}

func (lw LogWrap) Error(ctx context.Context, msg string, args ...any) {

	msg = fmt.Sprintf(msg, args...)
	lw.logger.Error(msg)

}

func (lw LogWrap) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	cfg := lw.config
	sql, rows := fc()
	elapsed := time.Since(begin)

	switch {
	case err != nil && (!errors.Is(err, gorm.ErrRecordNotFound) || !cfg.IgnoreRecordNotFoundError):
		lw.logger.Error("trace error", "error", err, "sql", sql, "rows", rows, "elapsed", elapsed)
	case elapsed > cfg.SlowThreshold && cfg.SlowThreshold != 0:
		slow := fmt.Sprintf("trace slow sql >= %v", cfg.SlowThreshold)
		lw.logger.Warn(slow, "sql", sql, "rows", rows, "elapsed", elapsed)
	default:
		lw.logger.Info("trace query", "sql", sql, "rows", rows, "elapsed", elapsed)
	}

}
