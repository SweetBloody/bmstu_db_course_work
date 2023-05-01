-- drop database if exists formula1;
-- create database formula1;

drop table if exists GrandPrix cascade;
create table public.GrandPrix(
    gp_id serial not null primary key,
    gp_season int not null check (gp_season > 1949),
    gp_name text not null,
    gp_date_num int not null check (gp_date_num between 0 and 32),
    gp_month text not null,
    gp_place text not null,
    gp_track_id int not null
);

drop table if exists Tracks cascade;
create table public.Tracks(
    track_id serial not null primary key,
    track_name text not null,
    track_country text not null,
    track_town text not null
);

drop table if exists QualificationResults cascade;
create table public.QualificationResults(
    qual_id serial not null primary key,
    qual_driver_place int not null,
    driver_id int not null,
    team_id int not null,
    q1_time time,
    q2_time time,
    q3_time time,
    gp_id int not null
);

drop table if exists RaceResults cascade;
create table public.RaceResults(
    race_id serial not null primary key,
    race_driver_place int,
    driver_id int not null,
    team_id int not null,
    gp_id int not null
);

drop table if exists Drivers cascade;
create table public.Drivers(
    driver_id serial not null primary key,
    driver_name text not null,
    driver_country text,
    driver_birth_date date
);

drop table if exists Teams cascade;
create table public.Teams(
    team_id serial not null primary key,
    team_name text not null,
    team_country text not null,
    team_base text not null
);

drop table if exists TeamsDrivers cascade;
create table public.TeamsDrivers(
    td_id serial not null primary key,
    driver_id int not null,
    team_id int not null,
    team_driver_season int not null check (team_driver_season > 1949)
);

drop table if exists Users cascade;
create table public.Users(
    user_id serial not null primary key,
    login text not null,
    password text not null,
    role int not null
);



set datestyle to 'dmy';
alter table GrandPrix add foreign key (gp_track_id) references public.Tracks(track_id);
alter table QualificationResults add foreign key (gp_id) references public.GrandPrix(gp_id);
alter table RaceResults add foreign key (gp_id) references public.GrandPrix(gp_id);
alter table TeamsDrivers add foreign key (driver_id) references public.Drivers(driver_id);
alter table TeamsDrivers add foreign key (team_id) references public.Teams(team_id);
-- alter table TeamsDrivers add primary key (driver_id, team_id);

