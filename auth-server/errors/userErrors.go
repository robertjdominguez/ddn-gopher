package customErrors

import "errors"

// TODO: Move this to an errors package...Rikin was right
var ErrEmptyUsername = errors.New("username cannot be an empty string")
