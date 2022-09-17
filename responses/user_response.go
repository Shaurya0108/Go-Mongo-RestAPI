package responses

type UserResponse struct {
    Id  int                    `json:"id"`
    Name string                 `json:"name"`
    Data    map[string]interface{} `json:"data"`
}