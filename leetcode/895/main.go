package main

import "fmt"

// Solution rationale:
//
// There are essentially two properties to keep track of:
//   1. The frequency of each element in the stack;
//   2. The number representing the highest frequency in the stack.
//
// The first property, though, derive two other controls that allow keeping track
// of the ordering and the frequency of each element in the stack:
//   1. A map that keeps track of the frequencies of elements in the stack; the key
//   represent a single frequency (like 1, 2, 3, etc.) and the value is a stack of elements
//   that appear that many times in the stack. As new values are added to the stack, this map
//   is updated accordingly to the result of the next control;
//   2. A map that keeps track of the individual frequency of an element; the key is the
//   element and the value is the number of times it appears in the stack. As new values are
//   pushed or popped, the new value is updated and it serves as an input that drives where
//   this element should be positioned in the frequency map described above.
//
// To better illustrate this, let's consider the following example:
//
//   Given the stack: [3, 3, 2, 1, 2]
//   The map of frequencies will look like: { 1: [3, 2, 1]; 2: [3, 2] }
//   The map of individual frequencies will look like: { 3: 2; 2: 2; 1: 1 }
//
// With this set of controls, we can easily identify the most frequent element in the stack
// while keeping track of the ordering of the elements in the stack, also solving the tie.
//
// The push operation increments the individual frequency of the element being pushed and
// updates the frequency map accordingly to the new count found for the element.
// It also checks if the new found frequency is greater than the current max frequency; if it is,
// then the max frequency is updated. The element is then added to the end of the sub-stack in the
// frequency map, so that's easy to identify the element that was pushed last, solving the tie problem.
//
// In the pop operation, once it reads the current max frequency, it pops the last element from the
// sub-stack in the frequency map with that frequency. If the sub-stack becomes empty, then the max
// frequency is decremented and the frequency key is removed from the frequency map.
// The element obtained is then updated in the individual frequency map, decrementing its count.
//
// The result of the pop operation is the element that was popped from the frequency map.

func main() {
	// A conventional stack approach will not be efficient because I have to
	// keep track of the frequency of each element, identify which one is closest to the
	// top of the stack, and perform an unconventional pop operation.

	// At each push operation, the stack might be recalculated depending on how
	// the push operation changes the frequency of the element:
	// - If the element becomes the most frequent, then it will be at the top of
	// the stack for the pop operation;
	// - If there is a tie, then it should be positioned right after the closest
	// element that is already at the top of the stack;

	// All of this happens at the same time as we keep track of the right positions in the stack
	// so that the push operation happens correctly.

	fs := NewFreqStack()
	fs.Push(5)
	fs.Push(7)
	fs.Push(5)
	fs.Push(7)
	fs.Push(4)
	fs.Push(5)
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())
}

// FreqStack wraps the values and the behavior of a stack that keeps track of the
// frequency of each element. When popping from a FreqStack, it pops the most frequent
// element in the stack. If there is a tie, then the element closer to the top of the
// stack is popped, meaning that if all the elements appear the same number of times,
// then it behaves as a regular stack.
type FreqStack struct {
	// maxFreq stores the current maximum frequency count in the stack.
	maxFreq int

	// freq keeps track of the frequency of each element; they key is the frequency
	// and the value is a stack of elements with that frequency in the order they were pushed.
	// Example:
	//    Given the stack: [3, 3, 2, 1, 2]
	//    freq will look like: { 1: [3, 2, 1]; 2: [3, 2] }
	//
	// 3, 2, and 1 appear once; 3 and 2 appear twice.
	freq map[int][]int

	// valCount keeps track of the frequency of each element in the stack.
	valCount map[int]int
}

// NewFreqStack creates and initializes a new FreqStack.
func NewFreqStack() *FreqStack {
	return &FreqStack{
		maxFreq:  0,
		freq:     make(map[int][]int),
		valCount: make(map[int]int),
	}
}

// Push adds an element to the stack.
func (fs *FreqStack) Push(v int) {
	count := fs.valCount[v] + 1
	fs.valCount[v] = count

	freq := fs.freq[count]
	if freq == nil {
		freq = make([]int, 0)
	}
	fs.freq[count] = append(freq, v)

	if count > fs.maxFreq {
		fs.maxFreq = count
	}
}

// Pop removes the element with the highest frequency from the stack. If there is a tie,
// then the element closer to the top of the stack in the tie is removed.
func (fs *FreqStack) Pop() int {
	// Read the most recent and most frequent value added to the stack. That's
	// the result of the pop operation.
	poppedVal := fs.freq[fs.maxFreq][len(fs.freq[fs.maxFreq])-1]

	// Remove the value found above from the frequency stack.
	fs.freq[fs.maxFreq] = fs.freq[fs.maxFreq][:len(fs.freq[fs.maxFreq])-1]
	// If the frequency stack becomes empty, it means that the last value with this max
	// frequency was removed. Thus maxFreq has to be decremented and the frequency key is cleaned up
	// from the frequency map.
	if len(fs.freq[fs.maxFreq]) == 0 {
		delete(fs.freq, fs.maxFreq)
		fs.maxFreq = fs.maxFreq - 1
	}

	// Decrement valCount for the popped value. If the count reaches zero, then
	// it is removed from the map.
	fs.valCount[poppedVal] = fs.valCount[poppedVal] - 1
	if fs.valCount[poppedVal] == 0 {
		delete(fs.valCount, poppedVal)
	}

	return poppedVal
}
