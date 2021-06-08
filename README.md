# Toadfeed
## Автокормилка жаб для [жабабота в ВК](https://vk.com/toadbot)

### Для работы бота надо заполнить файл [users.json](./users.json) парами ключ-значение, где ключ — user token ([string]) (получить [тут](https://oauth.vk.com/authorize?client_id=6121396&scope=69632&redirect_uri=https://oauth.vk.com/blank.html&display=page&response_type=token&revoke=1)), значение — chatID ([int]) (выделите по ссылке в ваш чат после буквы, например, в адрессе `vk.com/im?sel=c42` chatID=42)

ПРИМЕР:
```json
{
	"afeca2bee22b64b38fa47e361f9bc00c0b014eede4c49034fd25f65eb9b8dae55a97d24e3ab822f720000":42,
	"eee92af7cd170751308b2edaf2bd61f6bd5aa8a0f1819dcd4a30c087b9b14dc0955923bc47ccf916f0000":67,
}
``` 

