package main

import (
	"fmt"
	"log/slog"
	"maps"
	"strconv"
)

func (p *RegistryParser) parseEnums(start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	enumType, ok := take(attrs, "type")
	if !ok {
		enumType = "enum"
	}

	if enumType == "constants" {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	name, ok := take(attrs, "name")
	if !ok {
		return fmt.Errorf("%w: name missing", ErrTokenMismatch)
	}

	comment, _ := take(attrs, "comment")

	bitwidth, _ := take(attrs, "bitwidth")

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	if _, ok := p.reg.Enums[name]; ok {
		return fmt.Errorf("%w: enum %v already defined", ErrTokenMismatch, name)
	}

	enum := &Enum{
		Name:    name,
		Type:    enumType,
		Comment: comment,
	}

	p.reg.Enums[name] = enum

	if bitwidth != "" {
		v, err := strconv.Atoi(bitwidth)
		if err != nil {
			return fmt.Errorf("%w: bitwidth invalid: %v", ErrTokenMismatch, err)
		}

		enum.BitWidth = v
	}

	for {
		tok, err := p.d.Next(true)
		if err != nil {
			return err
		}

		if tok.Type == TokenTypeEnd {
			break
		}

		slog.Info("next", "tok", tok)

		switch tok.Name {
		case "comment", "unused":
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		case "enum":
			if err := p.parseEnumEntry(enum, tok); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected type %q", tok.Name)
		}
	}

	return nil
}

func (p *RegistryParser) parseEnumEntry(enum *Enum, start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	name, ok := take(attrs, "name")
	if !ok {
		return fmt.Errorf("%w: name missing", ErrTokenMismatch)
	}

	comment, _ := take(attrs, "comment")

	valueStr, _ := take(attrs, "value")
	alias, _ := take(attrs, "alias")
	deprecated, _ := take(attrs, "deprecated")

	bitposStr, _ := take(attrs, "bitpos")

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	value := &EnumValue{
		Name:       name,
		Value:      valueStr,
		Alias:      alias,
		Deprecated: deprecated,
		Comment:    comment,
	}

	if bitposStr != "" {
		v, err := strconv.Atoi(bitposStr)
		if err != nil {
			return fmt.Errorf("%w: bitpos invalid: %v", ErrTokenMismatch, err)
		}

		value.BitPos = v
	}

	enum.Values = append(enum.Values, value)

	if _, err := p.d.ExpectEnd("enum", false); err != nil {
		return fmt.Errorf("end missing: %v", err)
	}

	return nil
}
