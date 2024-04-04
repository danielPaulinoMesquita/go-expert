module novo

go 1.21.3

require (
	github.com/google/uuid v1.4.0
	golandProjects/awesomeProject/packaging/math v0.0.0-00010101000000-000000000000
)

/// go.sum looks like package-lock.json, it is responsible to save the version of lib or module downloaded by go.mod

replace golandProjects/awesomeProject/packaging/math => ./../math
