# О программе

Позволяет сохранять в файл, извлекать и манипулировать данными, введёнными пользователем.

Данные состоят из логина, пароля, расположения аккаунта и времени его создания

## Encrypted and decrypted

[Tutorial1](https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/)

[Tutorial2](https://dev.to/breda/secret-key-encryption-with-go-using-aes-316d)


Структура node имеет поле password, которое должно сохраняться в закодированном виде. 

Массив записей собирается в одну структуру (слайс строк), после чего сохраняется 
в отдельный файл с расширением .encr.

Ключ для шифровки - расшифровки необходимо расположить в новом файле с расширением .env, 
со значением KEY=V6U0GjlC97InPOPs9OIHAahRN3j3ZYPj

Также возможно сгенерировать новый ключ для новых записей (для старых - нужен прежний ключ)
[random key generator](https://acte.ltd/utils/randomkeygen)