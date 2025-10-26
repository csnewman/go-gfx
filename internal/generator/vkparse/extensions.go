package vkparse

import (
	"fmt"
	"log/slog"
	"maps"
	"strings"
)

func (p *RegistryParser) parseExtensions(start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	delete(attrs, "comment")

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
		case "extension":
			if err := p.parseExtension(tok); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected type %q", tok.Name)
		}
	}

	return nil
}

func (p *RegistryParser) parseExtension(start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnoreSupported(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	//if shouldIgnoreRatified(attrs) {
	//	if _, err := p.d.FindEnd(start.Name); err != nil {
	//		return err
	//	}
	//
	//	return nil
	//}

	delete(attrs, "ratified")

	delete(attrs, "author")
	delete(attrs, "comment")
	delete(attrs, "contact")
	delete(attrs, "nofeatures")
	delete(attrs, "type")
	delete(attrs, "deprecatedby")
	delete(attrs, "specialuse")
	delete(attrs, "promotedto")
	delete(attrs, "obsoletedby")
	delete(attrs, "provisional")
	delete(attrs, "sortorder")

	platform, ok := take(attrs, "platform")
	if ok && platform != "" && platform != "metal" && platform != "win32" {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	name, ok := take(attrs, "name")
	if !ok {
		return fmt.Errorf("missing name")
	}

	numberStr, ok := take(attrs, "number")
	if !ok {
		return fmt.Errorf("missing number")
	}

	depStr, _ := take(attrs, "depends")

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	feat := &Feature{
		Types:    make(map[string]struct{}),
		Commands: make(map[string]struct{}),

		Name:      name,
		ExtNumber: numberStr,
	}

	if depStr != "" {
		feat.ExtDepends = strings.Split(depStr, ",")
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
		case "deprecate", "remove":
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected type %q", tok.Name)
		}
	}

	p.reg.Extensions = append(p.reg.Extensions, feat)

	return nil
}
