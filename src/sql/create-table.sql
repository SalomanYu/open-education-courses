create table open_education(
    url text not null,
    title text not null,
    started_at date,
    finished_at date,
    img text,
    teachers_name text[],
    teachers_image text[],
    teachers_description text[],
    skills text,
    description text,
    requirements text,
    duration_in_week integer,
    lectures_count integer,
    has_certificate boolean,
    check (duration_in_week > 0)
);