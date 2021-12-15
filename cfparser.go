package cfparser

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Use large buffer for fast reading.
const bufSize = 1024

type CFParser struct {
	reader *bufio.Reader
	tree   *CPair
	line   string
	ignore string
	split  byte
}

// Most user should use this function to initialize a CFParser.
func NewCFParser(bindFile *os.File, IgnorePrefix string, split byte) *CFParser {
	parser := new(CFParser)
	parser.Bind(bindFile)
	parser.Ignore(IgnorePrefix)
	parser.Cut(split)
	return parser
}

// Bind a file to the CFParser and user should close the file correctly.
func (cfp *CFParser) Bind(file *os.File) {
	cfp.reader = bufio.NewReaderSize(file, bufSize)
}

// Tell parser to ignore lines starting with specific prefix.
func (cfp *CFParser) Ignore(prefix string) {
	cfp.ignore = prefix
}

// Use split to cut line to key and value.
func (cfp *CFParser) Cut(split byte) {
	cfp.split = split
}

// Start to read. Return the number of valid lines without comment and return -1 when missing parameters.
// NewCFParser contains all the parameters this function needed and ignore empty prefix is invalid.
func (cfp *CFParser) ReadAll() int {
	if cfp.reader == nil || len(cfp.ignore) == 0 || cfp.split == 0 {
		return -1
	}

	validCount := 0
	for cfp.readNext(&validCount) {
	}
	return validCount
}

// Find pair and get it. Return nil when can't find anything.
func (cfp *CFParser) Get(key string) *CPair {
	return get(cfp.tree, strings.ToLower(strings.TrimSpace(key)))
}

// Put a config pair by yourself to provide default config.
// You can also turn on the signalMode
// Most users don't need to do this, use ReadAll instead.
func (cfp *CFParser) Put(key string, val string) {
	key = strings.ToLower(strings.TrimSpace(key))
	val = strings.TrimSpace(val)
	cfp.tree = put(cfp.tree, key, val)
}

// TODO(AoXiang Li): Write the config tree to a file with standard format.
func (cfp *CFParser) Dump() {
}

func (cfp *CFParser) isComment() bool {
	return len(cfp.line) == 0 || strings.HasPrefix(cfp.line, cfp.ignore)
}

func (cfp *CFParser) readNext(validCount *int) bool {
	line, err := cfp.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return false
	}

	cfp.line = strings.TrimSpace(line)

	// Ignore comments.
	if !cfp.isComment() {
		// Find the first separator.
		find := strings.IndexByte(cfp.line, cfp.split)
		if find == -1 {
			return false
		}

		*validCount++

		// Preprocess the line.
		key := strings.ToLower(strings.TrimSpace(cfp.line[:find]))
		val := strings.TrimSpace(cfp.line[find+1:])
		cfp.tree = put(cfp.tree, key, val)
	}

	return err == nil
}
