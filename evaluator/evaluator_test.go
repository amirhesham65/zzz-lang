package evaluator

import (
	"testing"

	"github.com/amirhesham65/zzz-lang/lexer"
	"github.com/amirhesham65/zzz-lang/object"
	"github.com/amirhesham65/zzz-lang/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"yea", true},
		{"nah", false},
		{"!yea", false},
		{"!nah", true},
		{"!5", false},
		{"!!yea", true},
		{"!!nah", false},
		{"!!5", true},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"yea == yea", true},
		{"nah == nah", true},
		{"yea == nah", false},
		{"yea != nah", true},
		{"nah != yea", true},
		{"(1 < 2) == yea", true},
		{"(1 < 2) == nah", false},
		{"(1 > 2) == yea", false},
		{"(1 < 2) == nah", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{"fr (yea) { 10 }", 10},
		{"fr (nah) { 10 }", nil},
		{"fr (1) { 10 }", 10},
		{"fr (1 < 2) { 10 }", 10},
		{"fr (1 > 2) { 10 }", nil},
		{"fr (1 > 2) { 10 } lowkey { 20 }", 20},
		{"fr (1 < 2) { 10 } lowkey { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{
			`
			fr(10 > 1) {
				fr (10 > 1) {
					return 10;
				}
				return 1;
			}
			`, 10,
		},
		{
			`
			fr (10 > 1) {
			fr (10 > 1) {
				return 10;
			}

			return 1;
			}
			`, 10,
		},
		{
			`
			lit f = fun(x) {
			return x;
			x + 10;
			};
			f(10);`, 10,
		},
		{
			`
			lit f = fun(x) {
			lit result = x + 10;
			return result;
			return 10;
			};
			f(10);`, 20,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
   lit newAdder = fun(x) {
     fun(y) { x + y };
};
   lit addTwo = newAdder(2);
   addTwo(2);`
	testIntegerObject(t, testEval(input), 4)
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"lit a = 5; a;", 5},
		{"lit a = 5 * 5; a;", 25},
		{"lit a = 5; lit b = a; b;", 5},
		{"lit a = 5; lit b = a; lit c = a + b + 5; c;", 15},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + yea;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + yea; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-yea",
			"unknown operator: -BOOLEAN",
		},
		{
			"yea + nah;",
			"unknown operator: BOOLEAN + BOOLEAN",
		}, {
			"5; yea + nah; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"fr (10 > 1) { yea + nah; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
			fr (10 > 1) {
				fr (10 > 1) {
				  return yea + nah;
				}
				return 1; 
			}
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"undefined identifier: foobar",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
		{`spit("hello", "world")`, nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T(%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		}
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, expected=%d", result.Value, expected)
		return false
	}
	return true
}

func TestFunctionObject(t *testing.T) {
	input := "fun(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"lit identity = fun(x) { x; }; identity(5);", 5},
		{"lit identity = fun(x) { return x; }; identity(5);", 5},
		{"lit double = fun(x) { x * 2; }; double(5);", 10},
		{"lit add = fun(x, y) { x + y; }; add(5, 5);", 10},
		{"lit add = fun(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fun(x) { x; }(5)", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!";`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("str.Value is not \"Hello World!\", got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!";`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("str.Value is not \"Hello World!\", got=%q", str.Value)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, expected=%t", result.Value, expected)
		return false
	}
	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}
