package data

type User struct {
	Name       string `json:"name"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
}

type Source struct {
	SourceId    string `json:"sourceId"`
	Name        string `json:"sourceName"`
	owner       string `json:"owner"`
	description string `json:"description"`
}

type Variable struct {
	VariableId  string `json:"variableId"`
	Name        string `json:"variableName"`
	SourceId    string `json:"sourceId"`
	description string `json:"description"`
	unit        string `json:"unit"`
}

type UserGroup struct {
	Name string `json:"userGroupName"`
	Owner string `json:"userGroupOwner"`
	description string `json:"userGroupDesc"`
}

type VariableGroup struct {
	VarGroupId string `json:"varGroupId"`
	Name string `json:"varGroupName"`
	Owner string `json:"varGroupOwner"`
	description string `json:"varGroupDesc"`
}

func (user *User) String() string {
	return user.Name + "[" + user.Email + "]"
}

func (source *Source) String() string {
	return source.Name + "[" + source.SourceId + "]"
}

func (variable *Variable) String() string {
	return variable.Name + "[" + variable.SourceId + " : " + variable.VariableId + "]"
}
