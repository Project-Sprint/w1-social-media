CREATE TABLE IF NOT EXISTS "matchs" (
  id UUID PRIMARY KEY,
  matchCatId UUID NOT NULL,
  userCatId UUID NOT NULL,
  userId UUID NOT NULL,
  targetUserId UUID NOT NULL,
  message VARCHAR(120) NOT NULL CHECK (LENGTH(message) >= 5 AND LENGTH(message) <= 120),
  createdAt TIMESTAMP,
  FOREIGN KEY (matchCatId) REFERENCES cats(id),
  FOREIGN KEY (userCatId) REFERENCES cats(id),
  FOREIGN KEY (userId) REFERENCES users(id),
  FOREIGN KEY (targetUserId) REFERENCES users(id)
);
