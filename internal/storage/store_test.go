package storage

import (
	"errors"
	"testing"

	"github.com/jaekwon/testify/require"
)

func Test_Store_Put(t *testing.T) {
	testDataStore := &dataStore{err: errors.New("This is error")}
	st := &Store{db: testDataStore}

	require.NotNil(t, st)
	err := st.Put("", nil)

	require.Error(t, err)
	require.Equal(t, "storage: Store.Put st.db.Put error: This is error", err.Error())
}

func Test_Store_Put_tableTest(t *testing.T) {
	tc := []struct {
		name             string
		input            *dataStore
		expectedErrorMsg string
	}{
		{
			name:             "this should return foo-bar error",
			input:            &dataStore{err: errors.New("foo-bar")},
			expectedErrorMsg: "storage: Store.Put st.db.Put error: foo-bar",
		},
		{
			name:             "this should return bar-baz error",
			input:            &dataStore{err: errors.New("bar-baz")},
			expectedErrorMsg: "storage: Store.Put st.db.Put error: bar-baz",
		},
		{
			name:  "this should not return error",
			input: &dataStore{err: nil},
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {

			st := &Store{db: tc.input}

			require.NotNil(t, st)
			err := st.Put("", nil)

			if tc.expectedErrorMsg != "" {
				require.Equal(t, err.Error(), tc.expectedErrorMsg)
				return
			}

			require.NoError(t, err)

		})
	}

}
