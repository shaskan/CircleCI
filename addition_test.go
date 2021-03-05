package addition

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T){
	sum := Addition(4, 3)
	assert.Equal(t, 7, sum, "Sum not the same")
}