package formatter

type Config struct {
	IndentSize    uint8 `mapstructure:"indent-size"`
	MaxLineLength uint8 `mapstructure:"max-line-length"`

	Highlight bool
	Spaces    Spaces

	line int
}

type Spaces struct {
	Around Around
}

type Around struct {
	UnaryOperator          bool
	MultiplicativeOperator bool
}
