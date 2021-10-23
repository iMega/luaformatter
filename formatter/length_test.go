package formatter

import (
	"testing"
)

func TestStatementLength(t *testing.T) {
	type args struct {
		c    *Config
		code []byte
		p    printer
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{
			args: args{
				c: &Config{
					IndentSize:    4,
					MaxLineLength: 80,
					Alignment: Alignment{
						Table: AlignmentTable{
							KeyValuePairs: true,
							Comments:      true,
						},
					},
				},
				code: []byte(`
return (machine and machine.is_loaded_qwe) and "coffee brewing" or "fill your water"
                `),
				p: printer{},
			},
			want:    84,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := parse(tt.args.code)
			if err != nil {
				t.Errorf("parse() error = %v", err)
				return
			}

			got, err := StatementLength(tt.args.c, doc.Body, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("StatementLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("StatementLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
