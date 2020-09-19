# Units
[![Build](https://github.com/marstr/units/workflows/Build/badge.svg)](https://github.com/marstr/units/actions?query=workflow%3ABuild) [![PkgGoDev](https://pkg.go.dev/badge/github.com/marstr/units)](https://pkg.go.dev/github.com/marstr/units)

Find yourself wishing that you had type defintions for all of the units in your life, comparable to `time.Time`? This
library hopes to bring that level of convenience and type safety while working with all sorts of different quantities.
All of these types are defined to be 64 bits wide for maximum precision and range. Types which loan themselves to being
discrete are expressed as unsigned integers, types that are more prone to being continuous are of type `float64`.

# License
This library is available for use under the very permissive MIT license. The full text of which can be found here: 
[./LICENSE](./LICENSE)