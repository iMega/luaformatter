package formatter

import (
	"io"
	"reflect"
	"testing"
)

type fakeStatements struct {
	ID int
}

func (fakeStatements) InnerStatement(prev, cur *element) (bool, statement) {
	return false, nil
}

func (fakeStatements) Append(*element) {}

func (fakeStatements) AppendStatement(statement) {}

func (fakeStatements) IsEnd(prev, cur *element) (bool, bool) { return false, false }

func (fakeStatements) TypeOf() typeStatement { return 0 }

func (fakeStatements) GetBody(prevSt statement, cur *element) statement { return nil }

func (fakeStatements) GetStatement(prev, cur *element) statement { return nil }

func (fakeStatements) Format(*Config, printer, io.Writer) error { return nil }

func Test_chainStatments_Prev(t *testing.T) {
	tests := []struct {
		name   string
		csFunc func() *chainStatments
		want   statement
	}{
		{
			name: "chain is empty returns nil",
			csFunc: func() *chainStatments {
				return &chainStatments{}
			},
			want: nil,
		},
		{
			name: "chain with one statement returns nil",
			csFunc: func() *chainStatments {
				cs := chainStatments{}
				cs.Append(&fakeStatements{ID: 100})

				return &cs
			},
			want: &fakeStatements{
				ID: 100,
			},
		},
		{
			name: "chain with two statements returns prev statement",
			csFunc: func() *chainStatments {
				cs := chainStatments{}
				cs.Append(&fakeStatements{ID: 100})
				cs.Append(&fakeStatements{ID: 200})

				return &cs
			},
			want: &fakeStatements{
				ID: 200,
			},
		},
		{
			name: "optimistic returns prev statement",
			csFunc: func() *chainStatments {
				cs := &chainStatments{}
				cs.Append(&fakeStatements{ID: 100})
				cs.Append(&fakeStatements{ID: 200})
				cs.Append(&fakeStatements{ID: 300})

				return cs
			},
			want: &fakeStatements{
				ID: 300,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.csFunc().Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chainStatments.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}
