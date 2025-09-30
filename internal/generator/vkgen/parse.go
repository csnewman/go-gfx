package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

type RegistryParser struct {
	d   *Decoder
	reg *Registry
}

func parse(path string) (*Registry, error) {

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	rd := xml.NewDecoder(f)

	d := &Decoder{
		d: rd,
	}

	reg := &Registry{
		Types:   make(map[string]*Type),
		Aliases: make(map[string]string),
		Enums:   make(map[string]*Enum),
	}

	rg := &RegistryParser{
		d:   d,
		reg: reg,
	}

	if _, err := d.ExpectStart("registry", true); err != nil {
		return nil, err
	}

	for {

		tok, err := d.Next(true)
		if err != nil {
			return nil, err
		}

		if tok.Type == TokenTypeEnd {
			break
		}

		slog.Info("next", "tok", tok)

		switch tok.Name {
		case "comment", "platforms", "tags", "spirvextensions", "spirvcapabilities", "sync", "videocodecs":
			if _, err := d.FindEnd(tok.Name); err != nil {
				return nil, err
			}

			continue

		case "types":
			if err := rg.parseTypes(); err != nil {
				return nil, fmt.Errorf("parsing types: %w", err)
			}

			continue

		case "enums":
			if err := rg.parseEnums(tok); err != nil {
				return nil, fmt.Errorf("parsing types: %w", err)
			}

			continue

		case "commands", "feature", "extensions", "formats":
			if _, err := d.FindEnd(tok.Name); err != nil {
				return nil, err
			}

			// TODO: support

			continue

		}

		panic("stop")
	}

	for name, target := range reg.Aliases {
		targetType, ok := reg.Types[target]
		if !ok {
			panic(fmt.Sprintf("unknown type %s for alias %s", target, name))
		}

		targetType.Aliases = append(targetType.Aliases, name)
	}

	return reg, nil
}

type Decoder struct {
	d *xml.Decoder
}

type TokenType string

const (
	TokenTypeStart TokenType = "start"
	TokenTypeEnd   TokenType = "end"
	TokenTypeChar  TokenType = "char"
)

type Token struct {
	Type TokenType

	Name  string
	Attrs map[string]string

	Value string
}

var ErrTokenMismatch = errors.New("token mismatch")

func (d *Decoder) ExpectStart(name string, skipChar bool) (Token, error) {
	return d.Expect(Token{Type: TokenTypeStart, Name: name}, skipChar)
}

func (d *Decoder) ExpectChar(value string) (Token, error) {
	return d.Expect(Token{Type: TokenTypeChar, Value: value}, false)
}

func (d *Decoder) ExpectAnyChar() (Token, error) {
	found, err := d.Next(false)
	if err != nil {
		return Token{}, err
	}

	if found.Type != TokenTypeChar {
		return Token{}, fmt.Errorf("%w: %v", ErrTokenMismatch, found.Type)
	}

	return found, nil
}

func (d *Decoder) ExpectEnd(name string, skipChar bool) (Token, error) {
	return d.Expect(Token{Type: TokenTypeEnd, Name: name}, skipChar)
}

func (d *Decoder) ExpectSimpleText(name string, skipChar bool) (Token, string, error) {
	t, err := d.Expect(Token{Type: TokenTypeStart, Name: name}, skipChar)
	if err != nil {
		return Token{}, "", fmt.Errorf("expecting start token: %w", err)
	}

	valTok, err := d.Next(false)
	if err != nil {
		return Token{}, "", fmt.Errorf("failed to get next: %w", err)
	}

	if valTok.Type != TokenTypeChar {
		return Token{}, "", fmt.Errorf("%w: %v", ErrTokenMismatch, valTok.Type)
	}

	if _, err := d.Expect(Token{Type: TokenTypeEnd, Name: name}, false); err != nil {
		return Token{}, "", fmt.Errorf("expecting end token: %w", err)
	}

	return t, valTok.Value, nil
}

func (d *Decoder) Expect(tk Token, skipChar bool) (Token, error) {
	found, err := d.Next(skipChar)
	if err != nil {
		return Token{}, err
	}

	if found.Type != tk.Type || found.Name != tk.Name || found.Value != tk.Value {
		return Token{}, fmt.Errorf("%w: %v %q %q", ErrTokenMismatch, found.Type, found.Name, tk.Value)
	}

	return tk, nil
}

func (d *Decoder) Next(skipChar bool) (Token, error) {
	for {
		raw, err := d.d.Token()
		if err != nil {
			return Token{}, err
		}

		switch t := raw.(type) {
		case xml.StartElement:
			out := Token{
				Type:  TokenTypeStart,
				Name:  t.Name.Local,
				Attrs: make(map[string]string),
			}

			for _, attr := range t.Attr {
				out.Attrs[attr.Name.Local] = attr.Value
			}

			return out, nil
		case xml.CharData:
			if skipChar {
				continue
			}

			return Token{
				Type:  TokenTypeChar,
				Value: string(t),
			}, nil
		case xml.EndElement:
			return Token{
				Type: TokenTypeEnd,
				Name: t.Name.Local,
			}, nil
		case xml.Comment, xml.ProcInst, xml.Directive:
			continue
		default:
			panic(fmt.Sprintf("unhandled token type %T", raw))
		}
	}
}

func (d *Decoder) FindEnd(name string) (Token, error) {
	depth := 1

	for {
		tok, err := d.Next(true)
		if err != nil {
			return Token{}, err
		}

		switch tok.Type {
		case TokenTypeStart:
			depth++

		case TokenTypeEnd:
			depth--

			if depth == 0 {
				if tok.Name != name {
					return Token{}, fmt.Errorf("%w: end doesn't match %q %q", ErrTokenMismatch, name, tok.Name)
				}

				return tok, nil
			}
		}
	}
}
