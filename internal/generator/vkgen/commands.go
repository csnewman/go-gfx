package main

import (
	"fmt"
	"log/slog"
	"maps"
	"strings"
)

func (p *RegistryParser) parseCommands(start Token) error {
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
		case "command":
			if err := p.parseCommand(tok); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected type %q", tok.Name)
		}
	}

	return nil
}

func (p *RegistryParser) parseCommand(start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) || shouldIgnoreExport(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}

	if alias, ok := start.Attrs["alias"]; ok {
		name, ok := take(attrs, "name")
		if !ok {
			return fmt.Errorf("%w: name missing", ErrTokenMismatch)
		}

		slog.Info("command type", "alias", alias, "name", name)

		if _, ok := p.reg.Aliases[name]; ok {
			return fmt.Errorf("%w: alias %q already defined", ErrTokenMismatch, name)
		}

		p.reg.Aliases[name] = alias

		if _, err := p.d.ExpectEnd(start.Name, false); err != nil {
			return fmt.Errorf("end missing: %v", err)
		}

		return nil
	}

	delete(attrs, "successcodes")
	delete(attrs, "errorcodes")
	delete(attrs, "queues")
	delete(attrs, "allownoqueues")
	delete(attrs, "cmdbufferlevel")
	delete(attrs, "conditionalrendering")
	delete(attrs, "renderpass")
	delete(attrs, "tasks")
	delete(attrs, "comment")
	delete(attrs, "videocoding")

	if len(attrs) > 0 {
		return fmt.Errorf("%w: unprocessed attributes: %v", ErrTokenMismatch, attrs)
	}

	if _, err := p.d.ExpectStart("proto", true); err != nil {
		return fmt.Errorf("%w: %v", ErrTokenMismatch, err)
	}

	_, tyName, err := p.d.ExpectSimpleText("type", false)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrTokenMismatch, err)
	}

	slog.Info("proto", "type", tyName)

	if middleTok, err := p.d.ExpectAnyChar(); err != nil {
		return fmt.Errorf("expecting middle token: %w", err)
	} else if strings.TrimSpace(middleTok.Value) != "" {
		return fmt.Errorf("expecting middle token: %w %q", err, middleTok.Value)
	}

	_, commandName, err := p.d.ExpectSimpleText("name", false)
	if err != nil {
		return fmt.Errorf("expecting name: %w", err)
	}

	if _, err := p.d.ExpectEnd("proto", false); err != nil {
		return fmt.Errorf("%w: %v", ErrTokenMismatch, err)
	}

	slog.Info("proto", "command", commandName)

	if _, ok := p.reg.Commands[commandName]; ok {
		return fmt.Errorf("%w: command %v already defined", ErrTokenMismatch, commandName)
	}

	cmd := &Command{
		Name:       commandName,
		ReturnType: tyName,
	}

	p.reg.Commands[commandName] = cmd

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
		case "comment", "implicitexternsyncparams":
			if _, err := p.d.FindEnd(tok.Name); err != nil {
				return err
			}
		case "param":
			if err := p.parseCommandParam(cmd, tok); err != nil {
				return err
			}

			//if _, err := p.d.FindEnd(tok.Name); err != nil {
			//	return err
			//}
		//case "enum":
		//	if err := p.parseEnumEntry(enum, tok); err != nil {
		//		return err
		//	}
		default:
			return fmt.Errorf("unexpected type %q", tok.Name)
		}
	}

	return nil
}

func (p *RegistryParser) parseCommandParam(cmd *Command, start Token) error {
	attrs := maps.Clone(start.Attrs)

	if shouldIgnore(attrs) || shouldIgnoreExport(attrs) {
		if _, err := p.d.FindEnd(start.Name); err != nil {
			return err
		}

		return nil
	}
	delete(attrs, "optional")
	delete(attrs, "externsync")
	delete(attrs, "len")
	delete(attrs, "altlen")
	delete(attrs, "validstructs")
	delete(attrs, "noautovalidity")
	delete(attrs, "stride")
	delete(attrs, "objecttype")

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

	var memberClass MemberCategory

	switch middleTok.Value {
	case "":
		memberClass = MemberCategoryDirect
	case "*":
		memberClass = MemberCategoryPointer
	case "**", "*const*":
		memberClass = MemberCategoryPointer2
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

			cmd.Members = append(cmd.Members, CommandMember{
				Name:     fieldName,
				Category: MemberCategoryUnsupported,
			})

			return nil
		} else if strings.HasPrefix(tok.Value, ":") {
			// TODO: support
			if _, err := p.d.FindEnd(start.Name); err != nil {
				return err
			}

			cmd.Members = append(cmd.Members, CommandMember{
				Name:     fieldName,
				Category: MemberCategoryUnsupported,
			})

			return nil
		}

		panic("todo post text " + tok.Value)
	}

	slog.Info("member", "type", tyName, "name", fieldName, "class", memberClass)

	if _, err := p.d.ExpectEnd(start.Name, false); err != nil {
		return fmt.Errorf("end missing: %v", err)
	}

	cmd.Members = append(cmd.Members, CommandMember{
		Name:     fieldName,
		Type:     tyName,
		Category: memberClass,
	})

	return nil
}
