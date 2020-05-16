TOBACCOS = [
            ('Duft', 'Raspberry'),
            ('Duft', 'Cheesecake'),
            ('Musthave', 'Raspberry'),
            ('Musthave', 'Grapefruit'),
            ('Musthave', 'Orange'),
            ('Darkside', 'Lemon'),
            ('Darkside', 'Menthol'),
            ('Daily Hookah', 'Lemon'),
            ('Hypreme', 'Menthol'),
        ]

MIXES = [
            ('Raspberry', 'Grapefruit', 'Lemon', 'Menthol'),
            ('Raspberry', 'Cheesecake', 'Orange', 'Menthol'),
        ]


# Выбрать все табаки всех вкусов, которые нравятся пользователю
# Взять все миксы, в которых есть хотя бы 1 табак который нравится пользователю
# Составить хешмап вкусов к списку производителей
# Для каждого найденного микса собрать список производителей по каждому вкусу
def find_mixes(tobaccos, mixes):
    tobacco_dict = {}
    for manufacturer, flavour in tobaccos:
        if flavour not in tobacco_dict:
            tobacco_dict[flavour] = [manufacturer]
        elif manufacturer not in tobacco_dict[flavour]:
            tobacco_dict[flavour].append(manufacturer)

    result = []
    for mix in mixes:
        aggr_flavours = []
        for flavour in mix:
            aggr_flavours.append((flavour, tobacco_dict.get(flavour, [])))
        result.append(aggr_flavours)

    return result


def solve():
    res = find_mixes(TOBACCOS, MIXES)

    # pretty print
    for i, mix in enumerate(res):
        flavours = []
        l = ', '.join([x for x, _ in mix])
        print(f'Mix {i+1}:', l, '\n')
        print('-'*(len(l)+7))

        for flavour, manufacturers in mix:
            ml = ', '.join(manufacturers)
            print(f'{flavour} -> {ml}')

        if i != len(res) - 1 :
            print('\n###\n')


if __name__ == '__main__':
    solve()
