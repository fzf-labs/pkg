//go:build linux || darwin

package proc

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"sync/atomic"
	"syscall"
	"time"

	"golang.org/x/exp/slog"
)

// DefaultMemProfileRate is the default memory profiling rate.
// See also http://golang.org/pkg/runtime/#pkg-variables
const DefaultMemProfileRate = 4096

// started is non zero if a profile is running.
var started uint32

// Profile represents an active profiling session.
type Profile struct {
	// closers holds cleanup functions that run after each profile
	closers []func()

	// stopped records if a call to profile.Stop has been made
	stopped uint32
}

func (p *Profile) close() {
	for _, closer := range p.closers {
		closer()
	}
}

func (p *Profile) startBlockProfile() {
	fn := createDumpFile("block")
	f, err := os.Create(fn)
	if err != nil {
		slog.Error("profile: could not create block profile %q: %v", fn, err)
		return
	}

	runtime.SetBlockProfileRate(1)
	slog.Info("profile: block profiling enabled, %s", fn)
	p.closers = append(p.closers, func() {
		_ = pprof.Lookup("block").WriteTo(f, 0)
		_ = f.Close()
		runtime.SetBlockProfileRate(0)
		slog.Info("profile: block profiling disabled, %s", fn)
	})
}

func (p *Profile) startCPUProfile() {
	fn := createDumpFile("cpu")
	f, err := os.Create(fn)
	if err != nil {
		slog.Error("profile: could not create cpu profile %q: %v", fn, err)
		return
	}

	slog.Info("profile: cpu profiling enabled, %s", fn)
	_ = pprof.StartCPUProfile(f)
	p.closers = append(p.closers, func() {
		pprof.StopCPUProfile()
		_ = f.Close()
		slog.Info("profile: cpu profiling disabled, %s", fn)
	})
}

func (p *Profile) startMemProfile() {
	fn := createDumpFile("mem")
	f, err := os.Create(fn)
	if err != nil {
		slog.Error("profile: could not create memory profile %q: %v", fn, err)
		return
	}

	old := runtime.MemProfileRate
	runtime.MemProfileRate = DefaultMemProfileRate
	slog.Info("profile: memory profiling enabled (rate %d), %s", runtime.MemProfileRate, fn)
	p.closers = append(p.closers, func() {
		_ = pprof.Lookup("heap").WriteTo(f, 0)
		_ = f.Close()
		runtime.MemProfileRate = old
		slog.Info("profile: memory profiling disabled, %s", fn)
	})
}

func (p *Profile) startMutexProfile() {
	fn := createDumpFile("mutex")
	f, err := os.Create(fn)
	if err != nil {
		slog.Error("profile: could not create mutex profile %q: %v", fn, err)
		return
	}

	runtime.SetMutexProfileFraction(1)
	slog.Info("profile: mutex profiling enabled, %s", fn)
	p.closers = append(p.closers, func() {
		if mp := pprof.Lookup("mutex"); mp != nil {
			_ = mp.WriteTo(f, 0)
		}
		_ = f.Close()
		runtime.SetMutexProfileFraction(0)
		slog.Info("profile: mutex profiling disabled, %s", fn)
	})
}

func (p *Profile) startThreadCreateProfile() {
	fn := createDumpFile("threadcreate")
	f, err := os.Create(fn)
	if err != nil {
		slog.Error("profile: could not create threadcreate profile %q: %v", fn, err)
		return
	}

	slog.Info("profile: threadcreate profiling enabled, %s", fn)
	p.closers = append(p.closers, func() {
		if mp := pprof.Lookup("threadcreate"); mp != nil {
			_ = mp.WriteTo(f, 0)
		}
		_ = f.Close()
		slog.Info("profile: threadcreate profiling disabled, %s", fn)
	})
}

func (p *Profile) startTraceProfile() {
	fn := createDumpFile("trace")
	f, err := os.Create(fn)
	if err != nil {
		slog.Error("profile: could not create trace output file %q: %v", fn, err)
		return
	}

	if err := trace.Start(f); err != nil {
		slog.Error("profile: could not start trace: %v", err)
		return
	}

	slog.Info("profile: trace enabled, %s", fn)
	p.closers = append(p.closers, func() {
		trace.Stop()
		slog.Info("profile: trace disabled, %s", fn)
	})
}

// Stop stops the profile and flushes any unwritten data.
func (p *Profile) Stop() {
	if !atomic.CompareAndSwapUint32(&p.stopped, 0, 1) {
		// someone has already called close
		return
	}
	p.close()
	atomic.StoreUint32(&started, 0)
}

// StartProfile starts a new profiling session.
// The caller should call the Stop method on the value returned
// to cleanly stop profiling.
func StartProfile() Stopper {
	if !atomic.CompareAndSwapUint32(&started, 0, 1) {
		slog.Error("profile: Start() already called")
		return noopStopper
	}

	var prof Profile
	prof.startCPUProfile()
	prof.startMemProfile()
	prof.startMutexProfile()
	prof.startBlockProfile()
	prof.startTraceProfile()
	prof.startThreadCreateProfile()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		<-c

		slog.Info("profile: caught interrupt, stopping profiles")
		prof.Stop()

		signal.Reset()
		_ = syscall.Kill(os.Getpid(), syscall.SIGINT)
	}()

	return &prof
}

func createDumpFile(kind string) string {
	command := path.Base(os.Args[0])
	pid := syscall.Getpid()
	return path.Join(fmt.Sprintf("%s-%d-%s-%s.pprof",
		command, pid, kind, time.Now().Format(timeFormat)))
}
