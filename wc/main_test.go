package main

import (
    "bytes"
    "testing"
)

func TestCountLines(t *testing.T) {
    b := bytes.NewBufferString("wword1 word2 word3 word4\nword1 word2 word3 word4\nord1 word2 word3 word4\n")
    exp := 3
    res := count(b, true, false)
    if res != exp {
        t.Errorf("Expected %d, got %d\n", exp, res)
    }

}

func TestCountBytes(t *testing.T) {
    b := bytes.NewBufferString("word1 word2 word3 word4\n")
    exp := 24
    res := count(b, false, true)
    if res != exp {
        t.Errorf("Expected %d, got %d\n", exp, res)
    }

}

func TestCountWords(t *testing.T) {
    b := bytes.NewBufferString("word1 word2 word3 word4\n")
    exp := 4
    res := count(b, false, false)
    if res != exp {
        t.Errorf("Expected %d, got %d\n", exp, res)
    }

}
