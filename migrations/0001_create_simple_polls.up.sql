CREATE TABLE
    IF NOT EXISTS questions (
        id serial primary key,
        title VARCHAR not null
    );

CREATE TABLE
    IF NOT EXISTS choices (
        id serial primary key,
        choice_text VARCHAR not null,
        question_id integer,
        votes integer not null DEFAULT(0),
        foreign key(question_id) references questions(id) on delete cascade
    );