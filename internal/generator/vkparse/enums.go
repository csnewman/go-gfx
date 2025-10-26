package vkparse

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
		return p.parseConstants()
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
		value.HasBitPos = true
	}

	enum.Values = append(enum.Values, value)

	if _, err := p.d.ExpectEnd("enum", false); err != nil {
		return fmt.Errorf("end missing: %v", err)
	}

	return nil
}

func applyEnumExtension(extension FeatureEnumExtension, reg *Registry) {
	target, ok := reg.Enums[extension.Extends]
	if !ok {
		panic(fmt.Sprintf("enum %v not defined", extension.Extends))
	}

	for _, existing := range target.Values {
		if existing.Name == extension.Name {
			return
		}
	}

	switch extension.Type {
	case "offset":
		val := 1000000000 + ((extension.Ext - 1) * 1000) + extension.Offset

		if extension.NegDir {
			val = -val
		}

		target.Values = append(target.Values, &EnumValue{
			Name:  extension.Name,
			Value: strconv.Itoa(val),
		})
	case "bitpos":
		target.Values = append(target.Values, &EnumValue{
			Name:      extension.Name,
			BitPos:    extension.Bitpos,
			HasBitPos: true,
		})
	case "alias":
		target.Values = append(target.Values, &EnumValue{
			Name:  extension.Name,
			Alias: extension.Alias,
		})
	case "value":
		target.Values = append(target.Values, &EnumValue{
			Name:  extension.Name,
			Value: extension.Value,
		})
	default:
		panic(fmt.Sprintf("unknown type %v", extension.Type))
	}
}

func (p *RegistryParser) parseConstants() error {

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
		case "enum":
			if err := p.parseConstantEntry(tok); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected type %q", tok.Name)
		}
	}

	return nil
}

func (p *RegistryParser) parseConstantEntry(start Token) error {
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

	constType, ok := take(attrs, "type")
	if !ok {
		return fmt.Errorf("%w: type missing", ErrTokenMismatch)
	}

	comment, _ := take(attrs, "comment")

	valueStr, ok := take(attrs, "value")
	if !ok {
		return fmt.Errorf("%w: value missing", ErrTokenMismatch)
	}

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	if _, ok := p.reg.Constants[name]; ok {
		return fmt.Errorf("%w: constant %v already defined", ErrTokenMismatch, name)
	}

	p.reg.Constants[name] = &EnumValue{
		Name:      name,
		Value:     valueStr,
		Comment:   comment,
		ConstType: constType,
	}

	if _, err := p.d.ExpectEnd("enum", false); err != nil {
		return fmt.Errorf("end missing: %v", err)
	}

	return nil
}
