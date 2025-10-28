package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello Go"

	got := hello()

	if got != want {
		t.Fatalf("got %q; want %q", got, want)
	}
}

//…or push an existing repository from the command line
//git remote add origin https://github.com/sk8work/hello-go.git
//git branch -M main
//git push -u origin main
