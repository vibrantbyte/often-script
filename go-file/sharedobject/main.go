package main

/*
#include <stdio.h>
#include "libvibrant.h"
#cgo linux CFLAGS: -L./ -I./
#cgo linux LDFLAGS: -L./ -I./ -lhello
*/
import "C"

import (
	"fmt"
)

func main() {
	str := C.Sum(1,2)
	fmt.Println(str)
}