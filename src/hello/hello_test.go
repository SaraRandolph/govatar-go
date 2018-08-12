package main

import "testing"

func TestCalculate(t *testing.T) {
	if division(6) != 4 {
		t.Error("Expected result to equal 4")
		// f, err := os.OpenFile("/tmp/results.json", os.O_APPEND|os.O_WRONLY, 0600)
		// if err != nil {
		// 	panic(err)
		// }

		// defer f.Close()

		// if _, err = f.WriteString("Expected result to equal 4"); err != nil {
		// 	panic(err)
		// }
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{6, 4},
		{8, 5},
		{1, 1},
		{80, 41},
		{3, 2},
	}

	for _, test := range tests {
		if output := division(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if division(6) != 4 {
			b.Error("Expected result to equal 4")
		}
	}
}
