package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type item struct {
  name string
  price int
}

func printItems(i, j int, items []item) {
  fmt.Printf("%s %d, %s %d\n", items[i].name, items[i].price, items[j].name, items[j].price)
}

// main will run in O(n) where n is the number of items in the file. Essentially the algorithm runs through the list
// with two indices, so in the worst case both indices run through the entire list in O(2n), giving the runtime of O(n)
// after dropping the constants.
func main() {
  // Open the file passed as the first command line argument.
  filename := os.Args[1]
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  // Scan each line of the file and add each item to the list of items.
  items := make([]item, 0)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := strings.Split(string(scanner.Bytes()), ",")
    price, err := strconv.Atoi(strings.TrimSpace(line[1]))
    if err != nil {
      fmt.Println(err.Error())
      os.Exit(2)
    }
    items = append(items, item{name: line[0], price: price})
  }

  // Check that an error did not occur while scanning the file.
  if err := scanner.Err(); err != nil {
    fmt.Println(err.Error())
    os.Exit(3)
  }

  // Convert the second command line argument to an integer representing the target price.
  targetPrice, err := strconv.Atoi(os.Args[2])
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(4)
  }

  // Initialize the total price and the indices. If the first two prices are greater than the target price it will be
  // impossible to find a lower price combination because the prices are in sorted order.
  i, j := 0, 1
  totalPrice := items[i].price + items[j].price
  if totalPrice > targetPrice {
    fmt.Println("Not possible")
    return
  }

  // Increment the greater index while the price stays below the target price.
  for totalPrice < targetPrice {
    j++

    // If we've overindexed the greater index, break out of the loop and avoid calculating a new total price. The
    // greater index will be decremented to a valid value later.
    if j >= len(items) {
      break
    }

    totalPrice = items[i].price + items[j].price
  }

  // If we've hit the target price exactly, print the pair and return. We can guarantee the greater index is not
  // overindexed here because totalPrice == targetPrice implies totalPrice < targetPrice is false, meaning that we broke
  // out of the above loop before overindexing was possible.
  if totalPrice == targetPrice {
    printItems(i, j, items)
    return
  }

  // totalPrice is greater than targetPrice, so go back to the previous targetPrice.
  j--
  totalPrice = items[i].price + items[j].price

  // Increment the lesser index while the price stays below the target price.
  for totalPrice < targetPrice {
    i++

    // The two items must be distinct, so print the previous pair and return.
    if i == j {
      printItems(i-1, j, items)
      return
    }

    totalPrice = items[i].price + items[j].price
  }

  // If we've hit the target price exactly, print the pair and return.
  if totalPrice == targetPrice {
    printItems(i, j, items)
    return
  }

  // totalPrice is greater than targetPrice, so print the previous pair.
  printItems(i-1, j, items)
}
