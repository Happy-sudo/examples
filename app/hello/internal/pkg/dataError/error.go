package dataError

import (
	"github.com/pkg/errors"
)

func EntError(err error) error {
	switch {
	case ent.IsNotFound(err):
		return errors.Wrap(err, "todo item was not found")
	case ent.IsNotSingular(err):
		return errors.Wrap(err, "query multiple items")
	case ent.IsConstraintError(err):
		return errors.Wrap(err, "constraint failure")
	case ent.IsValidationError(err):
		return errors.Wrap(err, "validation failure")
	case ent.IsNotLoaded(err):
		return errors.Wrap(err, "loaded failure")
	case err != nil:
		return errors.Wrap(err, "deletion error:")
	}
	return nil
}
