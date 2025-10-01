package main

import (
	"fmt"
	"log/slog"
	"maps"
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

	p.reg.Types[name] = &Type{
		Name:     name,
		Category: CategoryStruct,

		StructReadOnly: readOnly,
	}

	//if _, err := p.d.ExpectEnd("type", false); err != nil {
	//	return fmt.Errorf("end missing: %v", err)
	//}

	if _, err := p.d.FindEnd(start.Name); err != nil {
		return err
	}

	return nil
}
