package vkparse

import (
	"fmt"
	"log/slog"
	"maps"
	"strings"
)

func (p *RegistryParser) parseStructType(start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	delete(attrs, "category")

	// unsupported
	delete(attrs, "requiredlimittype")
	delete(attrs, "structextends")
	delete(attrs, "allowduplicate")
	delete(attrs, "comment")

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

	readOnly := false

	retReadOnly, ok := take(attrs, "returnedonly")
	if ok && retReadOnly == "true" {
		readOnly = true
	}

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	if _, ok := p.reg.Types[name]; ok {
		return fmt.Errorf("%w: type %q already defined", ErrTokenMismatch, name)
	}

	ty := &Type{
		Name:     name,
		Category: CategoryStruct,

		StructReadOnly: readOnly,
	}

	p.reg.Types[name] = ty

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
		case "comment":
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		case "member":
			if err := p.parseStructMember(ty, tok); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected type %q", tok.Name)
		}
	}

	return nil
}

func (p *RegistryParser) parseStructMember(ty *Type, start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	delete(attrs, "optional")
	delete(attrs, "noautovalidity")
	delete(attrs, "limittype")
	delete(attrs, "len")
	delete(attrs, "altlen")
	delete(attrs, "values")
	delete(attrs, "deprecated")
	delete(attrs, "externsync")
	delete(attrs, "objecttype")
	delete(attrs, "featurelink")
	delete(attrs, "selector")

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	if tok, ok, err := p.d.PeekAnyChar(); err != nil {
		return fmt.Errorf("peek any-char: %w", err)
	} else if ok {
		tok.Value = strings.TrimSpace(tok.Value)
		tok.Value = strings.TrimSpace(strings.TrimPrefix(tok.Value, "const"))
		tok.Value = strings.TrimSpace(strings.TrimPrefix(tok.Value, "struct"))

		if tok.Value != "" {
			panic("todo pre text " + tok.Value)
		}
	}

	_, tyName, err := p.d.ExpectSimpleText("type", false)
	if err != nil {
		return fmt.Errorf("expecting type: %w", err)
	}

	middleTok, err := p.d.ExpectAnyChar()
	if err != nil {
		return fmt.Errorf("expecting middle token: %w", err)
	}

	middleTok.Value = strings.TrimSpace(middleTok.Value)
	middleTok.Value = strings.ReplaceAll(middleTok.Value, " ", "")

	var fieldClass FieldCategory

	switch middleTok.Value {
	case "":
		fieldClass = FieldCategoryDirect
	case "*":
		fieldClass = FieldCategoryPointer
	case "*const*":
		fieldClass = FieldCategoryPointer2
	default:
		panic(middleTok.Value)
	}

	_, fieldName, err := p.d.ExpectSimpleText("name", false)
	if err != nil {
		return fmt.Errorf("expecting name: %w", err)
	}

	if tok, ok, err := p.d.PeekAnyChar(); err != nil {
		return fmt.Errorf("peek any-char: %w", err)
	} else if ok {
		if strings.HasPrefix(tok.Value, "[") {
			// TODO: support
			if _, err := p.d.FindEnd(start.Name); err != nil {
				return err
			}

			ty.Fields = append(ty.Fields, StructField{
				Name:     fieldName,
				Category: FieldCategoryUnsupported,
			})

			return nil
		} else if strings.HasPrefix(tok.Value, ":") {
			// TODO: support
			if _, err := p.d.FindEnd(start.Name); err != nil {
				return err
			}

			ty.Fields = append(ty.Fields, StructField{
				Name:     fieldName,
				Category: FieldCategoryUnsupported,
			})

			return nil
		}

		panic("todo post text " + tok.Value)
	}

	slog.Info("member", "type", tyName, "name", fieldName, "class", fieldClass)

	if tok, ok, err := p.d.PeekStart("comment"); err != nil {
		return fmt.Errorf("peek comment: %w", err)
	} else if ok {
		if _, err := p.d.FindEnd(tok.Name); err != nil {
			return err
		}
	}

	if _, err := p.d.ExpectEnd(start.Name, false); err != nil {
		return fmt.Errorf("end missing: %v", err)
	}

	ty.Fields = append(ty.Fields, StructField{
		Name:     fieldName,
		Type:     tyName,
		Category: fieldClass,
	})

	return nil
}
