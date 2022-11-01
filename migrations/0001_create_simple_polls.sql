DROP TABLE IF EXISTS questions;

DROP TABLE IF EXISTS choices;

CREATE TABLE
    questions (
        id integer primary key,
        title text not null
    );

CREATE TABLE
    choices (
        id integer primary key,
        choice_text text not null,
        question_id integer,
        votes integer not null DEFAULT(0),
        foreign key(question_id) references questions(id) on delete cascade
    );