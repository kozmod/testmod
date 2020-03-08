package testmod

import "fmt" 

// Hi returns a friendly greeting
func Hi(name string) string {
   return fmt.Sprintf("Hi, %s", name)
}

func Hi2(name string) string {
  return fmt.Sprintf("Hi2, %s", name)
}
