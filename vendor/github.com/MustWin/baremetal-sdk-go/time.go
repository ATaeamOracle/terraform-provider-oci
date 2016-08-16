package baremetal

import "time"

type Time struct {
	time.Time
}

const baremetalTimeFormat = "2006-01-02T15:04:05.999+0000"

func (t *Time) UnmarshalJSON(data []byte) (e error) {
	t.Time, e = time.Parse(`"`+baremetalTimeFormat+`"`, string(data))
	return
}

func (t *Time) MarshalJSON() (buff []byte, e error) {
	s := t.Format(baremetalTimeFormat)
	buff = []byte(`"` + s + `"`)
	return
}
