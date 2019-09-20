package main

import "testing"

//a test file should have _test postfix
//a test function should have Test prefix
//a test function should take only one argument (t *testing.T) where we can t.Errorf(), t.Fail()

func TestHello1(t *testing.T) {
	got := Hello()
	want := "Hello, Go!"

	//anytime we can use t.Errorf to report an error

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got2 := HelloString("Go!")
	want2 := "Hello, Go!"

	if got2 != want2 {
		t.Errorf("got %q want %q", got2, want2)
	}
}

func TestHello2(t *testing.T) {

	t.Run("saying hello to go", func(t *testing.T) {
		got := Hello()
		want := "Hello, Go!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying hello to someone", func(t *testing.T) {
		got2 := HelloString("Go!")
		want2 := "Hello, Go!"

		if got2 != want2 {
			t.Errorf("got %q want %q", got2, want2)
		}
	})

}

func TestHello3(t *testing.T) {

	//t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems easier. 
	
    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := HelloString("Go!")
        want := "Hello, Go!"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'World'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

}
