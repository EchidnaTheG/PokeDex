package main

import (
	"testing"

)

func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
		{	input:    "  hello  world  ",
			expected: []string{"hello", "world"},
			},{
			input:    "1  hello  world  ",
			expected: []string{"1","hello", "world"},
			},{
			input:    "                       1 2                      4Dada 211",
			expected: []string{"1","2", "4Dada", "211"},
			},
		
		
	 }
	


	for _,c := range cases{
		actual := cleanInput(c.input)
		for i := range actual{
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("Word isn't Equal Expected Word")
			}
		}
	} 

}

/**func TestGetLocation(t *testing.T){
	cases := []struct{
		input string
		output string
	}{

	}
}
**/