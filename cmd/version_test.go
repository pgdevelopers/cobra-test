package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	cmd := versionCmd("version test")
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	output := string(out)
	fmt.Println(output)
	if string(out) != "version test" {
		t.Fatalf("expected \"%s\" got \"%s\"", "version test", output)
	}
}
