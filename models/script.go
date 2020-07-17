package models

type Script struct {
	Sql     string   `json:"sql"`
	Session *Session `json:"-" gorm:"-"`
}

func (s *Script) Execute(script *Script) error {
	tx := GetSessionTx(s.Session)
	_, err := tx.Raw(script.Sql).Rows()
	return err

}
func NewScriptModel(session *Session) *Script {
	return &Script{Session: session}
}
