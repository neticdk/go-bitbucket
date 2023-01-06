package bitbucket

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateTimeUnmarshal(t *testing.T) {
	d := &DateTime{}
	err := d.UnmarshalJSON([]byte(`1672996351030`))
	assert.NoError(t, err)
	assert.Equal(t, int64(1672996351), time.Time(*d).Unix())
}

func TestDateTimeMarshall(t *testing.T) {
	ts := time.Now()
	d := DateTime(ts)
	b, err := d.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(fmt.Sprintf("%d", ts.Unix()*1000)), b)
}
