CREATE TABLE todo (
    todo_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    user_id uuid NOT NULL ,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);