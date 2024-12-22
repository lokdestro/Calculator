package calc

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	calculator := New()

	tests := []struct {
		name        string
		expression  string
		expected    float64
		shouldError bool
	}{
		{
			name:       "Simple addition",
			expression: "2 + 2",
			expected:   4,
		},
		{
			name:       "Simple subtraction",
			expression: "10 - 4",
			expected:   6,
		},
		{
			name:       "Simple multiplication",
			expression: "3 * 3",
			expected:   9,
		},
		{
			name:       "Simple division",
			expression: "8 / 2",
			expected:   4,
		},
		{
			name:        "Division by zero",
			expression:  "10 / 0",
			shouldError: true,
		},
		{
			name:       "Parentheses precedence",
			expression: "(2 + 3) * 4",
			expected:   20,
		},
		{
			name:       "Operator precedence",
			expression: "2 + 3 * 4",
			expected:   14,
		},
		{
			name:       "Floating point numbers",
			expression: "5.5 + 1.2",
			expected:   6.7, // учтите возможные погрешности вычислений
		},
		{
			name:        "Invalid input",
			expression:  "2 + + 3",
			shouldError: true,
		},
		{
			name:       "Complex expression",
			expression: "10 + (5 * 2) - 3",
			expected:   17,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculator.Calc(test.expression)

			if test.shouldError {
				if err == nil {
					t.Fatalf("expected an error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != test.expected {
				t.Fatalf("expected %v but got %v", test.expected, result)
			}
		})
	}
}
