CREATE TABLE todos (
  uuid            char(36) unique primary key,
  title           varchar(40),
  description     varchar(350),
  status          integer,
  priority        integer,
  creation_date   timestamp,
  due_date        timestamp
);
