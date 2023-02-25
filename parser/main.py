from parse_pages import parse_seasons, parse_drivers_page, parse_teams_page, parse_track_page, convert_to_csv
import pandas as pd

URL_DRIVERS = 'https://www.f1-portal.ru/?content=info/handbook&sort=drivers'
URL_TEAMS = 'https://www.f1-portal.ru/?content=info/handbook&sort=teams'
URL_TRACKS = 'https://www.f1-portal.ru/?content=info/handbook&sort=tracks'


# Сброс ограничений на количество выводимых рядов
pd.set_option('display.max_rows', None)

# Сброс ограничений на число столбцов
pd.set_option('display.max_columns', None)

# Сброс ограничений на количество символов в записи
pd.set_option('display.max_colwidth', None)


def parse_f1():
    gp, qualifications, races = parse_seasons(2009, 2023)
    drivers = parse_drivers_page(URL_DRIVERS)
    teams = parse_teams_page(URL_TEAMS)
    tracks = parse_track_page(URL_TRACKS)

    convert_to_csv(gp, 'gp.csv')
    convert_to_csv(qualifications, 'qualifications.csv')
    convert_to_csv(races, 'races.csv')
    convert_to_csv(drivers, 'drivers.csv')
    convert_to_csv(teams, 'teams.csv')
    convert_to_csv(tracks, 'tracks.csv')


if __name__ == '__main__':
    parse_f1()





