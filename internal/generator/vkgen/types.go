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
		case "", "include", "define", "basetype", "handle":
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		case "enum":
			if err := p.parseEnumType(tok); err != nil {
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
			if err := p.parseStructType(tok); err != nil {
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

func (p *RegistryParser) parseEnumType(start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	delete(attrs, "category")

	name, ok := take(attrs, "name")
	if !ok {
		return fmt.Errorf("%w: name missing", ErrTokenMismatch)
	}

	if alias, ok := start.Attrs["alias"]; ok {
		slog.Info("enum type", "alias", alias)

		if _, ok := p.reg.Aliases[name]; ok {
			return fmt.Errorf("%w: alias %q already defined", ErrTokenMismatch, name)
		}

		p.reg.Aliases[name] = alias

		if _, err := p.d.ExpectEnd("type", false); err != nil {
			return fmt.Errorf("end missing: %v", err)
		}

		return nil
	}

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	if _, ok := p.reg.Types[name]; ok {
		return fmt.Errorf("%w: type %q already defined", ErrTokenMismatch, name)
	}

	p.reg.Types[name] = &Type{
		Name:     name,
		Category: CategoryEnum,
	}

	if _, err := p.d.ExpectEnd("type", false); err != nil {
		return fmt.Errorf("end missing: %v", err)
	}

	return nil
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
		name, ok := start.Attrs["name"]
		if !ok {
			panic("missing name")
		}

		slog.Info("bitmask type", "name", name, "alias", alias)

		if _, ok := p.reg.Aliases[name]; ok {
			return fmt.Errorf("%w: alias %q already defined", ErrTokenMismatch, name)
		}

		p.reg.Aliases[name] = alias

		if _, err := p.d.ExpectEnd("type", false); err != nil {
			return fmt.Errorf("end missing: %v", err)
		}

		return nil
	}

	requires, _ := take(attrs, "requires")
	bitValues, _ := take(attrs, "bitvalues")

	if requires == "" && bitValues != "" {
		requires = bitValues
	} else if requires != "" && bitValues != "" {
		return fmt.Errorf("%w: reqs and values %q or %q", ErrTokenMismatch, requires, bitValues)
	}

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

	if _, ok := p.reg.Types[nameVal]; ok {
		return fmt.Errorf("%w: type %q already defined", ErrTokenMismatch, nameVal)
	}

	parsed := &Type{
		Name:     nameVal,
		Category: CategoryBitmask,
		Requires: requires,
	}

	p.reg.Types[nameVal] = parsed

	switch typeName {
	case "VkFlags":
		parsed.BitmaskWidth = 32
	case "VkFlags64":
		parsed.BitmaskWidth = 64
	default:
		panic("unhandled type: " + typeName)
	}

	return nil
}
