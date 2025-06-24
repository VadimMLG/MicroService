package models

type Profile struct {
    ID          uint   `json:"id"`
    Login       string `json:"login"`
    Password    string `json:"-"` // Не сериализуется в JSON
    Phone       string `json:"phone"`
    Email       string `json:"email"`
    Name        string `json:"name"`
    Surname     string `json:"surname"`
    Description string `json:"description"`
    Photo       string `json:"photo"`
    Birthday    string `json:"birthday"` // Можно использовать time.Time
    City        string `json:"city"`
    Workplace   string `json:"workplace"`
}
