package takeaction_test

import (
	mastergo "github.com/bm4cs/mastergo/takeaction"
	"testing"
)

func Test_LongestString(t *testing.T) {
	t.Run("GivenSimpleStringInputs_WhenRunNormally_ThenLongestStringIdentified", func(t *testing.T) {
		want := 10
		if got := mastergo.LongestString("Six", "sleek", "swans", "swam", "swiftly", "southwards"); got != want {
			t.Errorf("LongestString() = %v, want %v", got, want)
		}
	})
}

func Test_Closures(t *testing.T) {
	t.Run("GivenClosureReturnedFunc_WhenCalledThreeTimes_TheVariableInOuterScopeIsIncremented", func(t *testing.T) {
		want := 3
		if got := mastergo.SingleClosure(); got != want {
			t.Errorf("Closures() = %v, want %v", got, want)
		}
	})
}

func Test_TwinClosures(t *testing.T) {
	t.Run("GivenAFunctionThatReturnsTwoClosures_WhenBothClosuresAreCalled_ThenTheOuterClosureStateIsShared", func(t *testing.T) {
		want := 35
		if got := mastergo.DoubleClosures(); got != want {
			t.Errorf("Closures() = %v, want %v", got, want)
		}
	})
}

func Test_CleverTracingWithDefer(t *testing.T) {
	t.Run("GivenAFunctionThatTracesUsingOneLineDeferAndAClosure_WhenCalled_ThenTheFunctionIsTracedBeginToEnd", func(t *testing.T) {
		mastergo.CleverTracingWithDefer()
	})
}