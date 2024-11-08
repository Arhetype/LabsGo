package Task3

func removePerson(people map[string]int, name string) {
	delete(people, name)
}
