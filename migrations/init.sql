CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE roles (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(50) UNIQUE NOT NULL,
  description text
);

CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  username varchar(100) UNIQUE NOT NULL,
  email varchar(150) UNIQUE NOT NULL,
  password_hash varchar(255) NOT NULL,
  full_name varchar(150),
  role_id uuid REFERENCES roles(id),
  created_at timestamptz DEFAULT now()
);

CREATE TABLE students (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid REFERENCES users(id),
  student_number varchar(50) UNIQUE,
  program_study varchar(150)
);

CREATE TABLE lecturers (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid REFERENCES users(id),
  lecturer_number varchar(50) UNIQUE,
  department varchar(150)
);

CREATE TABLE achievement_references (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  student_id uuid NOT NULL,
  mongo_achievement_id varchar(24) NOT NULL,
  status varchar(20) DEFAULT 'draft',
  submitted_at timestamptz,
  verified_at timestamptz,
  verified_by uuid,
  rejection_note text,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now()
);
