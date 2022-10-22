// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package log

import (
	"context"
	"os"
	"runtime"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config log 配置
type Config struct {
	Filename             string
	RollingLogMaxSize    int
	RollingLogMaxBackups int
	RollingLogMaxAge     int
	Compress             bool
	LastLog              func(ctx context.Context) []zap.Field
	Level                zapcore.Level
	StdOut               bool
}

// Options 配置选项
type Options func(config *Config)

// WithOutput 输出文件
func WithOutput(filename string) Options {
	return func(config *Config) {
		config.Filename = filename
	}
}

// WithStdout 标准输出
func WithStdout() Options {
	return func(config *Config) {
		config.StdOut = true
	}
}

// NewConfig 实例化日志配置
func NewConfig(options ...Options) *Config {
	for _, opt := range options {
		opt(defaultConfig)
	}
	return defaultConfig
}

var (
	// Logger 日志实例
	Logger        *zap.Logger
	defaultConfig = &Config{
		Filename:             getExecuteName(),
		RollingLogMaxSize:    1024 * 2,
		RollingLogMaxBackups: 10,
		RollingLogMaxAge:     7,
		Compress:             false,
		LastLog:              nil,
		Level:                zap.DebugLevel,
		StdOut:               false,
	}
)

func init() {
	InitLog(defaultConfig)
}

func InitLog(config *Config) {

	hook := lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.RollingLogMaxSize, // megabytes
		MaxBackups: config.RollingLogMaxBackups,
		MaxAge:     config.RollingLogMaxAge, // days
		Compress:   false,                   // disabled by default
	}

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= config.Level
	})

	var writeSyncer = []zapcore.WriteSyncer{
		zapcore.AddSync(&hook),
	}

	if defaultConfig.StdOut {
		// 同时打印到标准输出
		writeSyncer = append(writeSyncer, zapcore.AddSync(os.Stdout))
	}

	multiWriter := zapcore.NewMultiWriteSyncer(writeSyncer...)
	Logger = zap.New(
		zapcore.NewTee(zapcore.NewCore(consoleEncoder, multiWriter, highPriority)),
	).
		WithOptions(zap.AddCaller()).
		WithOptions(zap.AddCallerSkip(1))
}

// getExecuteName 获取当前运行时的文件名
func getExecuteName() string {
	var path = os.Args[0]
	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "\\", "/")
		if len(path) > 4 && path[len(path)-4:] == ".exe" {
			path = path[:len(path)-4]
		}
	}
	paths := strings.Split(path, "/")
	path = paths[len(paths)-1]
	return path + ".log"
}

func lastLog(ctx context.Context, f func(ctx context.Context) []zap.Field) []zap.Field {
	if f != nil {
		return f(ctx)
	}
	return nil
}

// Debug Debug日志
func Debug(ctx context.Context, msg string) {
	Logger.Debug(msg, lastLog(ctx, defaultConfig.LastLog)...)
}

// Info 信息日志
func Info(ctx context.Context, msg string) {
	Logger.Info(msg, lastLog(ctx, defaultConfig.LastLog)...)
}

// Warn 警告日志
func Warn(ctx context.Context, msg string) {
	Logger.Warn(msg, lastLog(ctx, defaultConfig.LastLog)...)
}

// Error 错误日志
func Error(ctx context.Context, msg string) {
	Logger.Error(msg, lastLog(ctx, defaultConfig.LastLog)...)
}
