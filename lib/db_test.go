//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  τλ,pkg,utils_test.go, does,Regression tests on the package functions 

    package lib

    import ("testing"; e "errors"; "fmt"; "net/http"; "io/ioutil"; )

//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  TEST SUITES

    func Test(t *testing.T) {

package main

import (
)

func main() {

  url := "http://localhost:8051/persons"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
  }
  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}
