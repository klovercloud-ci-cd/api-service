package enums

// STEP_TYPE steps type
type STEP_TYPE string

const (
	// BUILD build step
	BUILD = STEP_TYPE("BUILD")
	// DEPLOY deploy step
	DEPLOY = STEP_TYPE("DEPLOY")
)

// TRIGGER pipeline trigger options
type TRIGGER string

const (
	// AUTO pipeline trigger options is auto
	AUTO = TRIGGER("AUTO")
	// MANUAL pipeline trigger options is MANUAL
	MANUAL = TRIGGER("MANUAL")
)
// PERMISSIONS permission type
type PERMISSIONS string

const (
	// CREATE permission
	CREATE = PERMISSIONS("CREATE")
	// READ permission
	READ = PERMISSIONS("READ")
	// DELETE permission
	DELETE = PERMISSIONS("DELETE")
	// UPDATE permission
	UPDATE = PERMISSIONS("UPDATE")
)
