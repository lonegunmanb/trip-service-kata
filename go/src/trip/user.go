package trip

type User struct {
	trips   []*Trip
	friends []*User
}

func (u *User) GetTrips() []*Trip {
	return u.trips
}

func (u *User) AddTrip(trip *Trip) {
	u.trips = append(u.trips, trip)
}

func (u *User) GetFriends() []*User {
	return u.friends
}

func (u *User) AddFriends(friend *User) {
	u.friends = append(u.friends, friend)
}
