package blacklist

import (
  "os"
  "sort"
  "bufio"
)

type Domain struct {
  Name string
}

type Blacklist []Domain

// implement Sort interface for Blacklist
func (blacklist Blacklist) Len() int {
  return len(blacklist)
}

func (blacklist Blacklist) Less(i, j int) bool {
  // ascending alphabetical sort
  return blacklist[i].Name < blacklist[j].Name
}

func (blacklist Blacklist) Swap(i, j int) {
  blacklist[i], blacklist[j] = blacklist[j], blacklist[i]
}


func New(path string) (Blacklist, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  blacklist := Blacklist{}
  for scanner.Scan() {
    blacklist = append(blacklist, Domain{Name: scanner.Text()})
  }

  sort.Sort(blacklist)

  return blacklist, scanner.Err()
}
