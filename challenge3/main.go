package main

import (
  "fmt"
  "os"
)

// main runs in O(n^2) time and O(n) space where n is the number of Xs. It runs in O(n^2) because for each X it runs
// through every subsequent X and sets it to 1 for printing. It runs in O(n) space because it keeps a slice of the X
// indices.
func main() {
  // Find all the indices that have an X, keep track of them in a slice, and replace them with a 0.
  xIndices := make([]int, 0)
  bytes := []byte(os.Args[1])
  for idx, b := range bytes {
    if b == 'X' {
      xIndices = append(xIndices, idx)
      bytes[idx] = '0'
    }
  }

  // Print the string that has 0s in the place of all Xs.
  fmt.Printf("%s\n", bytes)

  // Run this loop while there are still indices to be replaced with 1s.
  for len(xIndices) > 0 {
    lastIndex := len(xIndices)-1

    // Replace each index of X with a 1, print the string and then set it back to 0.
    for i := lastIndex; i >= 0; i-- {
      bytes[xIndices[i]] = '1'
      fmt.Printf("%s\n", bytes)
      bytes[xIndices[i]] = '0'
    }

    // Set the last X index to a 1 and remove it from the slice of xIndices to be replaced with 1s.
    bytes[xIndices[lastIndex]] = '1'
    xIndices = xIndices[:lastIndex]
  }
}
