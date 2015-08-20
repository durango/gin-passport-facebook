package GinPassportFacebook

type Profile struct {
    Id            string `json:"id"`
    Email         string `json:"email"`
    FirstName    string `json:"first_name"`
    LastName     string `json:"last_name"`
    Hd            string `json:"hd"`
    Locale        string `json:"locale"`
    Name          string `json:"name"`
    Picture struct {
        Data struct {
            Url string `json:"url"`
        } `json:"data"`
    } `json:"picture"`
}
