CREATE TABLE IF NOT EXISTS transactions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	from_account INTEGER,
	to_account INTEGER,
	amount INTEGER,
	status INTEGER,
	flagged BOOLEAN,
	created_at DATETIME,
	updated_at DATETIME
);
