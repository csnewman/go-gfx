package main

import (
	"fmt"
	"log/slog"
	"maps"
	"strconv"
)

func (p *RegistryParser) parseFeature(start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	delete(attrs, "comment")
	delete(attrs, "name")
	delete(attrs, "depends")

	numberStr, ok := take(attrs, "number")
	if !ok {
		return fmt.Errorf("missing number")
	}

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	feat := &Feature{
		Version:  numberStr,
		Types:    make(map[string]struct{}),
		Commands: make(map[string]struct{}),
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
		case "require":
			if err := p.parseRequire(feat, tok); err != nil {
				return err
			}
		case "deprecate":
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected type %q", tok.Name)
		}
	}

	p.reg.Features = append(p.reg.Features, feat)

	return nil
}

func (p *RegistryParser) parseRequire(feat *Feature, start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	delete(attrs, "comment")
	delete(attrs, "depends")

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
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
		case "feature", "comment":
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		case "type":
			if err := p.parseRequireType(feat, tok); err != nil {
				return err
			}
		case "command":
			if err := p.parseRequireCommand(feat, tok); err != nil {
				return err
			}
		case "enum":
			if err := p.parseRequireEnum(feat, tok); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected type %q", tok.Name)
		}
	}

	return nil
}

func (p *RegistryParser) parseRequireType(feat *Feature, start Token) error {
	attrs := maps.Clone(start.Attrs)

	name, ok := take(attrs, "name")
	if !ok {
		return fmt.Errorf("%w: name missing", ErrTokenMismatch)
	}

	delete(attrs, "comment")

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	if _, err := p.d.ExpectEnd(start.Name, false); err != nil {
		return fmt.Errorf("end missing: %v", err)
	}

	feat.Types[name] = struct{}{}

	return nil
}

func (p *RegistryParser) parseRequireCommand(feat *Feature, start Token) error {
	attrs := maps.Clone(start.Attrs)

	delete(attrs, "comment")

	name, ok := take(attrs, "name")
	if !ok {
		return fmt.Errorf("%w: name missing", ErrTokenMismatch)
	}

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	if _, err := p.d.ExpectEnd(start.Name, false); err != nil {
		return fmt.Errorf("end missing: %v", err)
	}

	feat.Commands[name] = struct{}{}

	return nil
}

func (p *RegistryParser) parseRequireEnum(feat *Feature, start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	delete(attrs, "comment")
	delete(attrs, "deprecated")

	name, ok := take(attrs, "name")
	if !ok {
		return fmt.Errorf("%w: name missing", ErrTokenMismatch)
	}

	extends, isExtends := take(attrs, "extends")

	protect, _ := take(attrs, "protect")

	var (
		aliasStr string

		valueStr string

		offsetStr string
		extStr    string
		negDir    bool

		bitStr string

		empty bool
	)

	//if isExtends {
	if aliasStr, ok = take(attrs, "alias"); ok {
		delete(attrs, "deprecated")
	} else if valueStr, ok = take(attrs, "value"); ok {
	} else if offsetStr, ok = take(attrs, "offset"); ok {
		extStr, ok = take(attrs, "extnumber")
		if !ok {
			if feat.ExtNumber != "" {
				extStr = feat.ExtNumber
			} else {
				return fmt.Errorf("%w: extnumber missing", ErrTokenMismatch)
			}
		}

		dirStr, ok := take(attrs, "dir")

		if ok {
			if dirStr == "-" {
				negDir = true
			} else {
				panic(dirStr)
			}
		}
	} else if bitStr, ok = take(attrs, "bitpos"); ok {
	} else if !isExtends {
		empty = true
	} else {
		return fmt.Errorf("%w: offset missing", ErrTokenMismatch)
	}

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	if _, err := p.d.ExpectEnd(start.Name, false); err != nil {
		return fmt.Errorf("end missing: %v", err)
	}

	if isExtends {
		ext := FeatureEnumExtension{
			Name:    name,
			Extends: extends,
			Protect: protect,
		}

		if aliasStr != "" {
			ext.Type = "alias"
			ext.Alias = aliasStr
		} else if valueStr != "" {
			ext.Type = "value"
			ext.Value = valueStr
		} else if offsetStr != "" {
			offset, err := strconv.Atoi(offsetStr)
			if err != nil {
				return fmt.Errorf("%w: offset invalid: %v", ErrTokenMismatch, err)
			}

			extID, err := strconv.Atoi(extStr)
			if err != nil {
				return fmt.Errorf("%w: ext invalid: %v", ErrTokenMismatch, err)
			}

			ext.Type = "offset"
			ext.Offset = offset
			ext.Ext = extID
			ext.NegDir = negDir
		} else {
			bit, err := strconv.Atoi(bitStr)
			if err != nil {
				return fmt.Errorf("%w: ext invalid: %v", ErrTokenMismatch, err)
			}

			ext.Type = "bitpos"
			ext.Bitpos = bit
		}

		feat.EnumExtensions = append(feat.EnumExtensions, ext)
	} else if valueStr != "" {
		feat.Constants = append(feat.Constants, FeatureConstant{
			Name:    name,
			Value:   valueStr,
			Protect: protect,
		})
	} else if aliasStr != "" {
		feat.Constants = append(feat.Constants, FeatureConstant{
			Name:    name,
			Alias:   aliasStr,
			Protect: protect,
		})
	} else if empty {
		// drop
	} else {
		panic("unexpected version")
	}

	return nil
}
