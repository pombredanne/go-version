package misc

import (
	"testing"
)

var anyConstraints = map[string][]*Constraint{
	"*":       []*Constraint{},
	"*.*":     []*Constraint{},
	"*.x.*":   []*Constraint{},
	"x.x.x.*": []*Constraint{},
}

func TestParseAnyConstraints(t *testing.T) {
	for in, out := range anyConstraints {
		if x := ParseConstraints(in); len(x) != 0 {
			t.Errorf("FAIL: parseConstraints(%v) = {%s}: want {%s}", in, x, out)
		}
	}
}

var simpleConstraints = map[string][]*Constraint{
	"<>1.0.0":             []*Constraint{{"<>", "1.0.0.0"}},
	"!=1.0.0":             []*Constraint{{"!=", "1.0.0.0"}},
	">1.0.0":              []*Constraint{{">", "1.0.0.0"}},
	"<1.2.3.4":            []*Constraint{{"<", "1.2.3.4-dev"}},
	"<=1.2.3":             []*Constraint{{"<=", "1.2.3.0"}},
	">=1.2.3":             []*Constraint{{">=", "1.2.3.0"}},
	"=1.2.3":              []*Constraint{{"=", "1.2.3.0"}},
	"==1.2.3":             []*Constraint{{"=", "1.2.3.0"}},
	"1.2.3":               []*Constraint{{"=", "1.2.3.0"}},
	"=1.0":                []*Constraint{{"=", "1.0.0.0"}},
	"1.2.3b5":             []*Constraint{{"=", "1.2.3.0-beta5"}},
	">= 1.2.3":            []*Constraint{{">=", "1.2.3.0"}},
	">=dev-master":        []*Constraint{{">=", "9999999-dev"}},
	"dev-master":          []*Constraint{{"=", "9999999-dev"}},
	"dev-feature-a":       []*Constraint{{"=", "dev-feature-a"}},
	"dev-some-fix":        []*Constraint{{"=", "dev-some-fix"}},
	"dev-CAPS":            []*Constraint{{"=", "dev-CAPS"}},
	"dev-master as 1.0.0": []*Constraint{{"=", "9999999-dev"}},
	"<1.2.3.4-stable":     []*Constraint{{"<", "1.2.3.4"}},

	"<=3.0@dev":            []*Constraint{{"<=", "3.0.0.0"}},
	"1.0@dev":              []*Constraint{{"=", "1.0.0.0"}},                 //IgnoresStabilityFlag
	"1.0.x-dev#abcd123":    []*Constraint{{"=", "1.0.9999999.9999999-dev"}}, //IgnoresReferenceOnDevVersion
	"1.0.x-dev#trunk/@123": []*Constraint{{"=", "1.0.9999999.9999999-dev"}}, //IgnoresReferenceOnDevVersion
	//"1.0#abcd123":          []string{"=", "1.0.0.0"},                 //FailsOnBadReference
	//"1.0#trunk/@123":       []string{"=", "1.0.0.0"},                 //FailsOnBadReference
}

func TestParseConstraints(t *testing.T) {
	for in, out := range simpleConstraints {
		if x := ParseConstraints(in); x[0].String() != out[0].String() {
			t.Errorf("FAIL: parseConstraints(%v) = {%s}: want {%s}", in, x, out)
		}
	}
}

var wildcardConstraints = map[string][]*Constraint{
	"2.*":     []*Constraint{{">", "1.9999999.9999999.9999999"}, {"<", "2.9999999.9999999.9999999"}},
	"20.*":    []*Constraint{{">", "19.9999999.9999999.9999999"}, {"<", "20.9999999.9999999.9999999"}},
	"2.0.*":   []*Constraint{{">", "1.9999999.9999999.9999999"}, {"<", "2.0.9999999.9999999"}},
	"2.2.x":   []*Constraint{{">", "2.1.9999999.9999999"}, {"<", "2.2.9999999.9999999"}},
	"2.10.x":  []*Constraint{{">", "2.9.9999999.9999999"}, {"<", "2.10.9999999.9999999"}},
	"2.1.3.*": []*Constraint{{">", "2.1.2.9999999"}, {"<", "2.1.3.9999999"}},
	"0.*":     []*Constraint{{"<", "0.9999999.9999999.9999999"}},
}

func TestParseConstraintsWildcardConstraints(t *testing.T) {
	for in, out := range wildcardConstraints {
		if x := ParseConstraints(in); x[0].String() != out[0].String() && x[1].String() != out[1].String() {
			t.Errorf("FAIL: parseConstraints(%v) = {%s}: want {%s}", in, x, out)
		}
	}
}

var tildeConstraints = map[string][]*Constraint{
	"~1":         []*Constraint{{">=", "1.0.0.0"}, {"<", "2.0.0.0-dev"}},
	"~1.2":       []*Constraint{{">=", "1.2.0.0"}, {"<", "2.0.0.0-dev"}},
	"~1.2.3":     []*Constraint{{">=", "1.2.3.0"}, {"<", "1.3.0.0-dev"}},
	"~1.2.3.4":   []*Constraint{{">=", "1.2.3.4"}, {"<", "1.2.4.0-dev"}},
	"~1.2-beta":  []*Constraint{{">=", "1.2.0.0-beta"}, {"<", "2.0.0.0-dev"}},
	"~1.2-b2":    []*Constraint{{">=", "1.2.0.0-beta2"}, {"<", "2.0.0.0-dev"}},
	"~1.2-BETA2": []*Constraint{{">=", "1.2.0.0-beta2"}, {"<", "2.0.0.0-dev"}},
	"~1.2.2-dev": []*Constraint{{">=", "1.2.2.0-dev"}, {"<", "1.3.0.0-dev"}},
}

func TestParseConstraintsTildeConstraints(t *testing.T) {
	for in, out := range tildeConstraints {
		if x := ParseConstraints(in); x[0].String() != out[0].String() && x[1].String() != out[1].String() {
			t.Errorf("FAIL: parseConstraints(%v) = {%s}: want {%s}", in, x, out)
		}
	}
}

var multiConstraints = map[string][]*Constraint{
	">2.0,<=3.0":            []*Constraint{{">", "2.0.0.0"}, {"<=", "3.0.0.0"}},
	">2.0@stable,<=3.0@dev": []*Constraint{{">", "2.0.0.0"}, {"<=", "3.0.0.0-dev"}},
}

func TestParseConstraintsMultiConstraints(t *testing.T) {
	for in, out := range multiConstraints {
		if x := ParseConstraints(in); x[0].String() != out[0].String() && x[1].String() != out[1].String() {
			t.Errorf("FAIL: parseConstraints(%v) = {%s}: want {%s}", in, x, out)
		}
	}
}
