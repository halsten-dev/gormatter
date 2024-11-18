package gormatter

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	log.Println(Bold(Black, Red, "ATTENTION PLEASE !"))
	log.Println(Normal(Black, HighRed, "Attention !"))
}
