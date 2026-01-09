package models

type Task struct {
    ID        int64  `json:"id"`
    Title     string `json:"title"`
    Day       string `json:"day"` // YYYY-MM-DD
    IsDone    bool   `json:"isDone"`
    CreatedAt string `json:"createdAt"`
}
