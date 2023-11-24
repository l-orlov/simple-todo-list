const taskInput = document.getElementById('taskInput');
const addTaskButton = document.getElementById('addTask');
const taskList1 = document.getElementById('taskList1');
const taskList2 = document.getElementById('taskList2');
const taskList3 = document.getElementById('taskList3');

/*
Задача:
- подключить запросы на server, чтобы данные сохранялись в БД и подгружались из БД при обновление страницы
*/

const url = 'http://localhost:8080/tasks/';

function createTaskElement(id, title, status) {
    console.log(id, title, status)
    const listItem = document.createElement('li');
    listItem.innerHTML = `
        <span>${title}</span>
        <div>
        <button class="toTheBegining">В начало</button>
        <button class="inProcess">В процессе</button>
        <button class="ready">Готово</button>
        <button class="deleteTask">Удалить</button>
        </div>
    `;
    listItem.dataset.taskId = id
    listItem.dataset.taskTitle = title
    listItem.dataset.taskStatus = status


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

    console.log('status = ', status)
    switch (status) {
        case 1:
            taskList1.appendChild(listItem);
            break;
        case 2:
            taskList2.appendChild(listItem);
            break;
        case 3:
            taskList3.appendChild(listItem);
            break;
        default:
            console.log('invalid task status');
    }
}

function fetchTasks() {
    // Отправляем GET-запрос на сервер
    return fetch(url)
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

function createTask(title, status) {
    // Новая таска
    var task = {
        title: title,
        status: status,
    };
    // Опции запроса
    var requestOptions = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(task)
    };
    // Отправляем POST-запрос на сервер
    return fetch(url, requestOptions)
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

// Вызываем функцию и работаем с данными вне неё
// fetchTasks()
//     .then(data => {
//         // Перебираем элементы массива
//         for (let task of data) {
//             // Добавляем таску на общую доску тасок
//             createTaskElement(task.id, task.title, task.status)
//         }
//     })
//     .catch(error => {
//         // Обработка ошибок
//         console.error('There was an error outside fetchData:', error.message);
//     });

addTaskButton.addEventListener('click', async() => {
    const taskTitle = taskInput.value.trim();

    if (taskTitle === '') {
        alert('Пожалуйста, введите задачу.');
        return;
    }

    try {
        const data = await createTask(taskTitle, 1);

        const task = {
            id: data.id,
            status: data.status,
            title: data.title
        };

        console.log('task: ', task);

        createTaskElement(task.id, task.title, task.status);

        taskInput.value = '';
    } catch (error) {
        // Обработка ошибок
        console.error('error creating task:', error.message);
    }

    // createTaskElement(task.id, task.title, task.status)
    //
    // var task = {}
    // createTask(taskTitle, 1)
    //     .then(data => {
    //         task.id = data.id
    //         task.status = data.status
    //         task.title = data.title
    //     })
    //     .catch(error => {
    //         // Обработка ошибок
    //         console.error('There was an error outside fetchData:', error.message);
    //     });
    //
    // console.log('task: ', task)


    // taskInput.value = '';
});
