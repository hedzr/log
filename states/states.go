package states

import (
	"github.com/hedzr/log/buildtags"
	"github.com/hedzr/log/isdelve"
	"github.com/hedzr/log/trace"
)

// CmdrMinimal provides the accessors to debug/trace flags
type CmdrMinimal interface {
	InDebugging() bool
	GetDebugMode() bool
	SetDebugMode(b bool)
	GetDebugLevel() int // return debug level as a integer, 0..n, it represents count of `--debug` or set by caller explicitly
	SetDebugLevel(hits int)

	GetTraceMode() bool
	SetTraceMode(b bool)
	GetTraceLevel() int // return trace level as a integer, 0..n, it represents count of `--trace` or set by caller explicitly
	SetTraceLevel(hits int)

	IsNoColorMode() bool
	SetNoColorMode(b bool)
	CountOfNoColor() int
	SetNoColorCount(hits int)

	IsVerboseMode() bool
	IsVerboseModePure() bool
	SetVerboseMode(b bool)
	CountOfVerbose() int
	SetVerboseCount(hits int)

	IsQuietMode() bool
	SetQuietMode(b bool)
	CountOfQuiet() int
	SetQuietCount(hits int)
}

func Env() CmdrMinimal { return env }

var env = &minimalEnv{}

// minimalEnv structure holds the debug/trace flags and provides CmdrMinimal accessors
type minimalEnv struct {
	debugMode    bool
	debugLevel   int
	traceMode    bool
	traceLevel   int
	noColorMode  bool
	noColorCount int
	verboseMode  bool
	verboseCount int
	quietMode    bool
	quietCount   int
}

// InDebugging check if the delve debugger presents
func (e *minimalEnv) InDebugging() bool { return isdelve.Enabled }

// GetDebugMode return the debug boolean flag generally
func (e *minimalEnv) GetDebugMode() bool { return e.debugMode || isdelve.Enabled }

// SetDebugMode set the debug boolean flag generally
func (e *minimalEnv) SetDebugMode(b bool)    { e.debugMode = b }
func (e *minimalEnv) GetDebugLevel() int     { return e.debugLevel }
func (e *minimalEnv) SetDebugLevel(hits int) { e.debugLevel = hits }

// GetTraceMode return the trace boolean flag generally
func (e *minimalEnv) GetTraceMode() bool { return e.traceMode || trace.IsEnabled() }

// SetTraceMode set the trace boolean flag generally
func (e *minimalEnv) SetTraceMode(b bool)    { e.traceMode = b }
func (e *minimalEnv) GetTraceLevel() int     { return e.traceLevel }
func (e *minimalEnv) SetTraceLevel(hits int) { e.traceLevel = hits }

func (e *minimalEnv) IsNoColorMode() bool      { return e.noColorMode }
func (e *minimalEnv) SetNoColorMode(b bool)    { e.noColorMode = b }
func (e *minimalEnv) CountOfNoColor() int      { return e.noColorCount }
func (e *minimalEnv) SetNoColorCount(hits int) { e.noColorCount = hits }

func (e *minimalEnv) IsVerboseMode() bool      { return buildtags.VerboseEnabled || e.verboseMode }
func (e *minimalEnv) IsVerboseModePure() bool  { return e.verboseMode }
func (e *minimalEnv) SetVerboseMode(b bool)    { e.verboseMode = b }
func (e *minimalEnv) CountOfVerbose() int      { return e.verboseCount }
func (e *minimalEnv) SetVerboseCount(hits int) { e.verboseCount = hits }

func (e *minimalEnv) IsQuietMode() bool      { return e.quietMode }
func (e *minimalEnv) SetQuietMode(b bool)    { e.quietMode = b }
func (e *minimalEnv) CountOfQuiet() int      { return e.quietCount }
func (e *minimalEnv) SetQuietCount(hits int) { e.quietCount = hits }
