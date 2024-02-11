package main

import (
	"fmt"
	"strings"
)

func main() {
	loremIpsum := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas lobortis purus nec turpis facilisis,
	in placerat nisl fermentum. Nullam facilisis purus vel nulla tincidunt egestas et consectetur dui. Sed nec aliquam sapien,
	eget sagittis arcu. Aenean placerat augue vitae ante tincidunt, nec aliquam elit interdum. Ut posuere, felis iaculis bibendum maximus,
	neque ligula volutpat sapien, eget gravida lacus velit et arcu. Aliquam quis eros risus. Mauris eleifend, sem sed rhoncus condimentum,
	neque ligula viverra mauris, a tristique lectus orci id arcu. In hac habitasse platea dictumst. Interdum et malesuada fames ac ante ipsum
	primis in faucibus. In quis varius justo. Aliquam ac massa elit. Donec nec auctor sapien, ac tempor ante.`

	targetWord := "dolor"

	loremIpsum = strings.ToLower(loremIpsum)
	targetWord = strings.ToLower(targetWord)

	firstIndex := strings.Index(loremIpsum, targetWord)
	lastIndex := strings.LastIndex(loremIpsum, targetWord)

	fmt.Printf("The word %q appears at index (%d/%d) first and at index (%d/%d) last in the sentence.", targetWord, firstIndex, len(loremIpsum), lastIndex, len(loremIpsum))
}
