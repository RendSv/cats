package main

type Dog struct {
	tableName struct{} `pg:"dogs"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	Color     string   `json:"color"  pg:"color"`
	Breed     string   `json:"breed"  pg:"breed"`
}

// FindAllDogs Получить список котиков.
func FindAllDogs() []Dog {
	var dogs []Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dogs).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dogs
}

// CreateDog Создать котика.
func CreateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&dog).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

// FindDogById Получить котика по id.
func FindDogById(id string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dog).
		Where("id = ?", id).
		First()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

// DeleteDogById Удалить котика по id.
func DeleteDogById(id string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&dog).
		Where("id = ?", id).
		Delete()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

func UpdateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	oldDog := FindDogById(dog.ID)

	oldDog.Name = dog.Name

	oldDog.Color = dog.Color

	_, err := pgConnect.Model(&oldDog).
		Set("name = ?", oldDog.Name).
		Set("color = ?", oldDog.Color).
		Where("id = ?", oldDog.ID).
		Update()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return oldDog
}
