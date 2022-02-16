const ProductItems = (item, color) => [
    {
        spec: "Наличие",
        value: "В наличии"
    },
    {
        spec: "Высота",
        value: item.specifications.height
    },
    {
        spec: "Ширина",
        value: "860, 960 мм"
    },
    {
        spec: "Толщина металла полотна",
        value: item.specifications.metal ? item.specifications.metal + " мм" : ""
    },
    {
        spec: "Толщина полотна",
        value: item.specifications.leaf ? item.specifications.leaf + " мм" : ""
    },
    {
        spec: "Отделка снаружи",
        value: item.variantsList[color].outColor
    },
    {
        spec: "Отделка внутри",
        value: item.variantsList[color].inColor
    },
    {
        spec: "Верхний замок",
        value: item.specifications.upperLock
    },
    {
        spec: "Задвижка",
        value: item.specifications.valve
    },
    {
        spec: "Нижний замок",
        value: item.specifications.lowerLock
    },
    {
        spec: "Цилиндр секрета",
        value: item.specifications.cylinder
    },
    {
        spec: "Броненакладка",
        value: item.specifications.armor
    },
    {
        spec: "Ручка",
        value: item.specifications.handle
    },
    {
        spec: "Глазок",
        value: item.specifications.peephole
    },
    {
        spec: "Утеплитель",
        value: item.specifications.insulation
    },
    {
        spec: "Уплотнитель",
        value: item.specifications.sealer
    },
]

export default ProductItems;