package model

func (user *User) Insert() (err error) {
	uid, err := Db.InsertOne(user)
	user.Id = int(uid)
	return
}

func (user *User) Get() (has bool, err error) {
	return Db.Get(user)
}

func (user *User) Update() (err error) {
	_, err = Db.Update(user, User{Id: user.Id})
	return
}
