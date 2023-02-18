package main

type Cat struct {
	tableName struct{} `pg:"cats"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	IsSrtipe  bool     `json:"is_srtipe" pg:"is_srtipe"`
	Color     string   `json:"color"  pg:"color"`
}

// FindAllCats Получить список котиков.
func FindAllCats() []Cat {
	var cats []Cat
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&cats).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cats
}

// CreateCat Создать котика.
func CreateCat(cat Cat) Cat {
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&cat).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cat
}

// FindCatById Получить котика по id.
func FindCatById(id string) Cat {
	var cat Cat
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&cat).
		Where("id = ?", id).
		First()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cat
}

// DeleteCatById Удалить котика по id.
func DeleteCatById(id string) Cat {
	var cat Cat
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&cat).
		Where("id = ?", id).
		Delete()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cat
}

func UpdateCat(cat Cat) Cat {
	pgConnect := PostgresConnect()

	oldCat := FindCatById(cat.ID)

	oldCat.Name = cat.Name
	oldCat.IsSrtipe = cat.IsSrtipe
	oldCat.Color = cat.Color

	_, err := pgConnect.Model(&oldCat).
		Set("name = ?", oldCat.Name).
		Set("is_srtipe = ?", oldCat.IsSrtipe).
		Set("color = ?", oldCat.Color).
		Where("id = ?", oldCat.ID).
		Update()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return oldCat
}
