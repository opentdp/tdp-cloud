package dborm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"tdp-cloud/helper/logman"
)

type LogWrap struct {
	config *logger.Config
	logger *zap.Logger
}

func NewLogger() logger.Interface {

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

	lw.logger.Info(msg, zap.Any("data", args))

}

func (lw LogWrap) Warn(ctx context.Context, msg string, args ...any) {

	lw.logger.Warn(msg, zap.Any("data", args))

}

func (lw LogWrap) Error(ctx context.Context, msg string, args ...any) {

	lw.logger.Error(msg, zap.Any("data", args))

}

func (lw LogWrap) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	cfg := lw.config
	sql, rows := fc()
	elapsed := time.Since(begin)

	sqlZ := zap.String("sql", sql)
	rowsZ := zap.Int64("rows", rows)
	elapsedZ := zap.Duration("elapsed", elapsed)

	switch {
	case err != nil && (!errors.Is(err, gorm.ErrRecordNotFound) || !cfg.IgnoreRecordNotFoundError):
		lw.logger.Error("trace error", zap.Error(err), elapsedZ, sqlZ, rowsZ)
	case elapsed > cfg.SlowThreshold && cfg.SlowThreshold != 0:
		slow := fmt.Sprintf("trace slow sql >= %v", cfg.SlowThreshold)
		lw.logger.Warn(slow, elapsedZ, sqlZ, rowsZ)
	default:
		lw.logger.Info("trace query", elapsedZ, sqlZ, rowsZ)
	}

}
