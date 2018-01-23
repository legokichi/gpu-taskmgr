package io

import (
  "fmt"
  "stringset"
)

// +gen set
type String struct{}

type IO struct {
  savefile string
  input_dir string
  output_dir string
  queuing: []string
  //processing: Set<string>;
  //processed: Set<string>;
  //ignored: {[path: string]:  [number, string]};
}
/*
func (r *IO) Load() *IO, error {
  return r.width * r.height
}

func (r *IO) Fetch() *IO, error {
  return 2*r.width + 2*r.height
}
func (r *IO) Get() *IO, error {

}
func (r *IO) Commit() *IO, error {

}
func (r *IO) Ignore() *IO, error {

}
func (r *IO) Save() *IO, error {
  
}

func New() *IO {
  io := new(IO)
  return p
}
*/