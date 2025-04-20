// Variable Scope Helper functions
package interpreter

import "fmt"

// pushScope added a new scope map in scopes stack
func (v *HaruVisitor) pushScope() {
	v.scopes = append(v.scopes, make(map[string]Value))
}

// popScope removes most recent scope map from scopes stack
func (v *HaruVisitor) popScope() {
	if len(v.scopes) > 0 {
		v.scopes = v.scopes[:len(v.scopes)-1]
	}
}

// currentScope return present scope map being evaluated
func (v *HaruVisitor) currentScope() map[string]Value {
	return v.scopes[len(v.scopes)-1]
}

// declare adds new entry in current scope map
func (v *HaruVisitor) declare(name string, val Value) {
	v.currentScope()[name] = val
}

// assign updates the variables existing value with the new one
func (v *HaruVisitor) assign(name string, val Value) {
	for i := len(v.scopes) - 1; i >= 0; i-- {
		if _, ok := v.scopes[i][name]; ok {
			v.scopes[i][name] = val
			return
		}
	}

	runtimeErr(fmt.Sprintf("undefined variable '%s'", name))
}

// resolve finds and returns the Value struct by given variable name
// making sure stack is traversed based on local first,
// if variable is not found in scope stack empty Value struct is returned with bool false
func (v *HaruVisitor) resolve(name string) (Value, bool) {
	for i := len(v.scopes) - 1; i >= 0; i-- {
		if val, ok := v.scopes[i][name]; ok {
			return val, true
		}
	}

	return Value{}, false
}
