package data

// StudentAge public variable, array is static data type
var StudentAge [10]int

// StudentName public variable, slice is dynamic data type
var StudentName []string

// StudentStatus public variable, map
var StudentStatus map[string]bool

// WellKnownPorts : well known ports, map is dynamic data type
var WellKnownPorts map[string]int = map[string]int{"http": 80, "https": 443}
