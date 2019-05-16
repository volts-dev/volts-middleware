package polymer_import_rewrite

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/volts-dev/lexer"
)

type (
	parser struct {
		lexer          *lexer.TLexer
		pos            int // the pos of token
		tokens         []*lexer.TToken
		root           string // the name of root dir but not path
		path           string
		buffer         bytes.Buffer
		in_node_module bool // is it in node_modules
		in_same_module bool // is it in same modules
	}
)

// #新建解析器可以独立于TTemplate使用
func NewParser() *parser {
	return &parser{}
}

// Returns tokens[i] or NIL (if i >= len(tokens))
func (p *parser) get(i int) *lexer.TToken {
	if i < len(p.tokens) {
		//fmt.Println("Get", i, p.tokens[i])
		return p.tokens[i]
	}
	return nil
}

// Consume one token. It will be gone forever.
func (p *parser) next() {
	p.nextn(1)
}

// Consume N tokens. They will be gone forever.
func (p *parser) nextn(count int) {
	p.pos += count
}

// Returns the UNCONSUMED token count.
func (p *parser) remaining() int {
	return len(p.tokens) - p.pos
}

// Returns the current token.
func (p *parser) current() *lexer.TToken {
	return p.get(p.pos)
}

func (p parser) Buffer() bytes.Buffer {
	return p.buffer
}
func (p *parser) writeToBuf() {
	for _, t := range p.tokens {
		p.buffer.WriteString(t.Val)
	}
	p.buffer.WriteString("\n")

}

// Returns the current token.
func (p *parser) rest() {
	p.pos = 0
	p.in_node_module = false
	p.in_same_module = false
	p.tokens = make([]*lexer.TToken, 0)
}

// set the root dir name for reference
func (p *parser) SetRoot(dir string) {
	p.root = dir
}

// set the file full path for reference
func (p *parser) SetPath(path string) {
	p.path = path
}

// Returns the prior token.
func (p *parser) __Prior() *lexer.TToken {
	return p.get(p.pos - 1)
}

// Returns the CURRENT token if the given type AND value matches.
// It DOES NOT consume the token.
func (p *parser) peek(typ int, val string) *lexer.TToken {
	return p.peekN(0, typ, val)
}

// Returns the tokens[current position + shift] token if the
// given type AND value matches for that token.
// DOES NOT consume the token.
func (p *parser) peekN(shift int, typ int, val string) *lexer.TToken {
	t := p.get(p.pos + shift)
	if t != nil {
		//fmt.Println("PeekN", t)
		if t.Type == typ && t.Val == val {
			return t
		}
	}
	return nil
}

// Returns the CURRENT token if the given type matches.
// It DOES NOT consume the token.
func (p *parser) peekType(typ ...int) *lexer.TToken {
	return p.peekTypeN(0, typ...)
}

// Returns the tokens[current position + shift] token if the given type matches.
// DOES NOT consume the token for that token.
func (p *parser) peekTypeN(shift int, typ ...int) *lexer.TToken {
	t := p.get(p.pos + shift)
	if t != nil {
		if t.Type == lexer.SEMICOLON {
			return nil
		}
		//fmt.Println("PeekTypeN", t)
		for _, ty := range typ {
			if t.Type == ty {
				return t
			}
		}
	}
	return nil
}

func (p *parser) __skipWhitespace() {
	for {
		if p.current().Type != lexer.SAPCE {
			return
		}
		fmt.Println(lexer.TokenNames[p.current().Type])
		p.next()
	}

	return
}

func (self *parser) replace_import() {
	for self.remaining() > 0 {
		t := self.peekType(lexer.STRING)
		if t != nil {

			// make sure it is xxx/xxx but not /xxx/xx
			if self.path[0] == '/' {
				self.path = self.path[1:]
			}
			strs := strings.Split(self.path, "/")    // full url path without host
			pkg_name := strings.Split(t.Val, "/")[0] // package name
			//fmt.Println("impt", self.in_node_module, pkg_name, t.Line)
			// only @ and not . will be changed
			if pkg_name[0] == '@' || pkg_name[0] != '.' {
				if strings.Index(self.path, pkg_name) > -1 {
					self.in_same_module = true
				}
				var ext string
				var cnt int
				if self.in_node_module && self.in_same_module {
					for i, str := range strs {
						if str == pkg_name {
							cnt = len(strs) - i - 2
							//fmt.Println("str1", len(strs), strs, i, cnt)

							break
						}
					}

					if cnt < 0 {
						panic("repeat counter must biger than 0")
					}

					ext = strings.Repeat("../", cnt)
					t.Val = strings.Replace(t.Val, pkg_name+"/", ext, -1)

				} else if self.in_node_module {
					for i, str := range strs {
						if str == self.root {
							cnt = len(strs) - i
							//fmt.Println("str2", len(strs), strs, i, cnt)

							break
						}
					}
					cnt-- //len()-1
					cnt-- //node_modules
					cnt-- //file
					if cnt < 0 {
						panic("repeat counter must biger than 0")
					}
					ext = strings.Repeat("../", cnt)
					t.Val = ext + t.Val

					//t.Val = strings.Replace(t.Val, pkg_name+"//", ext, -1)
				} else {
					for i, str := range strs {
						if str == self.root {
							cnt = len(strs) - i - 2
							//fmt.Println("str3", len(strs), strs, i, cnt)

							break
						}
					}
					if cnt < 0 {
						panic("repeat counter must biger than 0")
					}
					//fmt.Println("str3", len(strs), strs, cnt)

					ext = strings.Repeat("../", cnt)
					t.Val = ext + "node_modules/" + t.Val
				}

				if !strings.HasSuffix(t.Val, ".js") {
					t.Val = t.Val + ".js"
				}
			}

			if pkg_name[0] == '.' {
				if !strings.HasSuffix(t.Val, ".js") {
					t.Val = t.Val + ".js"
				}
			}

			//fmt.Println("str4", t.Val)
			return
		}

		self.next()
	}

}

func (self *parser) parse() {
	if strings.Index(self.path, "node_modules") > -1 {
		self.in_node_module = true
	}

	//self.skipWhitespace()
	//t := self.current()
	//fmt.Println(t.Val)
	for self.remaining() > 0 {
		t := self.peek(lexer.IDENT, "import")
		if t != nil {
			//TODO 根据精确 import 位置确保无误
			// test the next token
			next := self.get(self.pos + 1)
			if next.Type == lexer.SAPCE {
				self.replace_import()
			}

		}

		self.next()
	} /*
		if t.Type == lexer.IDENT && t.Val == "import" {
			//idx := strings.Index(self.path, self.Root)
			self.replace_import()
		}
	*/
}
func (self *parser) Parse(input io.Reader) {
	// new the line scanner
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		t := scanner.Text()

		// only resolve the import line
		t = strings.TrimLeft(t, " ")
		if strings.HasPrefix(t, "import") {
			self.rest()
			lex, err := lexer.NewLexer(strings.NewReader(t))
			if err != nil {
				panic(err.Error())
			}

			self.lexer = lex
			for {
				token, ok := <-lex.Tokens
				if !ok {
					break
				}

				self.tokens = append(self.tokens, &token)
				//fmt.Println(lexer.PrintToken(token))
			}
			self.parse()
			self.writeToBuf()

		} else {
			self.buffer.WriteString(t + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//fmt.Println(self.buffer.String())
}
