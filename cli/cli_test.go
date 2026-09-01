package cli

import (
	"testing"

	"github.com/gregyr/go-utils/std"
)

func TestAddFlagAndAddOptAddCorrectly(t *testing.T) {
	c := New("test", "")

	t.Run("Test Flags", func(t *testing.T) {
		c.AddFlag(NewFlag("t", "test", "test info"))
		if len(c.flags) != 1 {
			t.Fatal("Expected length of Flags to be 1")
		}
		flag := c.flags[0]
		if flag.short != "t" {
			t.Errorf("Expected short version of flag to be 't'")
		}
		if flag.long != "test" {
			t.Errorf("Expected long version of flag to be 'test'")

		}
		if flag.info != "test info" {
			t.Errorf("Expected info of flag to be 'test info'")

		}
	})

	t.Run("Test Opts", func(t *testing.T) {
		c.AddOpt(NewOpt("o", "opttest", "opttest info"))
		if len(c.opts) != 1 {
			t.Fatal("Expected length of Opts to be 1")
		}
		opt := c.opts[0]
		if opt.short != "o" {
			t.Errorf("Expected short version of flag to be 'o'")
		}
		if opt.long != "opttest" {
			t.Errorf("Expected long version of flag to be 'opttest'")

		}
		if opt.info != "opttest info" {
			t.Errorf("Expected info of flag to be 'opttest info'")

		}
	})
}

func TestParseArgsReturnsEverythingFalseForNoArgs(t *testing.T) {
	c := New("test", "")

	c.AddFlag(NewFlag("t", "test", "test info"))
	c.AddOpt(NewOpt("o", "testopt", "testop info"))

	err := c.parseArgs([]string{})

	if err != nil {
		t.Fatalf("Failed with error: %s", err)
	}
	if !std.AllEqual(c.flagstates, false) {
		t.Error("Expected all flags to be false")
	}
}

func TestParseArgsSetsFlags(t *testing.T) {
	c := New("test", "")

	c.AddFlag(NewFlag("t", "test", "test info"))
	c.AddFlag(NewFlag("n", "notset", "this should not be set"))

	t.Run("Short Version", func(t *testing.T) {
		c.parseArgs([]string{"test", "-t"})

		if len(c.flags) != 2 {
			t.Fatalf("Length of flags should be 2")
		}
		flag_t := c.flags[0]
		if flag_t.short != "t" {
			t.Errorf("Expected the first flag to be flag for 't'")
		}
		flag_n := c.flags[1]
		if flag_n.short != "n" {
			t.Errorf("Expected the second flag to be flag for 'n'")
		}
		if len(c.flagstates) != 2 {
			t.Fatalf("Length of flagstates should be 2")
		}
		if c.flagstates[0] != true {
			t.Errorf("Expected flagstate for flag 't' to be true")
		}
		if c.flagstates[1] != false {
			t.Errorf("Expected flagstate for flag 'n' to be false")
		}
	})

	//reset
	c.flagstates = make([]bool, len(c.flags))
	t.Run("Long Version", func(t *testing.T) {
		c.parseArgs([]string{"test", "--test"})

		if len(c.flags) != 2 {
			t.Fatalf("Length of flags should be 2")
		}
		flag_t := c.flags[0]
		if flag_t.short != "t" {
			t.Errorf("Expected the first flag to be flag for 't'")
		}
		flag_n := c.flags[1]
		if flag_n.short != "n" {
			t.Errorf("Expected the second flag to be flag for 'n'")
		}
		if len(c.flagstates) != 2 {
			t.Fatalf("Length of flagstates should be 2")
		}
		if c.flagstates[0] != true {
			t.Errorf("Expected flagstate for flag 'test' to be true")
		}
		if c.flagstates[1] != false {
			t.Errorf("Expected flagstate for flag 'notset' to be false")
		}
	})

}

func TestParseArgsSetsOpts(t *testing.T) {
	c := New("test", "")

	c.AddOpt(NewOpt("t", "test", "test info"))
	c.AddOpt(NewOpt("n", "notset", "this should not be set"))

	t.Run("Short Version", func(t *testing.T) {
		c.parseArgs([]string{"test", "-t", "somevalue"})

		if len(c.opts) != 2 {
			t.Fatalf("Length of opts should be 2")
		}
		opt_t := c.opts[0]
		if opt_t.short != "t" {
			t.Errorf("Expected the first opt to be flag for 't'")
		}
		opt_n := c.opts[1]
		if opt_n.short != "n" {
			t.Errorf("Expected the second opt to be flag for 'n'")
		}
		if len(c.optsVals) != 2 {
			t.Fatalf("Length of flagstates should be 2")
		}
		if c.optsVals[0] != "somevalue" {
			t.Errorf("Expected value for opt 't' to be 'somevalue'")
		}
		if c.optsVals[1] != "" {
			t.Errorf("Expected value for opt 'n' to be empty")
		}
	})

	//reset
	c.optsVals = make([]string, len(c.opts))
	t.Run("Long Version", func(t *testing.T) {
		c.parseArgs([]string{"test", "--test", "somevalue"})

		if len(c.opts) != 2 {
			t.Fatalf("Length of opts should be 2")
		}
		opt_t := c.opts[0]
		if opt_t.short != "t" {
			t.Errorf("Expected the first opt to be flag for 't'")
		}
		opt_n := c.opts[1]
		if opt_n.short != "n" {
			t.Errorf("Expected the second opt to be flag for 'n'")
		}
		if len(c.optsVals) != 2 {
			t.Fatalf("Length of optsVals should be 2")
		}
		if c.optsVals[0] != "somevalue" {
			t.Errorf("Expected value for opt 't' to be 'somevalue'")
		}
		if c.optsVals[1] != "" {
			t.Errorf("Expected value for opt 'n' to be empty")
		}
	})

}
