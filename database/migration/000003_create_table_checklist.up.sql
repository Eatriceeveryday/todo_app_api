CREATE TABLE checklist(
    check_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    detail TEXT NOT NULL,
    is_completed BOOLEAN,
    todo_id uuid NOT NULL ,
    FOREIGN KEY (todo_id) REFERENCES todo(todo_id)
);