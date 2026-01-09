package handlers

import (
    "database/sql"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "personal-tracker/backend/internal/models"
)

type TaskHandlers struct {
    DB *sql.DB
}

func (h TaskHandlers) Register(r *gin.Engine) {
    api := r.Group("/api")
    api.GET("/tasks", h.listTasks)
    api.POST("/tasks", h.createTask)
    api.PATCH("/tasks/:id/toggle", h.toggleTask)
    api.DELETE("/tasks/:id", h.deleteTask)

    api.GET("/summary/week", h.weekSummary)
}

func (h TaskHandlers) listTasks(c *gin.Context) {
    day := c.Query("day")
    if day == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "day is required (YYYY-MM-DD)"})
        return
    }

    rows, err := h.DB.Query(`
        SELECT id, title, day, is_done, created_at
        FROM tasks
        WHERE day = ?
        ORDER BY created_at DESC, id DESC
    `, day)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    out := make([]models.Task, 0)
    for rows.Next() {
        var t models.Task
        var isDoneInt int
        if err := rows.Scan(&t.ID, &t.Title, &t.Day, &isDoneInt, &t.CreatedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        t.IsDone = isDoneInt == 1
        out = append(out, t)
    }

    c.JSON(http.StatusOK, out)
}

type createTaskReq struct {
    Title string `json:"title"`
    Day   string `json:"day"`
}

func (h TaskHandlers) createTask(c *gin.Context) {
    var req createTaskReq
    if err := c.ShouldBindJSON(&req); err != nil || req.Title == "" || req.Day == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "title and day are required"})
        return
    }

    res, err := h.DB.Exec(`INSERT INTO tasks (title, day, is_done) VALUES (?, ?, 0)`, req.Title, req.Day)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    id, _ := res.LastInsertId()
    c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h TaskHandlers) toggleTask(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    _, err = h.DB.Exec(`
        UPDATE tasks
        SET is_done = CASE WHEN is_done = 1 THEN 0 ELSE 1 END
        WHERE id = ?
    `, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}

func (h TaskHandlers) deleteTask(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    _, err = h.DB.Exec(`DELETE FROM tasks WHERE id = ?`, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}

type WeekSummaryItem struct {
    Day       string `json:"day"`
    Total     int    `json:"total"`
    Completed int    `json:"completed"`
}

func (h TaskHandlers) weekSummary(c *gin.Context) {
    start := c.Query("start") // YYYY-MM-DD
    if start == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "start is required (YYYY-MM-DD)"})
        return
    }

    rows, err := h.DB.Query(`
        SELECT day,
               COUNT(*) as total,
               SUM(CASE WHEN is_done = 1 THEN 1 ELSE 0 END) as completed
        FROM tasks
        WHERE day >= date(?) AND day <= date(?, '+6 day')
        GROUP BY day
        ORDER BY day ASC
    `, start, start)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    out := make([]WeekSummaryItem, 0)
    for rows.Next() {
        var item WeekSummaryItem
        if err := rows.Scan(&item.Day, &item.Total, &item.Completed); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        out = append(out, item)
    }

    c.JSON(http.StatusOK, out)
}
