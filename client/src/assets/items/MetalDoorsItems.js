const metalDoorCreationItems = [
    {
        title: "Для каталога",
        type: "checkbox",
        state_name: "for_catalog"

    },
    {
        title: "Номер",
        type: "text",
        state_name: "number"
    },
    {
        title: "Название",
        type: "text",
        state_name: "title"

    },
    {
        title: "Цвет",
        type: "text",
        // state_name: "number"

    },
    {
        title: "Цена",
        state_name: "prices",
        child: [
            {
                title: "Опт",
                type: "number",
                state_name: "opt"

            },
            {
                title: "Кр.Опт",
                type: "number",
                state_name: "k_opt"
            },
            {
                title: "Вип",
                type: "number",
                state_name: "vip"

            },
            {
                title: "Розница",
                type: "number",
                state_name: "roz"

            },
        ]
    },
    {
        title: "Высота",
        type: "number",
        state_name: "height",
        spec: true,


    },
    {
        title: "Толщина металла",
        type: "number",
        state_name: "metal",
        spec: true,


    },
    {
        title: "Толщина полотна",
        type: "number",
        state_name: "leaf",
        spec: true,


    },
    {
        title: "Верхний замок",
        type: "text",
        state_name: "upper_lock",
        spec: true,


    },
    {
        title: "Нижний замок",
        type: "text",
        state_name: "lower_lock",
        spec: true,

    },
    {
        title: "Задвижка",
        type: "text",
        state_name: "valve",
        spec: true,

    },
    {
        title: "Броненакладка",
        type: "text",
        state_name: "armor",
        spec: true,

    },
    {
        title: "Цилиндр",
        type: "text",
        state_name: "cylinder",
        spec: true,

    },
    {
        title: "Глазок",
        type: "text",
        state_name: "peephole",
        spec: true,

    },
    {
        title: "Утеплитель",
        type: "text",
        state_name: "insulation",
        spec: true,


    },
    {
        title: "Уплотнитель",
        type: "text",
        state_name: "sealer",
        spec: true,

    },
    {
        title: "Ручка",
        type: "text",
        state_name: "handle",
        spec: true,
    },
]

export default metalDoorCreationItems;