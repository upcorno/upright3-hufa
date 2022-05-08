package model

func (user *User) Insert() (err error) {
	_, err = Db.InsertOne(user)
	return
}

func (user *User) Get() (has bool, err error) {
	return Db.Get(user)
}

func (user *User) Update() (err error) {
	_, err = Db.Update(user, User{Id: user.Id})
	return
}
