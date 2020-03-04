package spruce

import (
	"strings"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := map[string]struct {
		input   [][]byte
		want    []byte
		wantErr error
	}{
		"merge-append-and-prepend-hint": {
			[][]byte{
				[]byte("key: value\narr:\n- initial"),
				[]byte("otherKey: value\narr:\n- (( prepend ))\n- top"),
				[]byte("arr:\n- (( append ))\n- last"),
				[]byte("key: override"),
			},
			[]byte("arr:\n- top\n- initial\n- last\nkey: override\notherKey: value\n"),
			nil,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotErr := Merge(test.input...)
			if gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr.Error()) {
				t.Errorf("expected %#v, got: %#v", test.wantErr, gotErr)
			}
			if string(test.want) != string(got) {
				t.Errorf("expected %#v, got: %#v", string(test.want), string(got))
			}
		})
	}
}
