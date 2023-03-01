// Copyright © 2020 Hedzr Yeh.

package log

// // CmdrMinimal provides the accessors to debug/trace flags
// type CmdrMinimal interface {
// 	states.CmdrMinimal
// }
//
// // MinimalEnv return the Env/CmdrMinimal object
// func MinimalEnv() CmdrMinimal { return states.Env() }
//
// // InDebugging check if the delve debugger presents
// func InDebugging() bool { return states.Env().InDebugging() }
//
// // GetDebugMode return the debug boolean flag generally
// func GetDebugMode() bool { return states.Env().GetDebugMode() }
//
// // GetTraceMode return the trace boolean flag generally
// func GetTraceMode() bool { return states.Env().GetTraceMode() }
//
// // SetDebugMode set the debug boolean flag generally
// func SetDebugMode(b bool) {
// 	if b && GetLevel() < DebugLevel {
// 		SetLevel(DebugLevel)
// 	}
// 	states.Env().SetDebugMode(b)
// }
//
// // SetTraceMode set the trace boolean flag generally
// func SetTraceMode(b bool) {
// 	if b {
// 		SetLevel(TraceLevel)
// 	}
// 	states.Env().SetTraceMode(b)
// }
