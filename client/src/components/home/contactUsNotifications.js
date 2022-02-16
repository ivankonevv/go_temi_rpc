const contactUsNotifications = {
    success: {
        group: 'test',
        title: 'Успешно',
        text: 'Ваша заявка успешно отправлена.',
        type: 'success'
    },
    error: {
        incorrect_input: {
            group: 'test',
            title: 'Неправильный ввод!',
            text: 'Кажется, вы некорректно заполнили форму.',
            type: 'error'
        },
        internal: {
            group: "test",
            title: "Ошибка",
            text: "Похоже, произошла какая-то ошибка. Попробуйте позже или позвоните по любому из указанных номеров.",
            type: "error"
        }
    },

}

export default contactUsNotifications;