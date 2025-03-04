package lines

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadLines(t *testing.T) {
	type testCase struct {
		Name     string
		Expected Lines
	}

	testCases := []testCase{
		{
			Name: "Success",
			Expected: Lines{
				Line{indices: []int{
					1, 1, 1, 1, 1,
				}},
				Line{indices: []int{
					0, 0, 0, 0, 0,
				}},
				Line{indices: []int{
					2, 2, 2, 2, 2,
				}},
				Line{indices: []int{
					0, 1, 2, 1, 0,
				}},
				Line{indices: []int{
					2, 1, 0, 1, 2,
				}},
			},
		},
	}

	for _, tc := range testCases {
		res, err := ReadLines()
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, 5, len(res))
		require.Equal(t, tc.Expected, res)
	}
}
