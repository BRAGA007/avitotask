
Все данные денег поступают в значении int и были приняты мною за копейки, так как это самый безопасный способ хранить деньги в базе данных.

Изначально поступает пустая таблица таблица.

Также все запросы реализованы методом POST так как отправлять GET запрос с body не принято.
#
Для того чтобы установить проект необходимо прописать в консоли git clone https://github.com/BRAGA007/avitotask

Для запуска проекта необходимо прописать docker-compose up --build
#
Создадим Userов c id 1 и id 2. Их будет достаточно чтобы рассмотреть весь функционал микросервиса.
![image](https://user-images.githubusercontent.com/66581773/202797881-75f57223-294e-448e-902b-6ce2d5e56cdf.png)
![image](https://user-images.githubusercontent.com/66581773/202797943-2aeda114-648e-47be-b3aa-b2b2a1d09414.png)
Теперь повторим тот же запрос для usera с id 1 его баланс увеличивается в задании.
![image](https://user-images.githubusercontent.com/66581773/202797912-e4c57066-1fe4-47d9-9962-c1abdd38bb2e.png)
Снятие денег со счета usera c id 1
![image](https://user-images.githubusercontent.com/66581773/202798121-131f1ebb-fa7f-460f-a52f-38bed5bd8a0b.png)
Показ баланса по id usera c id 1
![image](https://user-images.githubusercontent.com/66581773/202798200-8e1ae85c-df4f-4b1f-9c3d-a60642494a1e.png)
Перевод от первого пользователя второму
![image](https://user-images.githubusercontent.com/66581773/202798391-3c79f8b9-d700-4564-9ac1-84dbb49f4540.png)
Резервация денег со счета. Подтверждение списания денег происходит путем отправки такого же JSON
![image](https://user-images.githubusercontent.com/66581773/202798559-76be9fe9-c4e8-4d0e-b9f2-bafd857208f8.png)
Формирование .csv отчета о прибыли
![image](https://user-images.githubusercontent.com/66581773/202798983-176e3c0a-dc91-437d-8f68-5b0567744fa9.png)
Просмотр истории баланса. Можно менять сортировку и пагинацию путем изменения JSON
![image](https://user-images.githubusercontent.com/66581773/202799581-0368d33e-8ac1-4741-836b-67846917d4d2.png)
  





