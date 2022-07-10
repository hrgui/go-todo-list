CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_date = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TABLE IF NOT EXISTS todos (
  id SERIAL NOT NULL PRIMARY KEY, 
  title TEXT NOT NULL, 
  completed BOOLEAN NOT NULL DEFAULT false, 
  created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  updated_date TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE TRIGGER IF NOT EXISTS set_timestamp
BEFORE UPDATE ON todos
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


INSERT INTO todos (title) VALUES ('buy milk');
INSERT INTO todos (title) VALUES ('take out the trash');
INSERT INTO todos (title) VALUES ('learn go and vue');
