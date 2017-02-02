package main

import css "github.com/gorilla/css/scanner"
import "os"
import "bufio"
import "fmt"

func main () {
  output := make(map[string]int)

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    s := css.New(scanner.Text())
    for {
      token := s.Next()
      if token.Type == css.TokenEOF || token.Type == css.TokenError {
        break
      }

      if token.Type == css.TokenHash {
        if _, ok := output[token.Value]; ok {
          output[token.Value] = output[token.Value] + 1
        } else {
          output[token.Value] = 1
        }
      }
    }
  }

  for k, v := range output {
    fmt.Printf("%s\t %d\n", k, v)
  }
}
