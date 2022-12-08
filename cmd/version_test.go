package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestAbc(t *testing.T) {
	cmd := NewABCCMD()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	output := string(out)
	fmt.Println(output)
	if output != "abc test" {
		t.Fatalf("expected \"%s\" got \"%s\"", "abc test", output)
	}
}

func TestVersion(t *testing.T) {
	expected := "v.john.nate"
	cmd := NewVersionCMD(expected)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	output := string(out)
	fmt.Println(output)
	if output != expected {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, output)
	}
}
