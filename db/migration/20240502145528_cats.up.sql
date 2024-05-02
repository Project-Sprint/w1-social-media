CREATE TABLE IF NOT EXISTS "cats" (
  id UUID PRIMARY KEY,
  user_id UUID,
  name VARCHAR(30) NOT NULL CHECK (LENGTH(name) > 0),
  race VARCHAR NOT NULL CHECK (race IN ('Persian', 'Maine Coon', 'Siamese', 'Ragdoll', 'Bengal', 'Sphynx', 'British Shorthair', 'Abyssinian', 'Scottish Fold', 'Birman')),
  sex VARCHAR NOT NULL CHECK (sex IN ('male', 'female')),
  ageInMonth INTEGER NOT NULL CHECK (ageInMonth BETWEEN 1 AND 120082),
  description VARCHAR(200) NOT NULL CHECK (LENGTH(description) > 0),
  imageUrls TEXT[] NOT NULL CHECK (cardinality(imageUrls) >= 1 AND array(SELECT url FROM unnest(imageUrls) AS url WHERE url != '') = imageUrls),
  hasMatched BOOLEAN,
  createdAt TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
