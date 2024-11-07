-- dictionary
CREATE TABLE dict (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- parent id
    pid INTEGER,
    word TEXT NOT NULL,
    -- deffinaion
    def TEXT NOT NULL,
    is_root BOOLEAN NOT NULL
);
