ALTER TABLE courses
ADD CONSTRAINT FK_COURSES_CATEGORY
FOREIGN KEY (category_id) REFERENCES categories(ID);