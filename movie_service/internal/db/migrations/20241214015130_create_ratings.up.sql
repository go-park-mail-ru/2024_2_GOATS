CREATE TABLE ratings (
         id SERIAL PRIMARY KEY,
         user_id INT NOT NULL,
         movie_id INT NOT NULL,
         rating DECIMAL(10, 2) NOT NULL DEFAULT '0.0' CHECK (rating >= 0 AND rating <= 10),
         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
         updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

         UNIQUE (user_id, movie_id),
         CONSTRAINT fk_movie FOREIGN KEY (movie_id) REFERENCES movies (id) ON DELETE CASCADE
);

CREATE INDEX idx_ratings_user_id ON ratings(user_id);
CREATE INDEX idx_ratings_movie_id ON ratings(movie_id);
