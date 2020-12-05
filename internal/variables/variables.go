package variables

// VariableType : Enum for variable types
type VariableType int

//
const (
	ENVIRONMENT VariableType = iota
	PROTECTEDSSHKEY
	UNPROTECTEDSSHKEY
)

func (v VariableType) String() string {
	return [...]string{"ENVIRONMENT", "PROTECTEDSSHKEY", "UNPROTECTEDSSHKEY"}[v]
}

// Variable : interface for different variables
type Variable interface {
	AddVariable()
	EditVariable()
	DeleteVariable()
	GetType() VariableType
	GetName() string
}

// EnvironmentVariable : variable which will be injected as environment variable
type EnvironmentVariable struct {
	Type  VariableType
	Name  string
	Value string
}

func (env EnvironmentVariable) AddVariable() {

}

func (env EnvironmentVariable) EditVariable() {

}

func (env EnvironmentVariable) DeleteVariable() {

}

func (env EnvironmentVariable) GetType() VariableType {
	return env.Type
}

func (env EnvironmentVariable) GetName() string {
	return env.Name
}

// ProtectedSSHKey : Password protected SSH key in PEM format which will be injected into ssh-agent
type ProtectedSSHKey struct {
	Type     VariableType
	key      string
	password string
}

func (env ProtectedSSHKey) AddVariable() {

}

func (env ProtectedSSHKey) EditVariable() {

}

func (env ProtectedSSHKey) DeleteVariable() {

}

// UnprotectedSSHKey : Unprotected SSH key in PEM format which will be injected into ssh-agent
type UnprotectedSSHKey struct {
	Type VariableType
	key  string
}

func (env UnprotectedSSHKey) AddVariable() {

}

func (env UnprotectedSSHKey) EditVariable() {

}

func (env UnprotectedSSHKey) DeleteVariable() {

}
