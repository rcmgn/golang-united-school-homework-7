package coverage

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestNew(t *testing.T) {
	m, ok := New("1 2 3\n4 5 6\n7 8 9")
	fmt.Println(m, ok)
}

func TestMatrix_Rows(t *testing.T) {
	m, _ := New("1 2 3\n4 5 6\n7 8 9")
	//if ok != nil {
	fmt.Println(m.Rows())
	//}
}

func TestMatrix_Cols(t *testing.T) {
	m, _ := New("1 2 3\n4 5 6\n7 8 9")
	//if ok != nil {
	fmt.Println(m.Cols())
	//}
}
func TestMatrix_Set(t *testing.T) {
	m, _ := New("1 2 3\n4 5 6\n7 8 9")
	//if ok != nil {
	m.Set(1, 1, 100)
	fmt.Println(m)
	//}

}

func TestPeople_Len(t *testing.T) {
	p := People{{"Name1", "surname1", time.Now()}, {"Name2", "surname2", time.Now()}}
	fmt.Println(p)
	fmt.Println(p.Len())
}

func TestPeople_Swap(t *testing.T) {
	p := People{{"Name1", "surname1", time.Now()}, {"Name2", "surname2", time.Now()}}
	fmt.Println(p)
	p.Swap(0, 1)
	fmt.Println(p)
}
