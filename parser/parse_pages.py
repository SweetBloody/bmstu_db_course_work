import requests
from bs4 import BeautifulSoup
import pandas as pd
import re
import pycountry
import gettext

URL_SEASON_RESULTS = 'https://www.f1-portal.ru/?content=season%2Fresults&year='
PATH = 'parsed_data/'



# Запрос доступа к странице
def get_page(url):
    response = requests.get(url)
    if response.status_code != 200:
        print("Request error! URL: {}".format(url))
        return None
    else:
        print("Connection success! URL: {}".format(url))
    response.encoding = 'windows-1251'
    page = BeautifulSoup(response.text, 'lxml')
    return page


# Парсинг страницы с таблицей результатов (квалификация/гонка)
def parse_table(url, gp_ind, race=True):
    page = get_page(url)
    if page == None:
        return None
    table = page.find('table', 'table table-condensed table-bordered table-striped text-center')
    header = []
    for elem in table.find_all('th', 'text-center'):
        header.append(elem.text)
    header.append("gp_ind")
    data = []
    for row in table.find_all('tr'):
        data.append([col.text for col in row.find_all('td')] + [gp_ind])
    if race:
        data_table = pd.DataFrame(data[1:-2], columns=header)
    else:
        data_table = pd.DataFrame(data[1:], columns=header)
    # print(data_table)
    return data_table


# Парсинг блока с информацией о ГП
def parse_gp_block(gp_block):
    gp_name = gp_block.find('div', 'h4 list-group-item-heading')
    gp_dig = gp_block.find('div', 'gp-dig')
    gp_month = gp_block.find('small')
    gp_place = gp_block.find('em', 'text-muted')
    return [gp_name.text, gp_dig.text, gp_month.text, gp_place.text]


# Парсинг подблока с ссылками на результаты ГП
# type="race" - информация о результатах гонки
# type="qual" - информация о результатах квалификации
def parse_gp_results(gp_block, gp_ind, type="race"):
    gp_res = gp_block.find('a', 'list-group-item', href=re.compile(type))
    url = 'https://www.f1-portal.ru/' + gp_res.get('href')
    race = (type == "race")
    table = parse_table(url, gp_ind, race)
    # if table == None:
    #     return None
    return table


# Парсинг страницы сезона Формулы 1
def parse_season_page(url, gp_ind):
    season_page = get_page(url)
    if season_page == None:
        return None
    year = url[-4:]
    gp_info = []
    qual_info = pd.DataFrame()
    race_info = pd.DataFrame()
    for gp_block in season_page.find_all('div', 'list-group calendar-list'):
        block_info = parse_gp_block(gp_block)
        # gp_ind = len(gp_info)
        gp_info.append([year] + block_info)
        qual_info = pd.concat([qual_info, parse_gp_results(gp_block, gp_ind, "qual")], ignore_index=True)
        race_info = pd.concat([race_info, parse_gp_results(gp_block, gp_ind, "race")], ignore_index=True)
        gp_ind += 1
    # print(qual_info)
    gp_info = pd.DataFrame(gp_info)
    # print(gp_info)
    return gp_info, qual_info, race_info


# Парсим все сезоны
def parse_seasons(season_start, season_end):
    gp = pd.DataFrame()
    qualifications = pd.DataFrame()
    races = pd.DataFrame()

    for year in range(season_start, season_end):
        url = URL_SEASON_RESULTS + str(year)
        gp_ind = gp.index.size
        temp_gp, temp_qual, temp_race = parse_season_page(url, gp_ind)
        gp = pd.concat([gp, temp_gp], ignore_index=True)
        qualifications = pd.concat([qualifications, temp_qual], ignore_index=True)
        races = pd.concat([races, temp_race], ignore_index=True)
    return gp, qualifications, races



# Парсим страницу со всеми пилотами
def parse_drivers_page(url):
    driver_page = get_page(url)
    if driver_page == None:
        return None
    header = ['name', 'country', 'birth_date'   ]
    drivers = []
    delete_chars = str.maketrans('', '', '\t\n')
    for driver_block in driver_page.find_all('div', 'col-xs-12 col-sm-12 col-md-6 col-lg-6'):
        country = driver_block.find('div', re.compile('gp-flag flag')).get('title')
        name = driver_block.find('div', 'h4 media-heading').text.translate(delete_chars)
        birth_date = driver_block.find('p', 'text-muted text-ellipsis').get('title')
        drivers.append([name, country, birth_date])
    drivers_df = pd.DataFrame(drivers, columns=header)
    # print(drivers_df)
    return drivers_df


# Перевод кода страны в полное название
def get_country(code):
    russian = gettext.translation("iso3166", pycountry.LOCALES_DIR, languages=["ru"])
    russian.install()
    ru = pycountry.countries.get(alpha_2=code)
    country = _(ru.name)
    if country == 'Соединённое Королевство':
        country = 'Великобритания'
    elif country == 'Соединённые штаты':
        country = 'США'
    return country


# Парсинг страницы с командами
def parse_teams_page(url):
    teams_page = get_page(url)
    if teams_page == None:
        return None
    header = ['name', 'country', 'base']
    teams = []
    delete_chars = str.maketrans('', '', '\t\n')
    for team_block in teams_page.find_all('div', 'col-xs-12 col-sm-12 col-md-6 col-lg-6 f32'):
        country = team_block.find('div', re.compile('gp-flag flag')).get('class')[-1]
        country = get_country(country)
        name = team_block.find('div', 'h4 media-heading').text.translate(delete_chars)
        base = team_block.find('p', 'text-muted text-ellipsis').get('title')
        teams.append([name, country, base])
    teams_df = pd.DataFrame(teams, columns=header)
    print(teams_df)
    return teams_df


# Парсинг страницы с трассами
def parse_track_page(url):
    track_page = get_page(url)
    if track_page == None:
        return None
    header = ['name', 'country', 'town']
    tracks = []
    delete_chars = str.maketrans('', '', '\t\n')
    for track_block in track_page.find_all('div', 'col-xs-12 col-sm-6 col-md-4 col-lg-4 f32'):
        country = track_block.find('div', re.compile('gp-flag flag')).get('class')[-1]
        country = get_country(country)
        name = track_block.find('div', 'h4 media-heading').text.translate(delete_chars)
        town = track_block.find('p', 'text-muted text-ellipsis').text.split(',')[0]
        tracks.append([name, country, town])
    tracks_df = pd.DataFrame(tracks, columns=header)
    return tracks_df


# Перевод в csv файл
def convert_to_csv(data_frame, file_name, index=True):
    path_file = PATH + file_name
    data_frame.to_csv(path_file, sep=';', index=index)