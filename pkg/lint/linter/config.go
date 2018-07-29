package linter

const (
	PresetFormatting  = "format"
	PresetComplexity  = "complexity"
	PresetStyle       = "style"
	PresetBugs        = "bugs"
	PresetUnused      = "unused"
	PresetPerformance = "performance"
)

type Config struct {
	Linter           Linter
	EnabledByDefault bool
	DoesFullImport   bool
	NeedsSSARepr     bool
	InPresets        []string
	Speed            int // more value means faster execution of linter

	OriginalURL string // URL of original (not forked) repo, needed for autogenerated README
}

func (lc Config) WithFullImport() Config {
	lc.DoesFullImport = true
	return lc
}

func (lc Config) WithSSA() Config {
	lc.DoesFullImport = true
	lc.NeedsSSARepr = true
	return lc
}

func (lc Config) WithPresets(presets ...string) Config {
	lc.InPresets = presets
	return lc
}

func (lc Config) WithSpeed(speed int) Config {
	lc.Speed = speed
	return lc
}

func (lc Config) WithURL(url string) Config {
	lc.OriginalURL = url
	return lc
}

func (lc Config) NeedsProgramLoading() bool {
	return lc.DoesFullImport
}

func (lc Config) NeedsSSARepresentation() bool {
	return lc.NeedsSSARepr
}

func (lc Config) GetSpeed() int {
	return lc.Speed
}

func NewConfig(linter Linter) *Config {
	return &Config{
		Linter: linter,
	}
}