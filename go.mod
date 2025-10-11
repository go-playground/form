module github.com/go-playground/form/v4

go 1.21

retract (
	v4.2.3 // For retractions and revert of that code.
	v4.2.2 // Accidentally publiches breaking change
)

require github.com/go-playground/assert/v2 v2.2.0
