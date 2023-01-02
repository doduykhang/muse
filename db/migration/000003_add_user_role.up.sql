ALTER TABLE accounts
ADD COLUMN role VARCHAR(10) CHECK (role = 'ADMIN' OR role = 'USER')
