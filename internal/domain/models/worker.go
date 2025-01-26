package models

type Worker struct {
	id       string
	name     string
	password string
	login    string
}

func (w Worker) GetId() string {
	return w.id
}

func (w Worker) GetName() string {
	return w.name
}

func (w Worker) GetPassword() string {
	return w.password
}

func (w Worker) GetLogin() string {
	return w.login
}

func (w Worker) SetName(name string) {
	w.name = name
}

func (w Worker) SetPassword(password string) {
	w.password = password
}

func (w Worker) SetLogin(login string) {
	w.login = login
}
