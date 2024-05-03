CREATE TABLE IF NOT EXISTS "cats" (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  name VARCHAR(30) NOT NULL CHECK (LENGTH(name) > 0),
  race VARCHAR NOT NULL CHECK (race IN ('Persian', 'Maine Coon', 'Siamese', 'Ragdoll', 'Bengal', 'Sphynx', 'British Shorthair', 'Abyssinian', 'Scottish Fold', 'Birman')),
  sex VARCHAR NOT NULL CHECK (sex IN ('male', 'female')),
  ageInMonth INTEGER NOT NULL CHECK (ageInMonth BETWEEN 1 AND 120082),
  description VARCHAR(200) NOT NULL CHECK (LENGTH(description) > 0),
  imageUrls TEXT[] NOT NULL,
  hasMatched BOOLEAN,
  createdAt TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
