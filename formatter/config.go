package formatter

type Config struct {
	IndentSize    uint8 `mapstructure:"indent-size"`
	MaxLineLength uint8 `mapstructure:"max-line-length"`

	Highlight bool

	Spaces    Spaces
	Alignment Alignment `mapstructure:"alignment"`
}

func DefaultConfig() Config {
	return Config{
		IndentSize:    4,
		MaxLineLength: 80,
	}
}

type Spaces struct {
	Around Around
}

type Around struct {
	UnaryOperator          bool
	MultiplicativeOperator bool
}

type Alignment struct {
	Table AlignmentTable `mapstructure:"table"`
}

// AlignmentTable formatting tables in code
type AlignmentTable struct {
	// KeyValue = true
	// t = {
	//     key1   = value1,
	//     key10  = value10,
	//     key100 = value100,
	// }
	KeyValuePairs bool `mapstructure:"key-value-pairs"`

	// Comments = true
	// t = {
	//     key1 = value1,     -- comment
	//     key10 = value10,   -- comment
	//     key100 = value100, -- comment
	// }
	Comments bool `mapstructure:"comments"`
}
