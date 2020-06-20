package test

import "testing"

func TestAdd(t *testing.T) {
	//type args struct {
	//	a int
	//}
	tests := []struct {
		name string
		args int
		want int
	}{
		{"case 1",1,2},
		{"case 1",2,3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinus(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Minus(tt.args.a); got != tt.want {
				t.Errorf("Minus() = %v, want %v", got, tt.want)
			}
		})
	}
}