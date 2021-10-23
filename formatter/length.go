package formatter

import "bytes"

func StatementLength(c *Config, st statement, p printer) (uint8, error) {
	buf := bytes.NewBuffer(nil)
	if err := st.Format(c, p, buf); err != nil {
		return 0, err
	}

	l := uint8(buf.Len())
	buf.Reset()

	return l, nil
}
