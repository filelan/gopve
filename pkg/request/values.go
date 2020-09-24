package request

import (
	"net/url"
	"strconv"
	"time"

	"github.com/xabinapal/gopve/internal/types"
)

type Marshaler interface {
	Marshal() (string, error)
}

type Unmarshaler interface {
	Unmarshal(s string) error
}

type Values url.Values

func (v Values) AddString(k string, s string) {
	v[k] = []string{s}
}

func (v Values) AddInt(k string, i int) {
	v[k] = []string{strconv.Itoa(i)}
}

func (v Values) AddUint(k string, u uint) {
	v[k] = []string{strconv.Itoa(int(u))}
}

func (v Values) AddBool(k string, b bool) {
	v[k] = []string{types.PVEBool(b).String()}
}

func (v Values) AddTime(k string, t time.Time) {
	v[k] = []string{t.Format("2006-01-02 15:04:05")}
}

func (v Values) AddObject(k string, o Marshaler) error {
	val, err := o.Marshal()
	if err != nil {
		return err
	}

	v[k] = []string{val}

	return nil
}

func (v Values) ConditionalAddString(k string, s string, cond bool) {
	if cond {
		v.AddString(k, s)
	}
}

func (v Values) ConditionalAddInt(k string, i int, cond bool) {
	if cond {
		v.AddInt(k, i)
	}
}

func (v Values) ConditionalAddUint(k string, u uint, cond bool) {
	if cond {
		v.AddUint(k, u)
	}
}

func (v Values) ConditionalAddBool(k string, b bool, cond bool) {
	if cond {
		v.AddBool(k, b)
	}
}

func (v Values) ConditionalAddTime(k string, t time.Time, cond bool) {
	if cond {
		v.AddTime(k, t)
	}
}

func (v Values) ConditionalAddObject(k string, t Marshaler, cond bool) {
	if cond {
		v.AddObject(k, t)
	}
}
