package bad

import (
	"sync"
	"testing"
)

func TestGlobalNumAdd(t *testing.T) {
	ret := GlobalNumAdd(10)
	if ret != 10 {
		t.Errorf("expect 10,  actually:%d\n", ret)
	}
}

var mutex *sync.Mutex

func TestGlobalNumSub(t *testing.T) {
	ret := GlobalNumSub(1)
	if ret != -1 {
		t.Errorf("expect -1,  actually:%d\n", ret)
	}
}
func TestSomeFunc(t *testing.T) {
	mutex.Lock()
}

// Output:
//=== RUN   TestGlobalNumAdd
//--- PASS: TestGlobalNumAdd (0.00s)
//=== RUN   TestGlobalNumSub
//    bad_test.go:15: expect -1,  actually:9
//--- FAIL: TestGlobalNumSub (0.00s)
//FAIL
//coverage: 100.0% of statements
//exit status 1
//FAIL    golang-Testability-code-guide/global_var/bad    0.052s
