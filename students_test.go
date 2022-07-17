package coverage

import (
	"os"
	"sort"
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

func TestPeople_Len(t *testing.T) {
	var pip People
	pip = append(pip, Person{"Name1", "surname1", time.Now()})
	pip = append(pip, Person{"Name2", "surname2", time.Now()})
	if pip.Len() != len(pip) {
		t.Error("Error func Len")
	}
}

func TestPeople_Less(t *testing.T) {
	var pip People
	pip = append(pip, Person{"Name1", "surname1", time.Now()})
	pip = append(pip, Person{"Name2", "surname2", time.Now()})
	if !pip.Less(0, 1) {
		t.Error("Error func Less")
	}
}

func TestPeople_Less2(t *testing.T) {
	var pip People
	age := time.Now()
	pip = append(pip, Person{"Name", "surname1", age})
	pip = append(pip, Person{"Name", "surname2", age})
	if !pip.Less(0, 1) {
		t.Error("Error func Less2")
	}
}

func TestPeople_Swap(t *testing.T) {
	var pip People
	p0 := Person{"Name1", "surname1", time.Now()}
	p1 := Person{"Name2", "surname2", time.Now()}
	pip = append(pip, p0)
	pip = append(pip, p1)
	pip.Swap(0, 1)
	if !((pip[0] == p1) && (pip[1] == p0)) {
		t.Error("Error func Swap")
	}
}

func TestPeople_Sort(t *testing.T) {
	var pip People
	p0 := Person{"Name1", "surname1", time.Now()}
	p1 := Person{"Name2", "surname2", time.Now()}
	p2 := Person{"Name3", "surname3", time.Now()}
	pip = append(pip, p1, p2, p0)
	sort.Sort(pip)
	if (pip[0] != p0) || (pip[1] != p1) || (pip[2] != p2) {
		t.Error("Error sort.Sort")
	}
}

func TestNew(t *testing.T) {
	m, _ := New("1 2 3 \n 4 5 6 \n 7 8 9 \n 10 11 12")
	if (m.Rows()[3][2] != 12) || (m.Rows()[1][2] != 6) {
		t.Error("Error New Matrix")
	}
}

func TestMatrix_Rows(t *testing.T) {
	m, _ := New("1 2 3 \n 4 5 6 \n 7 8 9 \n 10 11 12")
	if (m.Rows()[3][2] != 12) || (m.Rows()[1][2] != 6) {
		t.Error("Error Rows")
	}
}

func TestMatrix_Cols(t *testing.T) {
	m, _ := New("1 2 3 \n 4 5 6 \n 7 8 9 \n 10 11 12")
	if (m.Cols()[1][1] != 5) || (m.Cols()[2][0] != 3) {
		t.Error("Error Cols")
	}
}
func TestMatrix_Set(t *testing.T) {
	var m *Matrix
	m, _ = New("1 2 3 \n 4 5 6 \n 7 8 9 \n 10 11 12")
	if !m.Set(1, 1, 100) {
		t.Error("Error Set1")
	}
	if !m.Set(1, 2, 100) {
		t.Error("Error Set2")
	}
	if m.Set(2, 3, 100) {
		t.Error("Error Set3")
	}
}
