package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("visit: https://github.com/Nine26/adcalmath for more updates")
	fmt.Printf("\n1)volume of prism\n\n2)circles\n\n3)solving rational numbers with coefficients\n\n4)perimeter of various shapes\n\n5)area of various shapes\n\n00)exit\n")
	fmt.Print("\npick a number: ")

	scanner.Scan()

	input00, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if input00 == 00 {
		os.Exit(0)
	}

	if input00 == 5 {

		fmt.Printf("\n1)area of triangle\n\n1.1)work out the height of the triangle\n\n2)area of irregular shapes\n") // unfinished code
		fmt.Print("\n\nplease pick a number: ")
		scanner.Scan()

		input, _ := strconv.ParseFloat(scanner.Text(), 64)

		if input == 1 {

			fmt.Printf("\nbase?: ")
			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Printf("\nheight?: ")
			scanner.Scan()

			input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var calcs = float64(0.5) * float64(input1) * float64(input2)

			fmt.Println("\nanswer: ", calcs)

		}

		if input == 1.1 {

			fmt.Printf("\narea?: ")
			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Printf("\nbase or height?: ")

			scanner.Scan()

			input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var calcs = float64(input1) * float64(input2)

			fmt.Println("\nanswer: ", calcs/2)

		}

	}

	if input00 == 4 {

		fmt.Printf("\n1)permiter of rectangle\n\n2)permiter irregular shapes\n\n3)permiter squares")
		fmt.Print("\n\nplease pick a number: ")
		scanner.Scan()

		input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		if input == 1 {

			fmt.Printf("\nwidth?: ")
			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Printf("\nlength?: ")
			scanner.Scan()

			input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var calcs = float64(input1)*2 + float64(input2)*2

			fmt.Println("\nanswer: ", calcs)

		}

		if input == 2 {

			fmt.Printf("\nwidth?: ")
			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Printf("\nlength?: ")
			scanner.Scan()

			input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var calcs = float64(input1) + float64(input2)

			fmt.Println("\nanswer: ", calcs*2)

		}

		if input == 3 {

			fmt.Printf("\ncm of side?:  ")
			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var calcs = float64(input1) * 4

			fmt.Println("\nanswer: ", calcs)

		}

	}

	if input00 == 3 {

		fmt.Printf("\n1)one fraction\n\n2)two fractions") //unfinished code so there might be bugs and wrong answers
		fmt.Print("\n\nplease pick a number: ")

		scanner.Scan()

		input0, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		if input0 == 1 {

			fmt.Printf("\nwhat is your numerator?: ")
			scanner.Scan()

			input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Printf("\nwhat is your denominator?: ")
			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Printf("\nrational number?: ")
			scanner.Scan()

			input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer = float64(input1) * float64(input2)

			fmt.Print("\nanswer: ", float64(answer)/float64(input))

		}
		if input0 == 2 {

			fmt.Print("\n\nwhat is your whole number: ")
			scanner.Scan()

			input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Printf("\n\nwhat is your numerator?: ")
			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Printf("\n\nwhat is your denominator?: ")
			scanner.Scan()

			input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var calcs = float64(input)*float64(input2) + float64(input1)

			fmt.Print("\n\nwhat is your other whole number: ")
			scanner.Scan()

			input3, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Printf("\n\nwhat is your numerator?: ")
			scanner.Scan()

			input4, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Printf("\n\nwhat is your denominator?: ")
			scanner.Scan()

			input5, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var calcs1 = float64(input3)*float64(input5) + float64(input4)
			var improper_fraction1 = float64(calcs1) / float64(input4)

			var calculations = (improper_fraction1 * float64(input2))
			var calculations1 = (calcs * float64(input5))

			fmt.Print("\n\nanswer: ", float64(calculations)/float64(calculations1))

		}

	}

	if input00 == 1 {

		fmt.Printf("\n1)prism\n\n2)triangle\n\n3)paralleogram\n\n4)trapezium\n")
		fmt.Print("please pick a number: ")

		scanner.Scan()

		input0, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		if input0 == 1 {

			fmt.Println("\narea of cross section? : ")

			scanner.Scan()

			input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Println("\nLength? :  ")

			scanner.Scan()

			input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Print("\nanswer = ", input*input2)
			fmt.Println("\npress any key to close")
			scanner.Scan()
		}

		if input0 == 3 {

			fmt.Println("\narea of cross section? : ")

			scanner.Scan()

			input3, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Println("left Length? :  ")

			scanner.Scan()

			input4, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			answer := input3 * input4

			fmt.Println("right side Length? :  ")

			scanner.Scan()

			input5, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Println("\nanswer = ", answer*input5)
			fmt.Println("\npress any key to close")
			scanner.Scan()

		}

		if input0 == 2 {

			fmt.Println("area of cross section? : ")

			scanner.Scan()

			input6, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Println("bottom length? :  ")

			scanner.Scan()

			input7, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			answer := input6 * input7

			fmt.Println("top length Length? :  ")

			scanner.Scan()

			input8, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			answer1 := answer / 2

			fmt.Println("\nanswer = ", answer1*input8)
			fmt.Println("\npress any key to close")
			scanner.Scan()

		}

		if input0 == 4 {

			fmt.Print("parrelal side 1? : ")

			scanner.Scan()

			input6, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Print("parrelal side 2? :  ")

			scanner.Scan()

			input7, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer = float64(input6 + input7)

			fmt.Print("height? :  ")

			scanner.Scan()

			input8, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			answer1 := (answer / 2)

			var answer3 = (answer1 * float64(input8))

			fmt.Print("side length ? :  ")

			scanner.Scan()

			input9, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Print("\nanswer = ", answer3*float64(input9))
			fmt.Println("\npress any key to close")
			scanner.Scan()

		}

	}

	if input00 == 2 {

		fmt.Printf("\n1)work out the circumfrence w radius\n\n2)work out the circumference w diameter\n\n3)work out radius with the circumference\n\n4)work out perimeter for different circles sizes\n\n5)work out perimeter for semicircle\n\n6)work out area of circle w radius\n\n7)work out area of circle w diameter\n\n8)work out radius of circle with area\n\n9)area of cylinder with radius\n\n10)area of cylinder with radius")
		fmt.Print("\n\nplease pick a number: ")

		scanner.Scan()

		input0, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		if input0 == 1 {

			fmt.Print("Radius? : ")

			scanner.Scan()

			input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer1 float64 = (float64(input))

			answer_notrounded := (answer1 * float64(2*math.Pi))

			fmt.Print("round answer to 1dp?\n1 = yes\n2 = no ")
			fmt.Print("\n\npick a number: ")

			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			if input1 == 1 {

				fmt.Println("\nanswer: ", math.Round(answer_notrounded*10)/10)
				fmt.Println("\n\npress any key to close")
				scanner.Scan()

			}

			if input1 == 2 {

				fmt.Println("\nanswer: ", answer_notrounded)
				fmt.Println("\n\npress any key to close")
				scanner.Scan()

			}

		}

		if input0 == 2 {

			fmt.Print("diameter : ")

			scanner.Scan()

			input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer1 float64 = (float64(input) / 2)

			answer_notrounded := (answer1 * float64(2*math.Pi))

			fmt.Print("\nround answer to 1dp?\n1 = yes\n2 = no ")
			fmt.Print("\n\npick a number: ")

			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			if input1 == 1 {

				fmt.Println("\nanswer: ", math.Round(answer_notrounded*10)/10)
				fmt.Println("\n\npress any key to close")
				scanner.Scan()

			}

			if input1 == 2 {

				fmt.Println("\nanswer: ", answer_notrounded)
				fmt.Println("\n\npress any key to close")
				scanner.Scan()

			}

		}

		if input0 == 3 {

			fmt.Print("circumference : ")

			scanner.Scan()

			input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer_notrounded = (float64(input) / (2 * math.Pi))

			fmt.Print("\nround answer to 1dp?\n1 = yes\n2 = no ")
			fmt.Print("\n\npick a number: ")

			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			if input1 == 1 {

				fmt.Println("\nanswer: ", math.Round(answer_notrounded*10)/10)
				fmt.Println("\npress any key to close")
				scanner.Scan()

			}

			if input1 == 2 {

				fmt.Println("\nanswer: ", answer_notrounded)
				fmt.Println("\n\npress any key to close")
				scanner.Scan()

			}

		}

		if input0 == 4 {

			fmt.Print("numerator? : ")

			scanner.Scan()

			input1, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Print("\ndenominator? :  ")

			scanner.Scan()

			input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer = float64(input1) / float64(input2)

			fmt.Print("\nradius? :  ")

			scanner.Scan()

			input3, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer_notrounded = (float64(answer)*(2*math.Pi)*float64(input3) + float64(2*input3))

			fmt.Print("\nround answer to 1dp?\n1 = yes\n2 = no ")
			fmt.Print("\npick a number: ")

			scanner.Scan()

			input4, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			if input4 == 1 {

				fmt.Println("\nanswer: ", math.Round(answer_notrounded*10)/10)
				fmt.Println("\npress any key to close")
				scanner.Scan()

			}

			if input4 == 2 {

				fmt.Println("\nanswer: ", answer_notrounded)
				fmt.Println("\n\npress any key to close")
				scanner.Scan()

			}

		}
		if input0 == 5 {

			fmt.Print("radius? :  ")

			scanner.Scan()

			input3, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer = float64(input3) / float64(2)

			var answer_notrounded = (float64(0.5)*(2*math.Pi)*float64(answer) + float64(2*answer))

			fmt.Print("\nround answer to 1dp?\n1 = yes\n2 = no ")
			fmt.Print("\n\npick a number: ")

			scanner.Scan()

			input4, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			if input4 == 1 {

				fmt.Println("answer: ", math.Round(answer_notrounded*10)/10)
				fmt.Println("press any key to close")
				scanner.Scan()

			}

			if input4 == 2 {

				fmt.Println("answer: ", answer_notrounded)
				fmt.Println("press any key to close")
				scanner.Scan()

			}

		}
		if input0 == 6 {

			fmt.Print("radius?: ")

			scanner.Scan()

			input3, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer_notrounded = (math.Pi * float64(input3*input3))

			fmt.Print("\nround answer to 1dp?\n1 = yes\n2 = no ")
			fmt.Print("\n\npick a number: ")

			scanner.Scan()

			input4, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			if input4 == 1 {

				fmt.Println("\nanswer: ", math.Round(answer_notrounded*10)/10)
				fmt.Println("\npress any key to close")
				scanner.Scan()

			}

			if input4 == 2 {

				fmt.Println("\nanswer: ", answer_notrounded)
				fmt.Println("\n\npress any key to close")
				scanner.Scan()

			}

		}
		if input0 == 7 {

			fmt.Print("diamter?: ")

			scanner.Scan()

			input3, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer = float64(input3) / 2

			var answer_notrounded = (math.Pi * float64(answer*answer))

			fmt.Print("\nround answer to 1dp?\n1 = yes\n2 = no ")
			fmt.Print("\n\npick a number: ")

			scanner.Scan()

			input4, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			if input4 == 1 {

				fmt.Println("answer: ", math.Round(answer_notrounded*10)/10)
				fmt.Println("press any key to close")
				scanner.Scan()

			}

			if input4 == 2 {

				fmt.Println("\nanswer: ", answer_notrounded)
				fmt.Println("\npress any key to close")
				scanner.Scan()

			}

		}
		if input0 == 8 {

			fmt.Print("area?: ")

			scanner.Scan()

			input3, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer = float64(input3) / math.Pi

			var answer_notrounded = (math.Sqrt(answer))

			fmt.Print("round answer to 1dp?\n1 = yes\n2 = no ")
			fmt.Print("\npick a number: ")

			scanner.Scan()

			input4, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			if input4 == 1 {

				fmt.Println("\nanswer: ", math.Round(answer_notrounded*10)/10)
				fmt.Println("\npress any key to close")
				scanner.Scan()

			}

			if input4 == 2 {

				fmt.Println("\nanswer: ", answer_notrounded)

				fmt.Println("\npress any key to close")
				scanner.Scan()

			}

		}

		if input0 == 9 {

			fmt.Print("radius?: ")

			scanner.Scan()

			input3, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Print("height?: ")

			scanner.Scan()

			input4, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer_notrounded = math.Pi * float64(input3) * float64(input3) * float64(input4)

			fmt.Print("round answer to 1dp?\n1 = yes\n2 = no ")
			fmt.Print("\npick a number: ")

			scanner.Scan()

			input5, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			if input5 == 1 {

				fmt.Println("\nanswer: ", math.Round(answer_notrounded*10)/10)
				fmt.Println("\npress any key to close")
				scanner.Scan()

			}

			if input5 == 2 {

				fmt.Println("\nanswer: ", answer_notrounded)

				fmt.Println("\npress any key to close")
				scanner.Scan()

			}

		}
		if input0 == 10 {

			fmt.Print("diameter?: ")

			scanner.Scan()

			input3, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			fmt.Print("height?: ")

			scanner.Scan()

			input4, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			var answer = float64(input3) / 2
			var answer_notrounded = math.Pi * float64(answer) * float64(answer) * float64(input4)

			fmt.Print("round answer to 1dp?\n1 = yes\n2 = no ")
			fmt.Print("\npick a number: ")

			scanner.Scan()

			input5, _ := strconv.ParseInt(scanner.Text(), 10, 64)

			if input5 == 1 {

				fmt.Println("\nanswer: ", math.Round(answer_notrounded*10)/10)
				fmt.Println("\npress any key to close")
				scanner.Scan()

			}

			if input5 == 2 {

				fmt.Println("\nanswer: ", answer_notrounded)

				fmt.Println("\npress any key to close")
				scanner.Scan()

			}

		}
	}
}
