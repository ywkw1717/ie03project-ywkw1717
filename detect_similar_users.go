package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

type UserBookPair [2]int

func (ubp UserBookPair) UserNum() int {
  return ubp[0]
}

func (ubp UserBookPair) BookNum() int {
  return ubp[1]
}

func main() {
  s := bufio.NewScanner(os.Stdin)

  s.Scan()

  inputValue := strings.Split(s.Text(), " ")
  ubp := UserBookPair{}

  for i, v := range inputValue {
    if num, err := strconv.Atoi(v); err != nil {
      fmt.Println(err)
      os.Exit(1)
    } else {
      ubp[i] = num
    }
  }

  fmt.Println(ubp.UserNum())
  fmt.Println(ubp.BookNum())
}
