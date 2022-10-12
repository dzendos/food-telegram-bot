# Food-telegram-bot
<img src = "./img/logo.jpg?raw=true" width = "30%" height = "30%" alt = "logo" align = center />

Простой бот для заказа доставки с сайтов без использования сторонних приложений.

## Мотивация

Создать телеграм бота в котором булет удобно:

* Заказывать доставку без установки сторонних приложений.
* Объединять доставку с друзьями.
* Разделять сумму заказа в зависимости от его состава.
* Упрощать оплату и вести учет долгов.

## Функционал

Бот имеет следующие функции:

* Создать заказ, выбрав ресторан доставки
* Создать ссылку-приглашение на вступление в общий заказ
* Собрать заказ из выбранного ресторана
* Собрать всю еду, выбранную всеми участниками заказа в один лист и рассчитать стоимость полученного заказа
* Оповещать всех участников заказа о необходимой сумме к оплате, с учетом доставки. Дополнительно присылать номер телефона заказчика, на которые необходимо перевести сумму

Благодаря многослойной архитектуре приложения бота легко расширять и перенастраивать под любой вид доставки.

## Дизайн

* Изменение дизайна в зависимости от темы приложения телеграм

Light theme     | Dark theme
:--------------:|:------------:
<img src = "./img/light_menu.jpg?raw=true" width = "30%" height = "30%" alt = "light_menu" align = center /> | <img src = "./img/darkj_menu.jpg?raw=true" width = "30%" height = "30%" alt = "darkj_menu" align = center />

* Всплывающее окно с возможностью выбора ресторанов и последующим составлением заказа
<img src = "./img/restaurant_choose.jpg?raw=true" width = "40%" height = "40%" alt = "restaurant_choose" align = center />

* Просмотр текущего состояния заказа

Light theme     | Dark theme
:--------------:|:------------:
<img src = "./img/list_of_order.png?raw=true" width = "30%" height = "30%" alt = "list_of_order" align = center /> | <img src = "./img/list_of_order_dark.png?raw=true" width = "30%" height = "30%" alt = "list_of_order_dark" align = center />

* Настраиваемое сообщение, отправляемое всем участникам заказа

## Команды

* /new_order - создать новый заказ
* /my_order - показать, что в вашем заказе
* /full_order - показать общий заказ
* /set_transaction_message - установить сообщение, приходящее всем коллегам, которые участвуют в вашем заказе
* /confirm order - подвтердить заказ

<img src = "./img/list_of_commands.jpg?raw=true" width = "30%" height = "30%" alt = "commands" align = center />

## Необходимые библиотеки
* [Go](https://go.dev/doc/install)
* [gotgbot](https://github.com/PaulSonOfLars/gotgbot)
* [goquery](https://github.com/PuerkitoBio/goquery)
* [errors](https://github.com/pkg/errors)

<img src = "./img/go.jpg?raw=true" width = "50%" height = "50%" alt = "go" align = center />

Используйте `go mod tidy` для того чтобы установить все необходимые пакеты

## Как запустить
Создайте файл config.yaml в папке data на примере config_example.yaml

Его структура такая:
* token - токен вашего бота
* port - порт через который ваш сервер будет слушать запросы с веб оболочки бота
* url - ссылка по которой будет хоститься веб оболочка бота

собирать при помощи `make run`
