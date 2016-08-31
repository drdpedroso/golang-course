package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	words("is2 Thi1s T4est 3a");
}

func words (s string) {
    a := strings.Split(s, " ")
    var bla = make(map[string]string)
    for i := 0; i < len(a); i++ {
        for y := 0; y < len(a[i]); y++ {
            if _, err := strconv.Atoi(string(a[i][y])); err == nil {
                bla[string(a[i][y])] = string(a[i])
            }
        }
    }
    for x := 1; x <= len(bla); x++ {
        var xT = strconv.Itoa(x);
        fmt.Println(bla[xT])
    }
}
