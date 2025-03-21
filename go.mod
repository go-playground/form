module github.com/go-playground/form/v4

go 1.13

retract (
	v4.2.2 // Accidentally publiches breaking change
	v4.2.3 // For retractions and revert of that code.
)

require github.com/go-playground/assert/v2 v2.0.1
