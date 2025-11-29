package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	// success - under size
	text := trim("text", 5)
	assert.Equal(t, "text", text)

	// success - at size
	text = trim("text", 4)
	assert.Equal(t, "text", text)

	// success - above size
	text = trim("text", 3)
	assert.Equal(t, "tex", text)
}

func TestBody(t *testing.T) {
	// success
	body := Body("\tbody\n", 4)
	assert.Equal(t, "body", body)
}

func TestHash(t *testing.T) {
	// success
	hash := Hash("text")
	assert.Equal(t, "982d9e3eb996f559e633f4d194def3761d909f5a3b647d1a851fead67c32c9d1", hash)
}

func TestJoin(t *testing.T) {
	// success
	join := Join("form", "alpha", "bravo")
	assert.Equal(t, "form.alpha.bravo", join)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME\n", 4)
	assert.Equal(t, "name", name)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(1000, 0).Local()

	// success
	tobj := Time("\t1000\n")
	assert.Equal(t, want, tobj)
}

func TestUUID(t *testing.T) {
	// success - new uuid
	uuid := UUID()
	assert.Len(t, uuid, 32)

	// success - existing uuid
	uuid = UUID("\tAAAABBBBCCCCDDDD1111222233334444\n")
	assert.Equal(t, "aaaabbbbccccdddd1111222233334444", uuid)
}

func TestUnix(t *testing.T) {
	// setup
	tobj := time.Unix(1000, 0).Local()

	// success
	unix := Unix(tobj)
	assert.Equal(t, "1000", unix)
}
