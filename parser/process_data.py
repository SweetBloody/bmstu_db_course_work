import pandas as pd
from parse_pages import convert_to_csv, PATH

# # Сброс ограничений на количество выводимых рядов
# pd.set_option('display.max_rows', None)
#
# # Сброс ограничений на число столбцов
# pd.set_option('display.max_columns', None)
#
# # Сброс ограничений на количество символов в записи
# pd.set_option('display.max_colwidth', None)

def convert_from_csv(file_name):
    data_frame = pd.read_csv(PATH + file_name, sep=';')
    return data_frame


def get_df_dict(file_name):
    df = convert_from_csv(file_name)
    diction = dict()
    for elem in df.iterrows():
        diction.update({elem[1]['name']: elem[0]})
    # print(drivers_dict)
    return diction


def update_team_names(df):
    df = df.replace(['BMW Sauber', 'Mercedes GP', 'Lotus Renault GP', 'Team Lotus', 'Marussia Virgin', 'Alfa Romeo Racing'],
               ['BMW', 'Mercedes', 'Renault', 'Lotus', 'Marussia', 'Alfa Romeo F1 Team'])
    return df


def update_results(df):
    df = df.replace(['Столкновение', 'Дифференциал', 'Подвеска', 'Разворот', 'Повреждение', 'Двигатель', 'Трансмиссия', 'Электрика', 'Авария', 'Давление масла', 'Коробка передач', 'Сход'],
               'DNF')
    return df


def update_driver_names(df):
    for elem in df.iterrows():
        name = elem[1]['Пилот']
        words = name.split(' ')
        if words[-1].isnumeric():
            df = df.replace([name], ' '.join(words[:-1]))
    # print(df)
    return df


def change_names_to_id(df, diction):
    for elem in diction:
        df = df.replace([elem], diction[elem])
    return df



if __name__ == '__main__':
    drivers_dict = get_df_dict('drivers.csv')
    teams_dict = get_df_dict('teams.csv')

    races_df = convert_from_csv('races.csv')
    races_df = update_driver_names(races_df)
    races_df = update_team_names(races_df)
    races_df = update_results(races_df)
    races_df = change_names_to_id(races_df, drivers_dict)
    races_df = change_names_to_id(races_df, teams_dict)

    qual_df = convert_from_csv('qualifications.csv')
    qual_df = update_driver_names(qual_df)
    qual_df = update_team_names(qual_df)
    qual_df = change_names_to_id(qual_df, drivers_dict)
    qual_df = change_names_to_id(qual_df, teams_dict)

    convert_to_csv(races_df, 'races.csv', False)
    convert_to_csv(qual_df, 'qualifications.csv', False)

    print(races_df)
    print(qual_df)
