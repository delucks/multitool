package main

func ExampleHexConvertBase() {
	_ = ConvertBase([]string{"basejump", "0xfeedface", "10"}, nil)
	_ = ConvertBase([]string{"basejump", "0xdeadbeef", "2"}, nil)
	_ = ConvertBase([]string{"basejump", "0xfeed", "8"}, nil)
	_ = ConvertBase([]string{"basejump", "0xfeed", "16"}, nil)
	// Output:
	// 4277009102
	// 11011110101011011011111011101111
	// 177355
	// feed
}

func ExampleOctalConvertBase() {
	_ = ConvertBase([]string{"basejump", "0o7431", "10"}, nil)
	_ = ConvertBase([]string{"basejump", "0o7431", "16"}, nil)
	_ = ConvertBase([]string{"basejump", "0o54307", "2"}, nil)
	_ = ConvertBase([]string{"basejump", "0o54307", "8"}, nil)
	// Output:
	// 3865
	// f19
	// 101100011000111
	// 54307
}

func ExampleBinaryConvertBase() {
	_ = ConvertBase([]string{"basejump", "0b11010001100001", "10"}, nil)
	_ = ConvertBase([]string{"basejump", "0b11010001100001", "16"}, nil)
	_ = ConvertBase([]string{"basejump", "0b1101", "2"}, nil)
	_ = ConvertBase([]string{"basejump", "0b1101", "8"}, nil)
	// Output:
	// 13409
	// 3461
	// 1101
	// 15
}
