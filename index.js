const taskInput = document.getElementById('taskInput');
const addTaskButton = document.getElementById('addTask');
const taskList1 = document.getElementById('taskList1');
const taskList2 = document.getElementById('taskList2');
const taskList3 = document.getElementById('taskList3');

/*
Задача:
- подключить запросы на server, чтобы данные сохранялись в БД и подгружались из БД при обновление страницы
*/

function fetchTask() {
    return fetch('http://localhost:8080/tasks/get')
        .then(response => {
            // Проверка на успешный ответ (код 200-299)
            if (!response.ok) {
                throw new Error(`Network response was not ok, status: ${response.status}`);
            }
            // Преобразование ответа в JSON
            return response.json();
        })
        .catch(error => {
            // Обработка ошибок
            console.error('There was a problem with the fetch operation:', error.message);
        });
}

function createTask(taskText) {
    const listItem = document.createElement('li');
    listItem.innerHTML = `
        <span>${taskText}</span>
        <div>
        <button class="toTheBegining">В начало</button>
        <button class="inProcess">В процессе</button>
        <button class="ready">Готово</button>
        <button class="deleteTask">Удалить</button>
        </div>
    `;

    const processButton1 = listItem.querySelector('.inProcess');
    processButton1.addEventListener('click', () => {

        taskList2.appendChild(listItem);
    });

    const readyButton1 = listItem.querySelector('.ready');
    readyButton1.addEventListener('click', () => {
        taskList3.appendChild(listItem);

    });

    const beginButton1 = listItem.querySelector('.toTheBegining');
    beginButton1.addEventListener('click', () => {
        taskList1.appendChild(listItem);

    });

    const deleteButton1 = listItem.querySelector('.deleteTask');
    deleteButton1.addEventListener('click', (event) => {
        const currentList = event.target.parentNode.parentNode.parentNode;
        if (currentList == taskList1) taskList1.removeChild(listItem);
        if (currentList == taskList2) taskList2.removeChild(listItem);
        if (currentList == taskList3) taskList3.removeChild(listItem);
    });

    return listItem
}

// Вызываем функцию и работаем с данными вне неё
fetchTask()
    .then(data => {
        // Перебираем элементы массива
        for (let item of data) {
            console.log(item);
            // Дополнительная обработка каждого элемента

            const currentListItem = createTask(item.title)

            switch (item.status) {
                case 1:
                    console.log('ToDo');
                    taskList1.appendChild(currentListItem);
                    break;
                case 2:
                    console.log('InProgress');
                    taskList2.appendChild(currentListItem);
                    break;
                case 3:
                    console.log('Done');
                    taskList3.appendChild(currentListItem);
                    break;
                case 4:
                    console.log('Deleted');
                    break;
                default:
                    console.log('Invalid selection.');
            }
        }
    })
    .catch(error => {
        // Обработка ошибок
        console.error('There was an error outside fetchData:', error.message);
    });

addTaskButton.addEventListener('click', () => {
    const taskText = taskInput.value.trim();

    if (taskText === '') {
        alert('Пожалуйста, введите задачу.');
        return;
    }

    const listItem = createTask(taskText)
    taskList1.appendChild(listItem);

    taskInput.value = '';
});
