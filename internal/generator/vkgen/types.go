package main

import (
	"fmt"
	"log/slog"
	"maps"
	"strings"
)

func (p *RegistryParser) parseTypes() error {
	for {
		tok, err := p.d.Next(true)
		if err != nil {
			return err
		}

		if tok.Type == TokenTypeEnd {
			break
		}

		if tok.Name == "comment" {
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}

			continue
		}

		if tok.Name != "type" {
			panic("unexpected token " + tok.Name)
		}

		slog.Info("next", "tok", tok)

		cat := tok.Attrs["category"]

		switch cat {
		case "", "include", "define", "basetype", "handle", "enum":
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		case "bitmask":
			if err := p.parseBitmaskType(tok); err != nil {
				return err
			}
		case "funcpointer":
			// TODO
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		case "struct":
			// TODO
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		case "union":
			// TODO
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected category %q", cat)
		}
	}

	return nil
}

func take[K comparable, V any](m map[K]V, k K) (V, bool) {
	v, ok := m[k]

	delete(m, k)

	return v, ok
}

func (p *RegistryParser) parseBitmaskType(start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	delete(attrs, "category")

	if alias, ok := start.Attrs["alias"]; ok {

		name, ok := start.Attrs["alias"]
		if !ok {
			panic("missing name")
		}

		slog.Info("bitmask type", "name", name, "alias", alias)

		if _, err := p.d.ExpectEnd("type", false); err != nil {
			return fmt.Errorf("end missing: %v", err)
		}

		return nil
	}

	requires, _ := take(attrs, "requires")
	bitValues, _ := take(attrs, "bitvalues")

	_ = bitValues

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	if _, err := p.d.ExpectChar("typedef "); err != nil {
		return fmt.Errorf("typedef missing: %v", err)
	}

	_, typeName, err := p.d.ExpectSimpleText("type", false)
	if err != nil {
		return fmt.Errorf("type missing: %v", err)
	}

	if tk, err := p.d.ExpectAnyChar(); err != nil {
		return fmt.Errorf("space missing: %v", err)
	} else if strings.TrimSpace(tk.Value) != "" {
		return fmt.Errorf("%w: space missing: %v", ErrTokenMismatch, tk.Value)
	}

	_, nameVal, err := p.d.ExpectSimpleText("name", false)
	if err != nil {
		return fmt.Errorf("name missing: %v", err)
	}

	slog.Info("bitmask type", "type", typeName, "name", nameVal, "requires", requires)

	if _, err := p.d.ExpectChar(";"); err != nil {
		return fmt.Errorf("semicolon missing: %v", err)
	}

	if _, err := p.d.ExpectEnd("type", false); err != nil {
		return fmt.Errorf("end missing: %v", err)
	}

	return nil
}
