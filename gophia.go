package main

// #include "gophia.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Data interface {
	Open() error
	Destroy() error
}

type Object struct {
	Obj unsafe.Pointer
}

type g_env struct {
	Env unsafe.Pointer
}

// sp_env
func Sp_env() (*g_env, error) {
	e, err := C.sp_env()
	if err != nil {
		return nil, fmt.Errorf("sp_env() failed")
	}

	return &g_env{Env: e}, nil
}

func (g *g_env) Open() error {
	_, err := C.g_open(g.Env)
	return err
}
func Sp_open(g *g_env) error {
	return g.Open()
}

func (g *g_env) Destroy() error {
	_, err := C.g_destroy(g.Env)
	return err
}
func Sp_destroy(g *g_env)

type g_ctl struct {
	Ctl unsafe.Pointer
}

// sp_ctl
func Sp_ctl(env *g_env) (*g_ctl, error) {
	ctl, err := C.g_ctl(env.Env)
	if err != nil {
		return nil, fmt.Errorf("sp_ctl() failed")
	}

	return &g_ctl{Ctl: ctl}, nil
}

// sp_object
func Sp_object(obj Object) (*Object, error) {
	o, err := C.g_object(obj.Obj)
	if err != nil {
		return nil, fmt.Errorf("sp_object() failed")
	}

	return &Object{Obj: o}, nil
}

// sp_set
func Sp_set(ctl *g_ctl, v ...interface{}) error {
	_, err := C.g_set__cfg(
		ctl.Ctl,
		C.CString("sophia.path"),
		C.CString("./sophia"),
	)

	return err
}

// sp_get
func Sp_get(ctl *g_ctl, args ...interface{}) (string, error) {
	var vp unsafe.Pointer
	var err error

	switch len(args) {
	case 0:
		return "", fmt.Errorf("sp_get failed: missing arguments")
	case 1:
		key := C.CString(args[0].(string))
		vp, err = C.g_get__key(ctl.Ctl, key)
		if err != nil {
			return "", fmt.Errorf("sp_get failed: g_get__key failed")
		}
		C.free(unsafe.Pointer(key))
	}

	v := C.g__vtoc(vp)
	return C.GoString(v), nil
}

func main() {
	env, err := Sp_env()
	if err != nil {
		panic(err)
	}

	ctl, err := Sp_ctl(env)
	if err != nil {
		panic(err)
	}
	Sp_set(ctl, "sophia.path", "./storage")
	Sp_set(ctl, "db", "test")
	Sp_set(ctl, "db.test.index.cmp", "string") /* string by default */
	v, err := Sp_get(ctl, "sophia.path")
	env.Open()
	defer env.Destroy()

	fmt.Printf("%#v", v)
}
