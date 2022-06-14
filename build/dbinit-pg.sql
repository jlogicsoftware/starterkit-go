DROP TABLE IF EXISTS todos;

CREATE TABLE todos (
	id SERIAL PRIMARY KEY,
	title varchar(255) NOT NULL,
	description varchar(255) NOT NULL,
	status BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO todos (id, title, description, status, created_at, updated_at) VALUES
(1, 'Test-1', 'Test 1', FALSE, '2020-01-01 00:00:00', '2020-02-01 00:00:00');
INSERT INTO todos (id, title, description, status, created_at, updated_at) VALUES
(2, 'Test-2', 'Test 2', FALSE, '2020-01-01 00:00:00', '2020-02-01 00:00:00');
INSERT INTO todos (id, title, description, status, created_at, updated_at) VALUES
(3, 'Test-3', 'Test 3', TRUE, '2020-01-01 00:00:00', '2020-02-01 00:00:00');
