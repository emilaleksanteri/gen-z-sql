package main

var tokens = map[string]string{
	"yoink":      "select",
	"fr":         "=",
	"skibity":    "from",
	"on_god":     "where",
	"goon":       "and",
	"edge":       "or",
	"bruh":       "not",
	"rizz":       "update",
	"short_king": "asc",
	"tall_king":  "desc",
	"yeet":       "delete",
	"slide":      "insert",
}

type parser struct {
	curr      byte
	pos       int
	input     string
	nextToken string
	output    string
}

func (p *parser) next() {
	if p.pos >= len(p.input) {
		p.curr = 0
		return
	}
	p.curr = p.input[p.pos]
	p.pos++
}

func (p *parser) peek() byte {
	if p.pos >= len(p.input) {
		return 0
	}
	return p.input[p.pos]
}

func (p *parser) parse() string {
	p.nextToken = p.parseToken()
	for p.nextToken != "" {
		p.output += p.nextToken
		p.nextToken = p.parseToken()
	}
	return p.output
}

func (p *parser) parseToken() string {
	var currToken string
	for p.curr == ' ' {
		p.next()
	}

	if p.curr == 0 {
		return ""
	}

	return currToken
}
