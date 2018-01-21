package bench

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func BenchmarkZapTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	w := zapcore.AddSync(stream)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkZapTextNegative(b *testing.B) {
	stream := &blackholeStream{}
	w := zapcore.AddSync(stream)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.ErrorLevel,
	)
	logger := zap.New(core)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkZapJSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	w := zapcore.AddSync(stream)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog",
				zap.String("rate", "15"),
				zap.Int("low", 16),
				zap.Float32("high", 123.2),
			)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkZapJSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	w := zapcore.AddSync(stream)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.ErrorLevel,
	)
	logger := zap.New(core)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog",
				zap.String("rate", "15"),
				zap.Int("low", 16),
				zap.Float32("high", 123.2),
			)
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}
