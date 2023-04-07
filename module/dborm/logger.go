package dborm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"tdp-cloud/helper/logman"
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

	lw.logger.Info(msg, logman.Any("data", args))

}

func (lw LogWrap) Warn(ctx context.Context, msg string, args ...any) {

	lw.logger.Warn(msg, logman.Any("data", args))

}

func (lw LogWrap) Error(ctx context.Context, msg string, args ...any) {

	lw.logger.Error(msg, logman.Any("data", args))

}

func (lw LogWrap) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	cfg := lw.config
	sql, rows := fc()
	elapsed := time.Since(begin)

	sqlF := logman.String("sql", sql)
	rowsF := logman.Int64("rows", rows)
	elapsedF := logman.Duration("elapsed", elapsed)

	switch {
	case err != nil && (!errors.Is(err, gorm.ErrRecordNotFound) || !cfg.IgnoreRecordNotFoundError):
		lw.logger.Error("trace error", logman.Any("error", err), sqlF, rowsF, elapsedF)
	case elapsed > cfg.SlowThreshold && cfg.SlowThreshold != 0:
		slow := fmt.Sprintf("trace slow sql >= %v", cfg.SlowThreshold)
		lw.logger.Warn(slow, sqlF, rowsF, elapsedF)
	default:
		lw.logger.Info("trace query", sqlF, rowsF, elapsedF)
	}

}
